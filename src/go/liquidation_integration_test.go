package liquidation

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"aavev3origin/liquidation/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	defaultFullnodeEndpoint          = "https://open-platform.nodereal.io/5d9c218e356942a6a9c577e2aadd174c/base"
	l2PoolAddress                    = "0xA238Dd80C259a72e81d7e4664a9801593F98d1c5"
	aaveOracleAddress                = "0x2Cc0Fc26eD4563A5ce5e8bdcfe1A2878676Ae156"
	liquidationProviderAddress       = "0xcf03ed89c297e54935bd0bb6ea33d16abca8ed94"
	maxEModeCategoryID         uint8 = 127
	reserveConfigBitSize       uint  = 256
)

func TestIntegrationMatchesOnchainResults(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL())
	if err != nil {
		t.Fatalf("failed to connect rpc: %v", err)
	}

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		t.Fatalf("failed to load latest header: %v", err)
	}

	blockNumber := new(big.Int).Set(header.Number)
	callOpts := &bind.CallOpts{
		BlockNumber: blockNumber,
		Context:     ctx,
	}

	pool, err := contracts.NewL2Pool(common.HexToAddress(l2PoolAddress), client)
	if err != nil {
		t.Fatalf("failed to init l2 pool: %v", err)
	}
	oracle, err := contracts.NewAaveOracle(common.HexToAddress(aaveOracleAddress), client)
	if err != nil {
		t.Fatalf("failed to init oracle: %v", err)
	}
	liquidationProvider, err := contracts.NewLiquidationProvider(common.HexToAddress(liquidationProviderAddress), client)
	if err != nil {
		t.Fatalf("failed to init liquidation provider: %v", err)
	}

	globalConfig, err := buildGlobalConfig(pool, callOpts, header.Time)
	if err != nil {
		t.Fatalf("failed to build global config: %v", err)
	}
	reserves, err := buildReserves(ctx, client, pool, oracle, callOpts)
	if err != nil {
		t.Fatalf("failed to build reserves: %v", err)
	}

	users := selectTestUsers(t)
	for _, user := range users {
		user := user
		t.Run(strings.ToLower(user.Hex()), func(t *testing.T) {
			input, collateralAsset, debtAsset, err := buildSimulationInputForUser(
				ctx,
				client,
				pool,
				reserves,
				globalConfig,
				callOpts,
				user,
			)
			if err != nil {
				t.Fatalf("failed to build input: %v", err)
			}

			onchainUserInfo, err := liquidationProvider.GetUserPositionFullInfo(callOpts, user)
			if err != nil {
				t.Fatalf("failed to fetch onchain user info: %v", err)
			}
			localUserInfo, err := GetUserPositionFullInfo(input)
			if err != nil {
				t.Fatalf("failed to compute local user info: %v", err)
			}
			assertUserInfoMatch(t, onchainUserInfo, localUserInfo)

			onchainLiq, err := liquidationProvider.GetLiquidationInfo(callOpts, user, collateralAsset, debtAsset)
			if err != nil {
				t.Fatalf("failed to fetch onchain liquidation info: %v", err)
			}
			localLiq, err := GetLiquidationInfo(input, collateralAsset, debtAsset)
			if err != nil {
				t.Fatalf("failed to compute local liquidation info: %v", err)
			}
			assertLiquidationInfoMatch(t, onchainLiq, localLiq)
		})
	}
}

func buildGlobalConfig(pool *contracts.L2Pool, callOpts *bind.CallOpts, timestamp uint64) (*GlobalConfig, error) {
	categories, err := fetchEModeCategories(pool, callOpts)
	if err != nil {
		return nil, err
	}
	return &GlobalConfig{
		CurrentTimestamp:   timestamp,
		LiquidationAllowed: true,
		EModeCategories:    categories,
	}, nil
}

