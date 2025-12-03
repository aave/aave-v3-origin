package liquidation

import (
  "fmt"
  "math/big"

  "github.com/ethereum/go-ethereum/common"
)

// GetUserPositionFullInfo replicates LiquidationDataProvider.getUserPositionFullInfo using local data.
func GetUserPositionFullInfo(input *SimulationInput) (*UserPositionFullInfo, error) {
  totalCollateral := big.NewInt(0)
  totalDebt := big.NewInt(0)
  ltvSum := big.NewInt(0)
  liqThresholdSum := big.NewInt(0)

  for _, position := range input.User.Positions {
    reserve, ok := input.Reserves[position.Asset]
    if !ok {
      return nil, fmt.Errorf("reserve data missing for asset %s", position.Asset.Hex())
    }

    unit := assetUnit(reserve.Config.Decimals)
    price := cloneBig(reserve.PriceInBaseCurrency)

    actualCollateralBalance := rayMul(position.CollateralScaledBalance, reserve.NormalizedIncome)
    actualDebtBalance := rayMul(position.DebtScaledBalance, reserve.NormalizedDebt)

    if position.UseAsCollateral && reserve.Config.LiquidationThreshold > 0 && !isZero(actualCollateralBalance) {
      collateralBase := div(mul(actualCollateralBalance, price), unit)
      totalCollateral = add(totalCollateral, collateralBase)

      ltv := reserve.Config.Ltv
      liqThreshold := reserve.Config.LiquidationThreshold
      if input.User.UserEModeCategory != 0 {
        if category, ok := input.Global.EModeCategories[input.User.UserEModeCategory]; ok &&
          reserveEnabledOnBitmap(category.CollateralBitmap, reserve.ID) {
          ltv = category.Ltv
          liqThreshold = category.LiquidationThreshold
        }
      }

      ltvSum = add(ltvSum, mul(collateralBase, bigFromUint64(ltv)))
      liqThresholdSum = add(liqThresholdSum, mul(collateralBase, bigFromUint64(liqThreshold)))
    }

    if position.IsBorrowing && !isZero(actualDebtBalance) {
      debtBase := div(mul(actualDebtBalance, price), unit)
      totalDebt = add(totalDebt, debtBase)
    }
  }

  avgLtv := big.NewInt(0)
  avgLiquidationThreshold := big.NewInt(0)
  if totalCollateral.Sign() > 0 {
    avgLtv = div(ltvSum, totalCollateral)
    avgLiquidationThreshold = div(liqThresholdSum, totalCollateral)
  }

  healthFactor := big.NewInt(0)
  if totalDebt.Sign() == 0 {
    healthFactor = cloneBig(maxUint256)
  } else {
    healthFactor = wadDiv(percentMul(totalCollateral, avgLiquidationThreshold), totalDebt)
  }

  return &UserPositionFullInfo{
    TotalCollateralBase:         totalCollateral,
    TotalDebtBase:               totalDebt,
    CurrentLiquidationThreshold: avgLiquidationThreshold,
    Ltv:                         avgLtv,
    HealthFactor:                healthFactor,
  }, nil
}

// GetLiquidationInfo mirrors LiquidationDataProvider.getLiquidationInfo with max debt liquidation.
func GetLiquidationInfo(input *SimulationInput, collateralAsset, debtAsset common.Address) (*LiquidationInfo, error) {
  return GetLiquidationInfoWithAmount(input, collateralAsset, debtAsset, nil)
}

