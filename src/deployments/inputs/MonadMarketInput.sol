// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import './MarketInput.sol';

/**
 * @title MonadMarketInput
 * @notice Deployment configuration for Aave V3 on Monad mainnet.
 *
 * Chain:        Monad (Chain ID 143)
 * EVM version:  Cancun-compatible
 * Native token: MON
 * WMON:         0x3bd359c1119da7da1d913d1c4d2b7c461115433a
 * Explorer:     https://monadscan.com
 * RPC:          https://rpc.monad.xyz
 *
 * Monad is an independent L1 — no sequencer uptime feed or PriceOracleSentinel needed.
 *
 * Addresses to confirm before mainnet deployment:
 *   MON_USD_FEED  — verify at https://data.chain.link/feeds/monad
 *   WMON          — verify at https://monadscan.com
 *   ETH_USD_FEED  — already verified from BGD infrastructure report (March 2026)
 */
contract MonadMarketInput is MarketInput {
  // WMON — Wrapped native MON token
  address constant WMON = 0x3bd359C1119dA7Da1D913D1C4D2B7c461115433A;

  // Chainlink ETH/USD price feed on Monad — used as the market reference currency.
  // Source: BGD Aave <> Monad Infrastructure / Technical Evaluation, March 2026.
  // Verify at: https://data.chain.link/feeds/monad
  address constant ETH_USD_FEED = 0x1B1414782B859871781bA3E4B0979b9ca57A0A04;

  // MON/USD price feed — native network base token price.
  // TODO: confirm address at https://data.chain.link/feeds/monad before mainnet deployment.
  // Leaving as address(0) skips UiPoolDataProvider deployment (non-critical for core protocol).
  address constant MON_USD_FEED = 0xBcD78f76005B7515837af6b50c7C52BCf73822fb;

  function _getMarketInput(
    address deployer
  )
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
    // All administrative roles are assigned to the deployer for initial deployment.
    // Transfer marketOwner and poolAdmin to governance multisig after deployment.
    roles.marketOwner = deployer;
    roles.emergencyAdmin = deployer;
    roles.poolAdmin = deployer;

    config.marketId = 'Aave V3 Monad Market';

    // providerId must be unique across all Aave markets registered in
    // PoolAddressesProviderRegistry. Using chain ID 143 as a natural choice.
    config.providerId = 143;

    config.oracleDecimals = 8; // Chainlink standard
    config.flashLoanPremium = 0.0005e4; // 0.05%

    // Monad is an independent L1 — no sequencer uptime feed.
    flags.l2 = false;
    config.l2SequencerUptimeFeed = address(0);
    config.l2PriceOracleSentinelGracePeriod = 0;

    config.wrappedNativeToken = WMON;

    // ETH/USD — standard market reference currency shared across Aave markets.
    config.marketReferenceCurrencyPriceInUsdProxyAggregator = ETH_USD_FEED;

    // MON/USD — network base token price for UiPoolDataProvider.
    // Set to address(0) until the feed address is confirmed on-chain.
    config.networkBaseTokenPriceInUsdProxyAggregator = MON_USD_FEED;

    // Deploy a fresh treasury and rewards controller (no existing contracts to reuse).
    config.treasury = address(0);
    config.incentivesProxy = address(0);

    // bytes32(0) → CREATE2 with zero salt; produces deterministic addresses.
    config.salt = bytes32(0);

    return (roles, config, flags, deployedContracts);
  }
}