func buildReserves(
	ctx context.Context,
	client *ethclient.Client,
	pool *contracts.L2Pool,
	oracle *contracts.AaveOracle,
	callOpts *bind.CallOpts,
) (map[common.Address]*ReserveData, error) {
	reserves := make(map[common.Address]*ReserveData)

	assets, err := pool.GetReservesList(callOpts)
	if err != nil {
		return nil, fmt.Errorf("get reserves list: %w", err)
	}

	prices, err := oracle.GetAssetsPrices(callOpts, assets)
	if err != nil {
		return nil, fmt.Errorf("get prices: %w", err)
	}
	if len(prices) != len(assets) {
		return nil, fmt.Errorf("price length mismatch, assets %d prices %d", len(assets), len(prices))
	}

	for i, asset := range assets {
		config, err := pool.GetConfiguration(callOpts, asset)
		if err != nil {
			return nil, fmt.Errorf("get configuration for %s: %w", asset.Hex(), err)
		}
		reserveData, err := pool.GetReserveData(callOpts, asset)
		if err != nil {
			return nil, fmt.Errorf("get reserve data for %s: %w", asset.Hex(), err)
		}
		normalizedIncome, err := pool.GetReserveNormalizedIncome(callOpts, asset)
		if err != nil {
			return nil, fmt.Errorf("get normalized income for %s: %w", asset.Hex(), err)
		}
		normalizedDebt, err := pool.GetReserveNormalizedVariableDebt(callOpts, asset)
		if err != nil {
			return nil, fmt.Errorf("get normalized debt for %s: %w", asset.Hex(), err)
		}
		liquidationGracePeriod, err := pool.GetLiquidationGracePeriod(callOpts, asset)
		if err != nil {
			return nil, fmt.Errorf("get liquidation grace period for %s: %w", asset.Hex(), err)
		}

		erc20, err := contracts.NewERC20(asset, client)
		if err != nil {
			return nil, fmt.Errorf("new erc20 %s: %w", asset.Hex(), err)
		}
		decimals, err := erc20.Decimals(callOpts)
		if err != nil {
			return nil, fmt.Errorf("decimals for %s: %w", asset.Hex(), err)
		}

		reserves[asset] = &ReserveData{
			Asset: asset,
			ID:    reserveData.Id,
			Config: ReserveConfig{
				Ltv:                    extractBits(config.Data, 0, 16),
				LiquidationThreshold:   extractBits(config.Data, 16, 16),
				LiquidationBonus:       extractBits(config.Data, 32, 16),
				LiquidationProtocolFee: extractBits(config.Data, 152, 16),
				Decimals:               decimals,
				IsActive:               config.Data.Bit(56) == 1,
				IsPaused:               config.Data.Bit(60) == 1,
			},
			LiquidationGracePeriodUntil: liquidationGracePeriod.Uint64(),
			PriceInBaseCurrency:         cloneBig(prices[i]),
			ATokenAddress:               reserveData.ATokenAddress.Hex(),
			VariableDebtTokenAddress:    reserveData.VariableDebtTokenAddress.Hex(),
			NormalizedIncome:            cloneBig(normalizedIncome),
			NormalizedDebt:              cloneBig(normalizedDebt),
		}
	}

	return reserves, nil
}

func buildSimulationInputForUser(
	ctx context.Context,
	client *ethclient.Client,
	pool *contracts.L2Pool,
	reserves map[common.Address]*ReserveData,
	global *GlobalConfig,
	callOpts *bind.CallOpts,
	user common.Address,
) (*SimulationInput, common.Address, common.Address, error) {
	userConfig, err := pool.GetUserConfiguration(callOpts, user)
	if err != nil {
		return nil, common.Address{}, common.Address{}, fmt.Errorf("get user config: %w", err)
	}
	userEMode, err := pool.GetUserEMode(callOpts, user)
	if err != nil {
		return nil, common.Address{}, common.Address{}, fmt.Errorf("get user emode: %w", err)
	}

	positions := make([]UserReservePosition, 0, len(reserves))
	collateralCandidates := make([]common.Address, 0)
	debtCandidates := make([]common.Address, 0)

	for _, reserve := range reserves {
		position := UserReservePosition{
			Asset:                   reserve.Asset,
			UseAsCollateral:         isUsingAsCollateral(userConfig.Data, reserve.ID),
			IsBorrowing:             isBorrowing(userConfig.Data, reserve.ID),
			CollateralScaledBalance: big.NewInt(0),
			DebtScaledBalance:       big.NewInt(0),
		}

		if position.UseAsCollateral {
			amount, err := scaledBalance(ctx, client, common.HexToAddress(reserve.ATokenAddress), user, callOpts)
			if err != nil {
				return nil, common.Address{}, common.Address{}, fmt.Errorf("get collateral balance for %s: %w", reserve.Asset.Hex(), err)
			}
			position.CollateralScaledBalance = amount
			if amount.Sign() > 0 {
				collateralCandidates = append(collateralCandidates, reserve.Asset)
			}
		}

		if position.IsBorrowing {
			amount, err := scaledBalance(ctx, client, common.HexToAddress(reserve.VariableDebtTokenAddress), user, callOpts)
			if err != nil {
				return nil, common.Address{}, common.Address{}, fmt.Errorf("get debt balance for %s: %w", reserve.Asset.Hex(), err)
			}
			position.DebtScaledBalance = amount
			if amount.Sign() > 0 {
				debtCandidates = append(debtCandidates, reserve.Asset)
			}
		}

		positions = append(positions, position)
	}

	collateralAsset := pickFirst(collateralCandidates)
	debtAsset := pickDistinct(debtCandidates, collateralAsset)
	if collateralAsset == (common.Address{}) || debtAsset == (common.Address{}) {
		return nil, common.Address{}, common.Address{}, fmt.Errorf("unable to find collateral/debt assets with balances")
	}

	userData := &UserData{
		Address:           user,
		UserEModeCategory: uint8(userEMode.Uint64()),
		Positions:         positions,
	}

	return &SimulationInput{
		Global:   global,
		Reserves: reserves,
		User:     userData,
	}, collateralAsset, debtAsset, nil
}

