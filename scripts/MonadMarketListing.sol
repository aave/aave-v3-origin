// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {AaveV3Payload} from '../src/contracts/extensions/v3-config-engine/AaveV3Payload.sol';
import {IAaveV3ConfigEngine as IEngine} from '../src/contracts/extensions/v3-config-engine/IAaveV3ConfigEngine.sol';
import {EngineFlags} from '../src/contracts/extensions/v3-config-engine/EngineFlags.sol';
import {ACLManager} from '../src/contracts/protocol/configuration/ACLManager.sol';

/**
 * @title MonadMarketListing
 * @notice Lists WMON and USDC on the freshly deployed Aave V3 Monad market.
 *
 * Parameters follow the Aave Monad activation plan:
 *   WMON — collateral-only at launch (LTV=0, borrowing disabled)
 *   USDC — borrowable stablecoin, no collateral use at launch
 *
 * The deployer must grant poolAdmin to this contract before calling execute(),
 * and this contract renounces it in _postExecute().
 */
contract MonadMarketListing is AaveV3Payload {
  // ── Addresses ────────────────────────────────────────────────────────────

  // WMON — Wrapped native MON token on Monad mainnet
  address public immutable WMON;

  // MON/USD price feed — use MockAggregator until confirmed on data.chain.link/feeds/monad
  address public immutable MON_USD_FEED;

  // USDC token — Circle CCTP v2 on Monad mainnet (proxy: 0x754704Bc059F8C67012fEd69BC8A327a5aafb603)
  address public immutable USDC;

  // USDC/USD Chainlink price feed on Monad
  // Source: BGD Aave <> Monad Infrastructure / Technical Evaluation, March 2026
  address public immutable USDC_USD_FEED;

  // Token implementations from the market deployment report
  address public immutable ATOKEN_IMPL;
  address public immutable VARIABLE_DEBT_IMPL;

  // ACLManager — to renounce poolAdmin in _postExecute
  ACLManager public immutable ACL_MANAGER;

  bytes32 constant POOL_ADMIN_ROLE =
    0x12ad05bde78c5ab75238ce885307f96ecd482bb402ef831f99e7018a0f169b7b;

  constructor(
    IEngine engine,
    address wmon,
    address monUsdFeed,
    address usdc,
    address usdcUsdFeed,
    address aTokenImpl,
    address variableDebtImpl,
    address aclManager
  ) AaveV3Payload(engine) {
    WMON = wmon;
    MON_USD_FEED = monUsdFeed;
    USDC = usdc;
    USDC_USD_FEED = usdcUsdFeed;
    ATOKEN_IMPL = aTokenImpl;
    VARIABLE_DEBT_IMPL = variableDebtImpl;
    ACL_MANAGER = ACLManager(aclManager);
  }

  // ── ConfigEngine listings ────────────────────────────────────────────────

  function newListingsCustom()
    public
    view
    override
    returns (IEngine.ListingWithCustomImpl[] memory)
  {
    IEngine.ListingWithCustomImpl[] memory listings = new IEngine.ListingWithCustomImpl[](2);

    // ── WMON ──────────────────────────────────────────────────────────────
    // Collateral-only at launch. No borrowing until chain matures.
    listings[0] = IEngine.ListingWithCustomImpl(
      IEngine.Listing({
        asset: WMON,
        assetSymbol: 'WMON',
        priceFeed: MON_USD_FEED,
        rateStrategyParams: IEngine.InterestRateInputData({
          optimalUsageRatio: 45_00,
          baseVariableBorrowRate: 0,
          variableRateSlope1: 7_00,
          variableRateSlope2: 300_00
        }),
        enabledToBorrow: EngineFlags.DISABLED,
        borrowableInIsolation: EngineFlags.DISABLED,
        withSiloedBorrowing: EngineFlags.DISABLED,
        flashloanable: EngineFlags.ENABLED,
        ltv: 0,
        liqThreshold: 0,
        liqBonus: 0,
        reserveFactor: 20_00,
        supplyCap: 500_000,
        borrowCap: 0,
        debtCeiling: 0,
        liqProtocolFee: 10_00
      }),
      IEngine.TokenImplementations({aToken: ATOKEN_IMPL, vToken: VARIABLE_DEBT_IMPL})
    );

    // ── USDC ──────────────────────────────────────────────────────────────
    // Borrowable stablecoin. No collateral use at launch.
    listings[1] = IEngine.ListingWithCustomImpl(
      IEngine.Listing({
        asset: USDC,
        assetSymbol: 'USDC',
        priceFeed: USDC_USD_FEED,
        rateStrategyParams: IEngine.InterestRateInputData({
          optimalUsageRatio: 90_00,
          baseVariableBorrowRate: 0,
          variableRateSlope1: 5_00,
          variableRateSlope2: 60_00
        }),
        enabledToBorrow: EngineFlags.ENABLED,
        borrowableInIsolation: EngineFlags.ENABLED,
        withSiloedBorrowing: EngineFlags.DISABLED,
        flashloanable: EngineFlags.ENABLED,
        ltv: 0,
        liqThreshold: 0,
        liqBonus: 0,
        reserveFactor: 10_00,
        supplyCap: 3_000_000,
        borrowCap: 2_700_000,
        debtCeiling: 0,
        liqProtocolFee: 10_00
      }),
      IEngine.TokenImplementations({aToken: ATOKEN_IMPL, vToken: VARIABLE_DEBT_IMPL})
    );

    return listings;
  }

  function getPoolContext() public pure override returns (IEngine.PoolContext memory) {
    return IEngine.PoolContext({networkName: 'Monad', networkAbbreviation: 'Mon'});
  }

  function _postExecute() internal override {
    // Renounce poolAdmin so this contract holds no lingering privileges
    ACL_MANAGER.renounceRole(POOL_ADMIN_ROLE, address(this));
  }
}
