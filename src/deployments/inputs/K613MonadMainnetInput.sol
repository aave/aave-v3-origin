// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./MarketInput.sol";

/**
 * @title K613MonadMainnetInput
 * @notice Market input configuration for Monad mainnet
 */
contract K613MonadMainnetInput is MarketInput {
    // Monad mainnet wrapped native token (WMON)
    address private constant MONAD_WMON = 0x3bd359C1119dA7Da1D913D1C4D2B7c461115433A;

    // Chainlink ETH/USD feed on Monad mainnet
    address private constant MONAD_ETH_USD_PRICE_FEED = 0x1B1414782B859871781bA3E4B0979b9ca57A0A04;

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
        roles.marketOwner = deployer;
        roles.emergencyAdmin = deployer;
        roles.poolAdmin = deployer;

        config.marketId = "K613 Aave v3 Monad";
        config.providerId = 143;
        config.oracleDecimals = 8;
        config.flashLoanPremiumTotal = 0.0005e4;
        config.flashLoanPremiumToProtocol = 0.0002e4;

        // L1 network: no sequencer uptime feed / sentinel setup
        flags.l2 = false;

        config.wrappedNativeToken = MONAD_WMON;
        config.networkBaseTokenPriceInUsdProxyAggregator = MONAD_ETH_USD_PRICE_FEED;
        config.marketReferenceCurrencyPriceInUsdProxyAggregator = MONAD_ETH_USD_PRICE_FEED;

        return (roles, config, flags, deployedContracts);
    }
}