// GetLiquidationInfoWithAmount mirrors LiquidationDataProvider.getLiquidationInfo with a custom debt cap.
func GetLiquidationInfoWithAmount(
  input *SimulationInput,
  collateralAsset,
  debtAsset common.Address,
  debtLiquidationAmount *big.Int,
) (*LiquidationInfo, error) {
  if debtLiquidationAmount == nil {
    debtLiquidationAmount = cloneBig(maxUint256)
  }

  liquidationInfo := &LiquidationInfo{
    MaxCollateralToLiquidate:      big.NewInt(0),
    MaxDebtToLiquidate:            big.NewInt(0),
    LiquidationProtocolFee:        big.NewInt(0),
    AmountToPassToLiquidationCall: big.NewInt(0),
  }

  userInfo, err := GetUserPositionFullInfo(input)
  if err != nil {
    return nil, err
  }
  liquidationInfo.UserInfo = userInfo

  collateralInfo, err := getCollateralFullInfo(input, collateralAsset)
  if err != nil {
    return nil, err
  }
  debtInfo, err := getDebtFullInfo(input, debtAsset)
  if err != nil {
    return nil, err
  }
  if isZero(debtInfo.DebtBalance) {
    return liquidationInfo, nil
  }

  liquidationInfo.CollateralInfo = collateralInfo
  liquidationInfo.DebtInfo = debtInfo

  if !canLiquidateHealthFactor(liquidationInfo.UserInfo.HealthFactor, input.Global.LiquidationAllowed) {
    return liquidationInfo, nil
  }

  collateralReserve, ok := input.Reserves[collateralAsset]
  if !ok {
    return nil, fmt.Errorf("reserve data missing for asset %s", collateralAsset.Hex())
  }
  debtReserve, ok := input.Reserves[debtAsset]
  if !ok {
    return nil, fmt.Errorf("reserve data missing for asset %s", debtAsset.Hex())
  }

  if !isReserveReadyForLiquidations(collateralReserve, true, input.Global.CurrentTimestamp) ||
    !isReserveReadyForLiquidations(debtReserve, false, input.Global.CurrentTimestamp) {
    return liquidationInfo, nil
  }

  if !isCollateralEnabledForUser(input.User, collateralReserve) {
    return liquidationInfo, nil
  }

  liquidationBonus := getLiquidationBonus(input, collateralReserve)
  maxDebtToLiquidate := getMaxDebtToLiquidate(liquidationInfo.UserInfo, collateralInfo, debtInfo, debtLiquidationAmount)

  collateralAmountToLiquidate, debtAmountToLiquidate, liquidationProtocolFee := getAvailableCollateralAndDebtToLiquidate(
    maxDebtToLiquidate,
    liquidationBonus,
    collateralInfo,
    debtInfo,
    &collateralReserve.Config,
  )

  maxCollateralToLiquidate, maxDebtToLiquidate, liquidationProtocolFee := adjustAmountsForGoodLeftovers(
    collateralAmountToLiquidate,
    debtAmountToLiquidate,
    liquidationProtocolFee,
    liquidationBonus,
    collateralInfo,
    debtInfo,
    &collateralReserve.Config,
  )

  liquidationInfo.MaxCollateralToLiquidate = maxCollateralToLiquidate
  liquidationInfo.MaxDebtToLiquidate = maxDebtToLiquidate
  liquidationInfo.LiquidationProtocolFee = liquidationProtocolFee

  if (!isZero(maxDebtToLiquidate) && maxDebtToLiquidate.Cmp(debtInfo.DebtBalance) == 0) ||
    (!isZero(maxCollateralToLiquidate) && maxCollateralToLiquidate.Cmp(collateralInfo.CollateralBalance) == 0) {
    liquidationInfo.AmountToPassToLiquidationCall = cloneBig(maxUint256)
  } else {
    liquidationInfo.AmountToPassToLiquidationCall = cloneBig(maxDebtToLiquidate)
  }

  return liquidationInfo, nil
}

func getCollateralFullInfo(input *SimulationInput, collateralAsset common.Address) (*CollateralFullInfo, error) {
  reserve, ok := input.Reserves[collateralAsset]
  if !ok {
    return nil, fmt.Errorf("reserve data missing for asset %s", collateralAsset.Hex())
  }

  unit := assetUnit(reserve.Config.Decimals)
  price := cloneBig(reserve.PriceInBaseCurrency)
  position, ok := findPosition(input.User.Positions, collateralAsset)
  if !ok {
    return nil, fmt.Errorf("position data missing for asset %s", collateralAsset.Hex())
  }
  balance := rayMul(position.CollateralScaledBalance, reserve.NormalizedIncome)
  balanceInBase := div(mul(balance, price), unit)

  return &CollateralFullInfo{
    AToken:                  reserve.ATokenAddress,
    CollateralBalance:       balance,
    CollateralBalanceInBase: balanceInBase,
    Price:                   price,
    AssetUnit:               unit,
  }, nil
}

func getDebtFullInfo(input *SimulationInput, debtAsset common.Address) (*DebtFullInfo, error) {
  reserve, ok := input.Reserves[debtAsset]
  if !ok {
    return nil, fmt.Errorf("reserve data missing for asset %s", debtAsset.Hex())
  }

  unit := assetUnit(reserve.Config.Decimals)
  price := cloneBig(reserve.PriceInBaseCurrency)
  position, ok := findPosition(input.User.Positions, debtAsset)
  if !ok {
    return nil, fmt.Errorf("position data missing for asset %s", debtAsset.Hex())
  }
  balance := rayMul(position.DebtScaledBalance, reserve.NormalizedDebt)
  balanceInBase := div(mul(balance, price), unit)

  return &DebtFullInfo{
    VariableDebtToken: reserve.VariableDebtTokenAddress,
    DebtBalance:       balance,
    DebtBalanceInBase: balanceInBase,
    Price:             price,
    AssetUnit:         unit,
  }, nil
}

