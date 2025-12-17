package liquidation

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	percentageFactor                    = big.NewInt(10_000)
	halfPercentageFactor                = big.NewInt(5_000)
	wad                                 = big.NewInt(1_000_000_000_000_000_000)                                                              // 1e18
	halfWad                             = big.NewInt(500_000_000_000_000_000)                                                                // 0.5e18
	ray                                 = func() *big.Int { v, _ := new(big.Int).SetString("1000000000000000000000000000", 10); return v }() // 1e27
	halfRay                             = func() *big.Int { v, _ := new(big.Int).SetString("500000000000000000000000000", 10); return v }()  // 0.5e27
	defaultLiquidationCloseFactor       = big.NewInt(5_000)                                                                                  // 0.5e4
	closeFactorHFThreshold              = big.NewInt(950_000_000_000_000_000)                                                                // 0.95e18
	healthFactorLiquidationThreshold    = big.NewInt(1_000_000_000_000_000_000)
	minHealthFactorLiquidationThreshold = big.NewInt(950_000_000_000_000_000) // 0.95e18
	minBaseMaxCloseFactorThreshold      = big.NewInt(200_000_000_000)         // 2000e8
	minLeftoverBase                     = big.NewInt(100_000_000_000)         // MIN_BASE_MAX_CLOSE_FACTOR_THRESHOLD / 2
	maxUint256                          = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
)

// EModeCategory mirrors the on-chain eMode configuration.
type EModeCategory struct {
	Ltv                  uint64
	LiquidationThreshold uint64
	LiquidationBonus     uint64
	CollateralBitmap     *big.Int
}

// ReserveConfig contains the subset of reserve configuration bits required by the liquidation helpers.
type ReserveConfig struct {
	Ltv                    uint64
	LiquidationThreshold   uint64
	LiquidationBonus       uint64
	LiquidationProtocolFee uint64
	Decimals               uint8
	IsActive               bool
	IsPaused               bool
}

// ReserveData collects the static data needed for calculations.
type ReserveData struct {
	Asset                       common.Address
	ID                          uint16
	Config                      ReserveConfig
	LiquidationGracePeriodUntil uint64
	PriceInBaseCurrency         *big.Int
	ATokenAddress               string
	VariableDebtTokenAddress    string
	NormalizedIncome            *big.Int // ray
	NormalizedDebt              *big.Int // ray
}

// UserReservePosition tracks a user's balances and configuration for a specific reserve.
type UserReservePosition struct {
	Asset                   common.Address
	UseAsCollateral         bool
	IsBorrowing             bool
	CollateralScaledBalance *big.Int
	DebtScaledBalance       *big.Int
}

// UserData aggregates per-user state.
type UserData struct {
	Address           common.Address
	UserEModeCategory uint8
	Positions         []UserReservePosition
}

// GlobalConfig hosts configuration shared across calculations.
type GlobalConfig struct {
	CurrentTimestamp uint64

	// 来自 IPriceOracleSentinel(priceOracleSentinel).isLiquidationAllowed()
	LiquidationAllowed bool
	EModeCategories    map[uint8]EModeCategory
}

// SimulationInput is the full snapshot required to evaluate liquidation data.
type SimulationInput struct {
	Global   *GlobalConfig
	Reserves map[common.Address]*ReserveData
	User     *UserData
}

// UserPositionFullInfo matches ILiquidationDataProvider.UserPositionFullInfo.
type UserPositionFullInfo struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}

// CollateralFullInfo matches ILiquidationDataProvider.CollateralFullInfo.
type CollateralFullInfo struct {
	AToken                  string
	CollateralBalance       *big.Int
	CollateralBalanceInBase *big.Int
	Price                   *big.Int
	AssetUnit               *big.Int
}

// DebtFullInfo matches ILiquidationDataProvider.DebtFullInfo.
type DebtFullInfo struct {
	VariableDebtToken string
	DebtBalance       *big.Int
	DebtBalanceInBase *big.Int
	Price             *big.Int
	AssetUnit         *big.Int
}

// LiquidationInfo mirrors ILiquidationDataProvider.LiquidationInfo.
type LiquidationInfo struct {
	UserInfo                      *UserPositionFullInfo
	CollateralInfo                *CollateralFullInfo
	DebtInfo                      *DebtFullInfo
	MaxCollateralToLiquidate      *big.Int
	MaxDebtToLiquidate            *big.Int
	LiquidationProtocolFee        *big.Int
	AmountToPassToLiquidationCall *big.Int
}
