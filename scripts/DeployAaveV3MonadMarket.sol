// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {DeployAaveV3MarketBatchedBase} from './misc/DeployAaveV3MarketBatchedBase.sol';
import {MonadMarketInput} from '../src/deployments/inputs/MonadMarketInput.sol';

/**
 * @title DeployAaveV3MonadMarket
 * @notice Deploys Aave V3 core protocol on Monad mainnet (Chain ID 143).
 *
 * Prerequisites (run once per chain, in order):
 *   1. make deploy-libs-one chain=monad   # deploys BorrowLogic + ConfiguratorLogic via CREATE2
 *   2. make deploy-libs-two chain=monad   # deploys FlashLoanLogic, LiquidationLogic,
 *                                         #   PoolLogic, SupplyLogic via CREATE2
 *   3. make deploy-monad                  # deploys full Aave V3 market
 *
 * Environment variables required in .env:
 *   RPC_MONAD            — Monad RPC endpoint (e.g. https://rpc.monad.xyz or Alchemy URL)
 *   LEDGER_SENDER        — deployer address on the Ledger hardware wallet
 *   MNEMONIC_INDEX       — HD derivation index (default 0)
 *   MONADSCAN_API_KEY    — MonadScan API key for contract verification (optional)
 *
 * Output:
 *   reports/aave-v3-monad-<timestamp>.json  — all deployed contract addresses
 *
 * Post-deployment:
 *   - Transfer marketOwner role to Aave governance executor (EXECUTOR_LVL_1 on Monad)
 *   - Transfer poolAdmin role to governance executor after guardian setup
 *   - Set emergencyAdmin to the Aave Guardian multisig
 */
contract DeployAaveV3MonadMarket is DeployAaveV3MarketBatchedBase, MonadMarketInput {}
