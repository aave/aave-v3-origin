// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {DeployAaveV3MarketBatchedBase} from "./misc/DeployAaveV3MarketBatchedBase.sol";
import {K613ArbitrumMarketInput} from "../src/deployments/inputs/K613ArbitrumMarketInput.sol";

/**
 * @title DeployArbitrumSepoliaMarket
 * @notice Deployment script for Aave v3 fork on Arbitrum Sepolia testnet
 * @dev This script deploys the complete Aave v3 protocol with:
 * - L2 Sequencer Uptime Feed integration
 * - Price Oracle Sentinel with grace period
 * - Fallback Oracle (ETH/USD)
 * - All core protocol contracts
 *
 * Usage:
 * forge script scripts/DeployArbitrumSepoliaMarket.sol:DeployArbitrumSepoliaMarket \
 *   --rpc-url $ARBITRUM_SEPOLIA_RPC_URL \
 *   --private-key $PRIVATE_KEY \
 *   --broadcast \
 *   --verify \
 *   --slow
 */
contract DeployArbitrumSepoliaMarket is DeployAaveV3MarketBatchedBase, K613ArbitrumMarketInput {}