func findPosition(positions []UserReservePosition, asset common.Address) (UserReservePosition, bool) {
  for _, pos := range positions {
    if pos.Asset == asset {
      return pos, true
    }
  }
  return UserReservePosition{}, false
}

func reserveEnabledOnBitmap(bitmap *big.Int, reserveID uint16) bool {
  if bitmap == nil {
    return false
  }
  return bitmap.Bit(int(reserveID)) == 1
}

func canLiquidateHealthFactor(healthFactor *big.Int, liquidationAllowed bool) bool {
  if healthFactor.Cmp(healthFactorLiquidationThreshold) >= 0 {
    return false
  }

  if !liquidationAllowed && healthFactor.Cmp(minHealthFactorLiquidationThreshold) >= 0 {
    return false
  }

  return true
}

func isReserveReadyForLiquidations(reserve *ReserveData, isCollateral bool, currentTimestamp uint64) bool {
  areLiquidationsAllowed := reserve.LiquidationGracePeriodUntil < currentTimestamp
  collateralEnabled := true
  if isCollateral {
    collateralEnabled = reserve.Config.LiquidationThreshold != 0
  }

  return reserve.Config.IsActive && !reserve.Config.IsPaused && areLiquidationsAllowed && collateralEnabled
}

func isCollateralEnabledForUser(user *UserData, collateralReserve *ReserveData) bool {
  for _, pos := range user.Positions {
    if pos.Asset == collateralReserve.Asset && pos.UseAsCollateral {
      return true
    }
  }
  return false
}

func getLiquidationBonus(input *SimulationInput, collateralReserve *ReserveData) *big.Int {
  if input.User.UserEModeCategory != 0 {
    if category, ok := input.Global.EModeCategories[input.User.UserEModeCategory]; ok &&
      reserveEnabledOnBitmap(category.CollateralBitmap, collateralReserve.ID) {
      return bigFromUint64(category.LiquidationBonus)
    }
  }
  return bigFromUint64(collateralReserve.Config.LiquidationBonus)
}

func getMaxDebtToLiquidate(
  userInfo *UserPositionFullInfo,
  collateralInfo *CollateralFullInfo,
  debtInfo *DebtFullInfo,
  debtLiquidationAmount *big.Int,
) *big.Int {
  maxDebtToLiquidate := cloneBig(debtInfo.DebtBalance)

  if collateralInfo.CollateralBalanceInBase.Cmp(minBaseMaxCloseFactorThreshold) >= 0 &&
    debtInfo.DebtBalanceInBase.Cmp(minBaseMaxCloseFactorThreshold) >= 0 &&
    userInfo.HealthFactor.Cmp(closeFactorHFThreshold) > 0 {
    totalDefaultLiquidatableDebtInBase := percentMul(userInfo.TotalDebtBase, defaultLiquidationCloseFactor)
    if debtInfo.DebtBalanceInBase.Cmp(totalDefaultLiquidatableDebtInBase) > 0 {
      maxDebtToLiquidate = div(
        mul(totalDefaultLiquidatableDebtInBase, debtInfo.AssetUnit),
        debtInfo.Price,
      )
    }
  }

  return minBig(maxDebtToLiquidate, debtLiquidationAmount)
}

func getAvailableCollateralAndDebtToLiquidate(
  maxDebtToLiquidate *big.Int,
  liquidationBonus *big.Int,
  collateralInfo *CollateralFullInfo,
  debtInfo *DebtFullInfo,
  collateralConfig *ReserveConfig,
) (*big.Int, *big.Int, *big.Int) {
  liquidationProtocolFeePercentage := bigFromUint64(collateralConfig.LiquidationProtocolFee)

  maxBaseCollateral := div(
    mul(mul(debtInfo.Price, maxDebtToLiquidate), collateralInfo.AssetUnit),
    mul(collateralInfo.Price, debtInfo.AssetUnit),
  )
  maxCollateralToLiquidate := percentMul(maxBaseCollateral, liquidationBonus)

  collateralAmountToLiquidate := big.NewInt(0)
  debtAmountToLiquidate := big.NewInt(0)
  if maxCollateralToLiquidate.Cmp(collateralInfo.CollateralBalance) > 0 {
    collateralAmountToLiquidate = cloneBig(collateralInfo.CollateralBalance)
    debtAmountToLiquidate = percentDiv(
      div(
        mul(mul(collateralInfo.Price, collateralAmountToLiquidate), debtInfo.AssetUnit),
        mul(debtInfo.Price, collateralInfo.AssetUnit),
      ),
      liquidationBonus,
    )
  } else {
    collateralAmountToLiquidate = maxCollateralToLiquidate
    debtAmountToLiquidate = cloneBig(maxDebtToLiquidate)
  }

  liquidationProtocolFee := big.NewInt(0)
  if !isZero(liquidationProtocolFeePercentage) {
    bonusCollateral := sub(collateralAmountToLiquidate, percentDiv(collateralAmountToLiquidate, liquidationBonus))
    liquidationProtocolFee = percentMul(bonusCollateral, liquidationProtocolFeePercentage)
    collateralAmountToLiquidate = sub(collateralAmountToLiquidate, liquidationProtocolFee)
  }

  return collateralAmountToLiquidate, debtAmountToLiquidate, liquidationProtocolFee
}

