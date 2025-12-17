// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import 'forge-std/Script.sol';
import 'forge-std/console.sol';

import {AaveV3BatchOrchestration} from '../src/deployments/projects/aave-v3-batched/AaveV3BatchOrchestration.sol';
import {IAaveV3ConfigEngine} from '../src/contracts/extensions/v3-config-engine/IAaveV3ConfigEngine.sol';
import {EngineFlags} from '../src/contracts/extensions/v3-config-engine/EngineFlags.sol';
import {MockAggregatorSetPrice} from '../tests/invariants/utils/mocks/MockAggregatorSetPrice.sol';
import {SequencerOracle} from '../src/contracts/mocks/oracle/SequencerOracle.sol';
import {TestnetERC20} from '../src/contracts/mocks/testnet-helpers/TestnetERC20.sol';
import {IERC20} from '../src/contracts/dependencies/openzeppelin/contracts/IERC20.sol';
import {IPool} from '../src/contracts/interfaces/IPool.sol';
import {IACLManager} from '../src/contracts/interfaces/IACLManager.sol';
import {IPoolAddressesProvider} from '../src/contracts/interfaces/IPoolAddressesProvider.sol';
import {LiquidationDataProvider} from '../src/contracts/helpers/LiquidationDataProvider.sol';
import {MarketReport, MarketConfig, DeployFlags, Roles} from '../src/deployments/interfaces/IMarketReportTypes.sol';

struct DeployedAssets {
  TestnetERC20 usdc;
  TestnetERC20 usdt;
  TestnetERC20 wbtc;
  TestnetERC20 weth;
  MockAggregatorSetPrice usdcOracle;
  MockAggregatorSetPrice usdtOracle;
  MockAggregatorSetPrice wbtcOracle;
  MockAggregatorSetPrice wethOracle;
  SequencerOracle sequencerOracle;
}

/**
 * @notice One-shot script to deploy a full Aave v3 market on Base Sepolia, together with
 *         4 mock assets and adjustable price feeds.
 * @dev Transactions are broadcasted; set PRIVATE_KEY env. RPC handled by forge args.
 */
