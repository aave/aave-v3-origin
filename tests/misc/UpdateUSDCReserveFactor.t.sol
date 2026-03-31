// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import 'forge-std/Test.sol';
import {IACLManager} from '../../src/contracts/interfaces/IACLManager.sol';
import {IPool} from '../../src/contracts/interfaces/IPool.sol';
import {DataTypes} from '../../src/contracts/protocol/libraries/types/DataTypes.sol';
import {ReserveConfiguration} from '../../src/contracts/protocol/libraries/configuration/ReserveConfiguration.sol';
import {UpdateUSDCReserveFactor} from '../../scripts/UpdateUSDCReserveFactor.sol';

/**
 * @title UpdateUSDCReserveFactorTest
 * @notice Plan.md Section 10 — Tenderly fork shortcut.
 *
 *   Validates that the USDC reserve factor payload works correctly on the
 *   Monad Tenderly fork WITHOUT deploying any governance infrastructure
 *   (no CrossChainController, PayloadsController, or Executor).
 *
 *   The test pranks the deployer EOA (who holds DEFAULT_ADMIN in ACLManager)
 *   to grant POOL_ADMIN to the payload contract, then calls execute() directly.
 *
 * Run:
 *   forge test --match-contract UpdateUSDCReserveFactorTest --fork-url $RPC_MONAD -vvv
 */
contract UpdateUSDCReserveFactorTest is Test {
  using ReserveConfiguration for DataTypes.ReserveConfigurationMap;

  // ── Deployed addresses (reports/1774822427-market-deployment.json) ────────
  address internal constant ACL_MANAGER = 0x56Ae0156ac4A2d4a38B98201183BC917c6851832;
  address internal constant POOL        = 0x021b65d3F66b5e63e7Ac1Fe94bB56124F381cD30;

  // Deployer EOA — holds DEFAULT_ADMIN + POOL_ADMIN on the Monad Tenderly fork.
  // Confirmed by: cast call $ACL "hasRole(bytes32,address)(bool)" $DEFAULT_ADMIN_ROLE $DEPLOYER
  address internal constant DEPLOYER    = 0x82205F524ABb0468bA5f1Bbc1439Aa7F6669fb96;

  // ── Asset ─────────────────────────────────────────────────────────────────
  address internal constant USDC        = 0x754704Bc059F8C67012fEd69BC8A327a5aafb603;

  UpdateUSDCReserveFactor internal payload;

  function setUp() public {
    vm.createSelectFork(vm.rpcUrl('monad'));
    payload = new UpdateUSDCReserveFactor();
  }

  /// @dev Sanity check: USDC reserve factor starts at 10% before any change.
  function test_usdcReserveFactor_startsWith10Pct() public view {
    DataTypes.ReserveConfigurationMap memory cfg = IPool(POOL).getConfiguration(USDC);
    assertEq(cfg.getReserveFactor(), 1000, 'initial reserveFactor must be 1000 (10%)');
  }

  /// @dev Core test: execute the payload and verify the reserve factor changes to 15%.
  function test_executePayload_changesReserveFactorTo15Pct() public {
    // Grant POOL_ADMIN to the payload contract so it can call PoolConfigurator.
    // (In production the Executor Level 1 holds POOL_ADMIN and delegatecalls the payload.)
    vm.prank(DEPLOYER);
    IACLManager(ACL_MANAGER).addPoolAdmin(address(payload));

    payload.execute();

    DataTypes.ReserveConfigurationMap memory cfg = IPool(POOL).getConfiguration(USDC);
    assertEq(cfg.getReserveFactor(), 1500, 'reserveFactor must be 1500 (15%) after execution');
  }

  /// @dev Guard test: no other bits in the reserve configuration bitmask change.
  function test_noOtherConfigBitsChanged() public {
    DataTypes.ReserveConfigurationMap memory before = IPool(POOL).getConfiguration(USDC);

    vm.prank(DEPLOYER);
    IACLManager(ACL_MANAGER).addPoolAdmin(address(payload));
    payload.execute();

    DataTypes.ReserveConfigurationMap memory after_ = IPool(POOL).getConfiguration(USDC);

    // Reserve factor occupies bits 64-79 of the configuration word.
    // Mask those bits out and assert everything else is identical.
    uint256 rfMask = ~(uint256(0xFFFF) << 64);
    assertEq(
      before.data & rfMask,
      after_.data & rfMask,
      'only the reserveFactor bits (64-79) should change'
    );
  }

  /// @dev Confirm that calling execute() a second time is idempotent on the reserve factor.
  function test_idempotent_doubleExecute() public {
    vm.prank(DEPLOYER);
    IACLManager(ACL_MANAGER).addPoolAdmin(address(payload));

    payload.execute();
    payload.execute();

    DataTypes.ReserveConfigurationMap memory cfg = IPool(POOL).getConfiguration(USDC);
    assertEq(cfg.getReserveFactor(), 1500, 'reserve factor should still be 1500 after second execute');
  }
}
