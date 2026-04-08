// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {DeployAaveV3MarketBatchedBase} from "./misc/DeployAaveV3MarketBatchedBase.sol";
import {K613MonadMainnetInput} from "../src/deployments/inputs/K613MonadMainnetInput.sol";

/**
 * @title DeployMonadMainnetMarket
 * @notice Deployment script for Aave v3 fork on Monad mainnet
 *
 * Usage:
 * forge script scripts/DeployMonadMainnetMarket.sol:DeployMonadMainnetMarket \
 *   --rpc-url $RPC_MONAD \
 *   --private-key $PRIVATE_KEY \
 *   --broadcast \
 *   --verify \
 *   --slow
 */
contract DeployMonadMainnetMarket is DeployAaveV3MarketBatchedBase, K613MonadMainnetInput {}
