// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./MarketInput.sol";

/**
 * @title K613ArbitrumMarketInput
 * @notice Market input configuration for Arbitrum Sepolia L2 testnet
 * @dev This configuration includes:
 * - L2 Sequencer Uptime Feed for Chainlink oracle validation
 * - WETH address for Arbitrum Sepolia
 * - Market ID and Provider ID for Arbitrum Sepolia
 * - Grace period for Price Oracle Sentinel
 */
contract K613ArbitrumMarketInput is MarketInput {
    // Arbitrum Sepolia Testnet addresses
    // L2 Sequencer Uptime Feed (Chainlink) - Arbitrum Sepolia
    address private constant ARBITRUM_SEPOLIA_SEQUENCER_UPTIME_FEED = 0x4Da69f028A5790fCcFeE3E3506357863c3c4ecc3;

    // WETH on Arbitrum Sepolia
    address private constant ARBITRUM_SEPOLIA_WETH = 0x980B62Da83eFf3D4576C647993b0c1D7faf17c73;

    // Grace period for Price Oracle Sentinel (3600 seconds = 1 hour)
    uint256 private constant GRACE_PERIOD = 3600;

    // Chainlink ETH/USD Price Feed for Arbitrum Sepolia
    // Standard Aave fallback: ETH/USD price feed
    // This will be used as fallback oracle for AaveOracle (FallbackOracle contract)
    // Fallback oracle is used when:
    // - Asset source is not set (address(0))
    // - Chainlink aggregator returns invalid price (<= 0)
    // This is the standard Aave approach - no custom fallback needed for MVP
    // Source: Chainlink Data Feeds for Arbitrum Sepolia
    // https://docs.chain.link/data-feeds/price-feeds/addresses
    address private constant ARBITRUM_SEPOLIA_ETH_USD_PRICE_FEED = 0x2d3bBa5e0A9Fd8EAa45Dcf71A2389b7C12005b1f;

    function _getMarketInput(address deployer)
        internal
        pure
        override
        returns (
            Roles memory roles,
            MarketConfig memory config,
            DeployFlags memory flags,
            MarketReport memory deployedContracts
        )
    {
        // Setup roles
        // For MVP/Testnet: All roles can be the same address (deployer)
        // For Production: Roles should be separated:
        //   - POOL_ADMIN: Main admin for pool configuration
        //   - EMERGENCY_ADMIN: Emergency pause/unpause
        //   - RISK_ADMIN: Risk parameter updates (set via PostDeploySetup.sol)
        //   - Market Owner: Ownership of PoolAddressesProvider
        roles.marketOwner = deployer;
        roles.emergencyAdmin = deployer;
        roles.poolAdmin = deployer;

        // Market configuration for Arbitrum Sepolia
        config.marketId = "K613 Aave v3 Fork ";
        config.providerId = 421614; // Arbitrum Sepolia chain ID
        config.oracleDecimals = 8;
        config.flashLoanPremiumTotal = 0.0005e4;
        config.flashLoanPremiumToProtocol = 0.0002e4;

        // L2 specific configuration
        config.l2SequencerUptimeFeed = ARBITRUM_SEPOLIA_SEQUENCER_UPTIME_FEED;
        config.l2PriceOracleSentinelGracePeriod = GRACE_PERIOD;
        config.wrappedNativeToken = ARBITRUM_SEPOLIA_WETH;

        // Set ETH/USD price feed for fallback oracle (standard Aave approach)
        // This will be used in FallbackOracle contract deployed during AaveOracle setup
        // FallbackOracle returns ETH/USD price when:
        // - Asset source is not set (address(0))
        // - Chainlink aggregator returns invalid price (<= 0)
        // This is the standard Aave approach - no custom fallback needed for MVP
        config.networkBaseTokenPriceInUsdProxyAggregator = ARBITRUM_SEPOLIA_ETH_USD_PRICE_FEED;
        config.marketReferenceCurrencyPriceInUsdProxyAggregator = ARBITRUM_SEPOLIA_ETH_USD_PRICE_FEED;

        // Set L2 flag
        flags.l2 = true;

        return (roles, config, flags, deployedContracts);
    }
}