func adjustAmountsForGoodLeftovers(
  collateralAmountToLiquidate *big.Int,
  debtAmountToLiquidate *big.Int,
  liquidationProtocolFee *big.Int,
  liquidationBonus *big.Int,
  collateralInfo *CollateralFullInfo,
  debtInfo *DebtFullInfo,
  collateralConfig *ReserveConfig,
) (*big.Int, *big.Int, *big.Int) {
  if sub(add(collateralAmountToLiquidate, liquidationProtocolFee), collateralInfo.CollateralBalance).Sign() < 0 &&
    sub(debtAmountToLiquidate, debtInfo.DebtBalance).Sign() < 0 {
    collateralLeftoverInBaseCurrency := div(
      mul(sub(sub(collateralInfo.CollateralBalance, collateralAmountToLiquidate), liquidationProtocolFee), collateralInfo.Price),
      collateralInfo.AssetUnit,
    )
    debtLeftoverInBaseCurrency := div(
      mul(sub(debtInfo.DebtBalance, debtAmountToLiquidate), debtInfo.Price),
      debtInfo.AssetUnit,
    )

    if collateralLeftoverInBaseCurrency.Cmp(minLeftoverBase) < 0 ||
      debtLeftoverInBaseCurrency.Cmp(minLeftoverBase) < 0 {
      collateralDecreaseAmountInBaseCurrency := big.NewInt(0)
      debtDecreaseAmountInBaseCurrency := big.NewInt(0)

      if collateralLeftoverInBaseCurrency.Cmp(minLeftoverBase) < 0 {
        collateralDecreaseAmountInBaseCurrency = sub(minLeftoverBase, collateralLeftoverInBaseCurrency)
      }
      if debtLeftoverInBaseCurrency.Cmp(minLeftoverBase) < 0 {
        debtDecreaseAmountInBaseCurrency = sub(minLeftoverBase, debtLeftoverInBaseCurrency)
      }

      if collateralDecreaseAmountInBaseCurrency.Cmp(debtDecreaseAmountInBaseCurrency) > 0 {
        collateralDecreaseAmount := div(
          mul(collateralDecreaseAmountInBaseCurrency, collateralInfo.AssetUnit),
          collateralInfo.Price,
        )
        collateralAmountToLiquidate = sub(collateralAmountToLiquidate, collateralDecreaseAmount)
        debtAmountToLiquidate = percentDiv(
          div(
            mul(mul(collateralInfo.Price, collateralAmountToLiquidate), debtInfo.AssetUnit),
            mul(debtInfo.Price, collateralInfo.AssetUnit),
          ),
          liquidationBonus,
        )
      } else {
        debtDecreaseAmount := div(
          mul(debtDecreaseAmountInBaseCurrency, debtInfo.AssetUnit),
          debtInfo.Price,
        )
        debtAmountToLiquidate = sub(debtAmountToLiquidate, debtDecreaseAmount)
        collateralAmountToLiquidate = percentMul(
          div(
            mul(mul(debtInfo.Price, debtAmountToLiquidate), collateralInfo.AssetUnit),
            mul(collateralInfo.Price, debtInfo.AssetUnit),
          ),
          liquidationBonus,
        )
      }

      liquidationProtocolFeePercentage := bigFromUint64(collateralConfig.LiquidationProtocolFee)
      if !isZero(liquidationProtocolFeePercentage) {
        bonusCollateral := sub(collateralAmountToLiquidate, percentDiv(collateralAmountToLiquidate, liquidationBonus))
        liquidationProtocolFee = percentMul(bonusCollateral, liquidationProtocolFeePercentage)
        collateralAmountToLiquidate = sub(collateralAmountToLiquidate, liquidationProtocolFee)
      }
    }
  }

  return collateralAmountToLiquidate, debtAmountToLiquidate, liquidationProtocolFee
}