func scaledBalance(
	ctx context.Context,
	client *ethclient.Client,
	token common.Address,
	user common.Address,
	callOpts *bind.CallOpts,
) (*big.Int, error) {
	reader, err := contracts.NewIScaledBalanceToken(token, client)
	if err != nil {
		return nil, err
	}
	return reader.ScaledBalanceOf(callOpts, user)
}

func extractBits(data *big.Int, offset, length uint) uint64 {
	if data == nil || offset+length > reserveConfigBitSize {
		return 0
	}
	mask := new(big.Int).Lsh(big.NewInt(1), length)
	mask.Sub(mask, big.NewInt(1))
	value := new(big.Int).And(new(big.Int).Rsh(data, offset), mask)
	return value.Uint64()
}

func isBorrowing(data *big.Int, reserveID uint16) bool {
	if data == nil {
		return false
	}
	return data.Bit(int(reserveID)<<1) == 1
}

func isUsingAsCollateral(data *big.Int, reserveID uint16) bool {
	if data == nil {
		return false
	}
	return data.Bit((int(reserveID)<<1)+1) == 1
}

func fetchEModeCategories(pool *contracts.L2Pool, callOpts *bind.CallOpts) (map[uint8]EModeCategory, error) {
	categories := make(map[uint8]EModeCategory)
	for id := uint8(0); id <= maxEModeCategoryID; id++ {
		data, err := pool.GetEModeCategoryData(callOpts, id)
		if err != nil {
			return nil, fmt.Errorf("get emode category %d: %w", id, err)
		}
		isEmpty := data.Ltv == 0 && data.LiquidationThreshold == 0 && data.LiquidationBonus == 0
		if isEmpty && id != 0 && len(categories) > 0 {
			break
		}
		if isEmpty {
			continue
		}

		bitmap, err := pool.GetEModeCategoryCollateralBitmap(callOpts, id)
		if err != nil {
			return nil, fmt.Errorf("get emode bitmap %d: %w", id, err)
		}
		categories[id] = EModeCategory{
			Ltv:                  uint64(data.Ltv),
			LiquidationThreshold: uint64(data.LiquidationThreshold),
			LiquidationBonus:     uint64(data.LiquidationBonus),
			CollateralBitmap:     cloneBig(bitmap),
		}
	}

	return categories, nil
}

func pickFirst(addrs []common.Address) common.Address {
	if len(addrs) == 0 {
		return common.Address{}
	}
	return addrs[0]
}

func pickDistinct(candidates []common.Address, avoid common.Address) common.Address {
	for _, addr := range candidates {
		if addr != avoid {
			return addr
		}
	}
	return pickFirst(candidates)
}

func selectTestUsers(t *testing.T) []common.Address {
	t.Helper()

	rawUsers := loadUserList(t)
	count := targetUserCount()
	if count > len(rawUsers) {
		count = len(rawUsers)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(rawUsers))

	selected := make([]common.Address, 0, count)
	for i := 0; i < count; i++ {
		selected = append(selected, rawUsers[perm[i]])
	}
	return selected
}