contract DeployAaveV3BaseSepolia is Script {
  function run() external {
    uint256 deployerPk = _getPrivateKey();
    address deployer = vm.addr(deployerPk);

    vm.startBroadcast(deployerPk);

    // 1) Deploy test tokens
    DeployedAssets memory a;
    a.usdc = new TestnetERC20('USDC', 'USDC', 6, deployer);
    a.usdt = new TestnetERC20('USDT', 'USDT', 6, deployer);
    a.wbtc = new TestnetERC20('WBTC', 'WBTC', 8, deployer);
    a.weth = new TestnetERC20('WETH', 'WETH', 18, deployer);

    // Mint 1e9 units (before decimals) to deployer
    uint256 baseSupply = 1e9;
    a.usdc.mint(deployer, baseSupply * 10 ** a.usdc.decimals());
    a.usdt.mint(deployer, baseSupply * 10 ** a.usdt.decimals());
    a.wbtc.mint(deployer, baseSupply * 10 ** a.wbtc.decimals());
    a.weth.mint(deployer, baseSupply * 10 ** a.weth.decimals());

    // 2) Deploy adjustable price feeds (8 decimals)
    a.usdcOracle = new MockAggregatorSetPrice(1e8); // $1
    a.usdtOracle = new MockAggregatorSetPrice(1e8); // $1
    a.wbtcOracle = new MockAggregatorSetPrice(88700e8); // $88,700
    a.wethOracle = new MockAggregatorSetPrice(3078e8); // $3,078

    // 3) Deploy sequencer uptime oracle (for sentinel) and mark as healthy
    a.sequencerOracle = new SequencerOracle(deployer);
    a.sequencerOracle.setAnswer(false, block.timestamp);

    // 4) Build config + deploy Aave V3 (L2Pool)
    Roles memory roles;
    roles.marketOwner = deployer;
    roles.poolAdmin = deployer;
    roles.emergencyAdmin = deployer;

    MarketConfig memory config;
    config.marketId = 'Aave V3 Base Sepolia';
    config.providerId = 84532;
    config.oracleDecimals = 8;
    config.flashLoanPremium = 0.0005e4; // 0.05%
    config.networkBaseTokenPriceInUsdProxyAggregator = address(a.wethOracle);
    config.marketReferenceCurrencyPriceInUsdProxyAggregator = address(a.usdcOracle);
    config.l2SequencerUptimeFeed = address(a.sequencerOracle);
    config.l2PriceOracleSentinelGracePeriod = 1 hours;
    config.wrappedNativeToken = address(a.weth); // treated as plain ERC20

    DeployFlags memory flags;
    flags.l2 = true;

    MarketReport memory report;
    report = AaveV3BatchOrchestration.deployAaveV3(deployer, roles, config, flags, report);

    // 5) List assets via ConfigEngine with custom implementations
    IAaveV3ConfigEngine configEngine = IAaveV3ConfigEngine(report.configEngine);
    IAaveV3ConfigEngine.PoolContext memory context = IAaveV3ConfigEngine.PoolContext({
      networkName: 'Base Sepolia',
      networkAbbreviation: 'BSP'
    });

    IAaveV3ConfigEngine.ListingWithCustomImpl[] memory listings = new IAaveV3ConfigEngine
      .ListingWithCustomImpl[](4);

    IAaveV3ConfigEngine.InterestRateInputData memory rateParams = IAaveV3ConfigEngine
      .InterestRateInputData({
        optimalUsageRatio: 45_00,
        baseVariableBorrowRate: 0,
        variableRateSlope1: 4_00,
        variableRateSlope2: 60_00
      });

    listings[0] = _buildListing(
      address(a.usdc),
      'USDC',
      address(a.usdcOracle),
      rateParams,
      report.aToken,
      report.variableDebtToken
    );
    listings[1] = _buildListing(
      address(a.usdt),
      'USDT',
      address(a.usdtOracle),
      rateParams,
      report.aToken,
      report.variableDebtToken
    );
    listings[2] = _buildListing(
      address(a.wbtc),
      'WBTC',
      address(a.wbtcOracle),
      rateParams,
      report.aToken,
      report.variableDebtToken
    );
    listings[3] = _buildListing(
      address(a.weth),
      'WETH',
      address(a.wethOracle),
      rateParams,
      report.aToken,
      report.variableDebtToken
    );

    IACLManager(report.aclManager).addAssetListingAdmin(address(configEngine));
    IACLManager(report.aclManager).addRiskAdmin(address(configEngine));
    configEngine.listAssetsCustom(context, listings);

    LiquidationDataProvider liquidationDataProvider = new LiquidationDataProvider(
      report.poolProxy,
      report.poolAddressesProvider
    );

    // Log outputs for convenience
    console.log('Deployer', deployer);
    console.log('USDC', address(a.usdc));
    console.log('USDT', address(a.usdt));
    console.log('WBTC', address(a.wbtc));
    console.log('WETH', address(a.weth));
    console.log('Oracle USDC', address(a.usdcOracle));
    console.log('Oracle USDT', address(a.usdtOracle));
    console.log('Oracle WBTC', address(a.wbtcOracle));
    console.log('Oracle WETH', address(a.wethOracle));
    console.log('Sequencer oracle', address(a.sequencerOracle));
    console.log('Pool', report.poolProxy);
    console.log('ConfigEngine', report.configEngine);
    console.log('AaveOracle', report.aaveOracle);
    console.log('LiquidationDataProvider', address(liquidationDataProvider));

    // Ensure sequencer oracle timestamp is older than grace period to allow borrowing
    a.sequencerOracle.setAnswer(false, block.timestamp - 2 hours);

    _quickHealthCheck(report, deployer, a);

    vm.stopBroadcast();
  }

  function _buildListing(
    address asset,
    string memory symbol,
    address priceFeed,
    IAaveV3ConfigEngine.InterestRateInputData memory rateParams,
    address aTokenImpl,
    address vTokenImpl
  ) internal pure returns (IAaveV3ConfigEngine.ListingWithCustomImpl memory listing) {
    listing.base = IAaveV3ConfigEngine.Listing({
      asset: asset,
      assetSymbol: symbol,
      priceFeed: priceFeed,
      rateStrategyParams: rateParams,
      enabledToBorrow: EngineFlags.ENABLED,
      borrowableInIsolation: EngineFlags.DISABLED,
      withSiloedBorrowing: EngineFlags.DISABLED,
      flashloanable: EngineFlags.ENABLED,
      ltv: 82_50,
      liqThreshold: 86_00,
      liqBonus: 5_00,
      reserveFactor: 10_00,
      supplyCap: 0, // uncapped
      borrowCap: 0, // uncapped
      debtCeiling: 0,
      liqProtocolFee: 10_00
    });

    listing.implementations = IAaveV3ConfigEngine.TokenImplementations({
      aToken: aTokenImpl,
      vToken: vTokenImpl
    });
  }

  /// @dev Simple post-deploy sanity: supply/borrow/repay flow + price bump
  function _quickHealthCheck(
    MarketReport memory report,
    address deployer,
    DeployedAssets memory a
  ) internal {
    IPool pool = IPool(report.poolProxy);

    // Approvals
    IERC20(address(a.usdc)).approve(address(pool), type(uint256).max);
    IERC20(address(a.weth)).approve(address(pool), type(uint256).max);

    // a) Supply 1 WETH as collateral
    uint256 wethAmount = 1 ether;
    pool.supply(address(a.weth), wethAmount, deployer, 0);
    pool.setUserUseReserveAsCollateral(address(a.weth), true);

    // b) Supply 1000 USDC, keep collateral disabled
    uint256 usdcAmount = 1000 * 1e6;
    pool.supply(address(a.usdc), usdcAmount, deployer, 0);
    pool.setUserUseReserveAsCollateral(address(a.usdc), false);

    // c) Borrow 900 USDC (variable rate)
    uint256 borrowAmount = 900 * 1e6;
    pool.borrow(address(a.usdc), borrowAmount, 2, 0, deployer);

    // d) Repay all USDC (repay max to cover interest)
    pool.repay(address(a.usdc), type(uint256).max, 2, deployer);

    // e) Update WETH price to 3100 USD
    a.wethOracle.setLatestAnswer(3100e8);

    // Log account data
    (
      uint256 totalCollateralBase,
      uint256 totalDebtBase,
      uint256 availableBorrowsBase,
      uint256 currentLiquidationThreshold,
      uint256 ltv,
      uint256 healthFactor
    ) = pool.getUserAccountData(deployer);

    console.log('Post-check collateral (base)', totalCollateralBase);
    console.log('Post-check debt (base)', totalDebtBase);
    console.log('Available borrows (base)', availableBorrowsBase);
    console.log('Liq threshold', currentLiquidationThreshold);
    console.log('LTV', ltv);
    console.log('Health factor', healthFactor);
    console.log('WETH price now', a.wethOracle.latestAnswer());
    console.log('USDC price (ref currency)', a.usdcOracle.latestAnswer());
  }

  function _getPrivateKey() internal returns (uint256) {
    string memory raw = vm.envString('PRIVATE_KEY');
    bytes memory b = bytes(raw);
    bool hasPrefix = b.length >= 2 && b[0] == '0' && (b[1] == 'x' || b[1] == 'X');
    string memory normalized = hasPrefix ? raw : string(abi.encodePacked('0x', raw));
    return uint256(vm.parseBytes32(normalized));
  }
}
