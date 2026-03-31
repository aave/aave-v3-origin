// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {AaveV3Payload} from '../src/contracts/extensions/v3-config-engine/AaveV3Payload.sol';
import {IAaveV3ConfigEngine as IEngine} from '../src/contracts/extensions/v3-config-engine/IAaveV3ConfigEngine.sol';
import {EngineFlags} from '../src/contracts/extensions/v3-config-engine/EngineFlags.sol';

/**
 * @title UpdateUSDCReserveFactor
 * @notice Governance payload that increases the USDC reserve factor on Aave V3 Monad
 *         from 10% (1000 bps) to 15% (1500 bps).
 *
 * Tenderly fork shortcut:
 *   Deploy this contract, grant it POOL_ADMIN in ACLManager, then call execute() directly.
 *   This bypasses the full a.DI governance bridge for rapid local testing.
 *
 * Production flow:
 *   The Executor Level 1 would delegatecall this via PayloadsController after a
 *   governance vote on Ethereum. The Executor holds POOL_ADMIN, so no manual grant needed.
 *
 * Execute() call chain (all delegatecalls until final PoolConfigurator call):
 *   payload.execute()
 *     → address(CONFIG_ENGINE).functionDelegateCall(updateBorrowSide)
 *     → BORROW_ENGINE.functionDelegateCall(executeBorrowSide)
 *     → poolConfigurator.setReserveFactor(USDC, 1500)   [msg.sender = this contract]
 *
 * Therefore this contract must have POOL_ADMIN or RISK_ADMIN before execute() is called.
 */
contract UpdateUSDCReserveFactor is AaveV3Payload {
  // USDC — Circle CCTP v2 on Monad (proxy)
  address public constant USDC = 0x754704Bc059F8C67012fEd69BC8A327a5aafb603;

  // New reserve factor: 15.00% expressed as basis points (1 bps = 0.01%)
  uint256 public constant NEW_RESERVE_FACTOR = 15_00;

  // ConfigEngine deployed by aave-v3-origin on the Monad Tenderly fork
  // Source: reports/1774822427-market-deployment.json
  address internal constant MONAD_CONFIG_ENGINE = 0x521F472C6b1EDBb3B6b7790e83841Ac29223F4fA;

  constructor() AaveV3Payload(IEngine(MONAD_CONFIG_ENGINE)) {}

  /// @dev Change USDC reserve factor from 10% to 15%. All other borrow params unchanged.
  function borrowsUpdates()
    public
    pure
    override
    returns (IEngine.BorrowUpdate[] memory)
  {
    IEngine.BorrowUpdate[] memory updates = new IEngine.BorrowUpdate[](1);
    updates[0] = IEngine.BorrowUpdate({
      asset: USDC,
      enabledToBorrow: EngineFlags.KEEP_CURRENT,
      flashloanable: EngineFlags.KEEP_CURRENT,
      borrowableInIsolation: EngineFlags.KEEP_CURRENT,
      withSiloedBorrowing: EngineFlags.KEEP_CURRENT,
      reserveFactor: NEW_RESERVE_FACTOR
    });
    return updates;
  }

  function getPoolContext() public pure override returns (IEngine.PoolContext memory) {
    return IEngine.PoolContext({networkName: 'Monad', networkAbbreviation: 'Mon'});
  }
}