func loadUserList(t *testing.T) []common.Address {
	t.Helper()
	content, err := os.ReadFile(filepath.Join(".", "test-user.txt"))
	if err != nil {
		t.Fatalf("failed to read test-user.txt: %v", err)
	}

	lines := strings.Fields(string(content))
	users := make([]common.Address, 0, len(lines))
	for _, line := range lines {
		users = append(users, common.HexToAddress(strings.TrimSpace(line)))
	}
	if len(users) == 0 {
		t.Fatalf("no users found in test-user.txt")
	}
	return users
}

func targetUserCount() int {
	if raw := os.Getenv("TEST_USER_COUNT"); raw != "" {
		if v, err := strconv.Atoi(raw); err == nil && v > 0 {
			return v
		}
	}
	return 1
}

func rpcURL() string {
	if v := os.Getenv("FULLNODE_ENDPOINT"); v != "" {
		return v
	}
	if v := os.Getenv("RPC_URL"); v != "" {
		return v
	}
	return defaultFullnodeEndpoint
}

func assertUserInfoMatch(t *testing.T, onchain contracts.ILiquidationDataProviderUserPositionFullInfo, local *UserPositionFullInfo) {
	t.Helper()

	compareBig(t, "total collateral", onchain.TotalCollateralInBaseCurrency, local.TotalCollateralBase)
	compareBig(t, "total debt", onchain.TotalDebtInBaseCurrency, local.TotalDebtBase)
	compareBig(t, "current liquidation threshold", onchain.CurrentLiquidationThreshold, local.CurrentLiquidationThreshold)
	compareBig(t, "ltv", onchain.Ltv, local.Ltv)
	compareBig(t, "health factor", onchain.HealthFactor, local.HealthFactor)
}

func assertLiquidationInfoMatch(
	t *testing.T,
	onchain contracts.ILiquidationDataProviderLiquidationInfo,
	local *LiquidationInfo,
) {
	t.Helper()

	assertUserInfoMatch(t, onchain.UserInfo, local.UserInfo)

	if !strings.EqualFold(onchain.CollateralInfo.AToken.Hex(), local.CollateralInfo.AToken) {
		t.Fatalf("collateral aToken mismatch: onchain %s local %s", onchain.CollateralInfo.AToken.Hex(), local.CollateralInfo.AToken)
	}
	if !strings.EqualFold(onchain.DebtInfo.VariableDebtToken.Hex(), local.DebtInfo.VariableDebtToken) {
		t.Fatalf("variable debt token mismatch: onchain %s local %s", onchain.DebtInfo.VariableDebtToken.Hex(), local.DebtInfo.VariableDebtToken)
	}

	compareBig(t, "collateral balance", onchain.CollateralInfo.CollateralBalance, local.CollateralInfo.CollateralBalance)
	compareBig(t, "collateral balance in base", onchain.CollateralInfo.CollateralBalanceInBaseCurrency, local.CollateralInfo.CollateralBalanceInBase)
	compareBig(t, "collateral price", onchain.CollateralInfo.Price, local.CollateralInfo.Price)
	compareBig(t, "collateral asset unit", onchain.CollateralInfo.AssetUnit, local.CollateralInfo.AssetUnit)

	compareBig(t, "debt balance", onchain.DebtInfo.DebtBalance, local.DebtInfo.DebtBalance)
	compareBig(t, "debt balance in base", onchain.DebtInfo.DebtBalanceInBaseCurrency, local.DebtInfo.DebtBalanceInBase)
	compareBig(t, "debt price", onchain.DebtInfo.Price, local.DebtInfo.Price)
	compareBig(t, "debt asset unit", onchain.DebtInfo.AssetUnit, local.DebtInfo.AssetUnit)

	compareBig(t, "max collateral to liquidate", onchain.MaxCollateralToLiquidate, local.MaxCollateralToLiquidate)
	compareBig(t, "max debt to liquidate", onchain.MaxDebtToLiquidate, local.MaxDebtToLiquidate)
	compareBig(t, "liquidation protocol fee", onchain.LiquidationProtocolFee, local.LiquidationProtocolFee)
	compareBig(t, "amount to pass to liquidation call", onchain.AmountToPassToLiquidationCall, local.AmountToPassToLiquidationCall)
}

func compareBig(t *testing.T, label string, onchain *big.Int, local *big.Int) {
	t.Helper()

	if normalize(onchain).Cmp(normalize(local)) != 0 {
		t.Fatalf("%s mismatch: onchain %s local %s", label, onchain.String(), local.String())
	}
}

func normalize(v *big.Int) *big.Int {
	if v == nil {
		return big.NewInt(0)
	}
	return v
}
