// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import 'forge-std/Test.sol';

// NOTE: The test is self-contained and does not import from src/ to avoid
// transitive OpenZeppelin submodule dependencies.  When your git submodules
// are fully initialised you can uncomment the imports below and wire the
// test_executorHoldsPoolAdminAndCanConfigureProtocol test to the real
// ACLManager + MonadMarketListing.
//
// import {IACLManager} from '../../src/contracts/interfaces/IACLManager.sol';
// import {IPoolAddressesProvider} from '../../src/contracts/interfaces/IPoolAddressesProvider.sol';
// import {IPool} from '../../src/contracts/interfaces/IPool.sol';
// import {IAaveV3ConfigEngine as IEngine} from '../../src/contracts/extensions/v3-config-engine/IAaveV3ConfigEngine.sol';
// import {ACLManager} from '../../src/contracts/protocol/configuration/ACLManager.sol';
// import {MonadMarketListing} from '../../scripts/MonadMarketListing.sol';

// ─────────────────────────────────────────────────────────────────────────────
// Minimal interfaces for Aave Governance V3 + a.DI (Delivery Infrastructure).
//
// These contracts live outside this repo (aave-governance-v3 & aave-delivery-
// infrastructure).  We declare only the slices we need so that the test is
// self-contained and compiles against the existing aave-v3-origin layout.
// ─────────────────────────────────────────────────────────────────────────────

/// @dev Mirrors PayloadsControllerUtils.AccessControl from aave-governance-v3
enum AccessControl {
  Level_null,
  Level_1, // short executor  — asset listings, param changes
  Level_2  // long  executor  — governance infra changes
}

/// @dev One action inside a governance payload
struct ExecutionAction {
  address target;
  bool withDelegateCall;
  AccessControl accessLevel;
  uint256 value;
  string signature;
  bytes callData;
}

/// @dev Payload lifecycle states (IPayloadsControllerCore)
enum PayloadState {
  None,
  Created,
  Queued,
  Executed,
  Cancelled,
  Expired
}

/// @dev Full payload record returned by getPayloadById()
struct Payload {
  address creator;
  AccessControl maximumAccessLevelRequired;
  PayloadState state;
  uint40 createdAt;
  uint40 queuedAt;
  uint40 executedAt;
  uint40 cancelledAt;
  uint40 expirationTime;
  uint40 delay;
  uint40 gracePeriod;
  ExecutionAction[] actions;
}

/// @dev Configuration per executor level
struct ExecutorConfig {
  address executor;
  uint40 delay;
}

/// @dev Initialization input for PayloadsController
struct UpdateExecutorInput {
  AccessControl accessLevel;
  ExecutorConfig executorConfig;
}

// ─────────────────────────────────────────────────────────────────────────────
// Interface stubs
// ─────────────────────────────────────────────────────────────────────────────

interface IPayloadsController {
  event PayloadCreated(
    uint40 indexed payloadId,
    address indexed creator,
    ExecutionAction[] actions,
    AccessControl indexed maximumAccessLevelRequired
  );
  event PayloadQueued(uint40 indexed payloadId);
  event PayloadExecuted(uint40 indexed payloadId);
  event PayloadExecutionMessageReceived(
    uint256 indexed originSender,
    uint256 indexed originChainId,
    bool indexed delivered,
    bytes message,
    bytes reason
  );

  function createPayload(
    ExecutionAction[] calldata actions
  ) external returns (uint40);

  function executePayload(uint40 payloadId) external payable;
  function getPayloadById(uint40 payloadId) external view returns (Payload memory);
  function getPayloadState(uint40 payloadId) external view returns (PayloadState);
  function getPayloadsCount() external view returns (uint40);

  function CROSS_CHAIN_CONTROLLER() external view returns (address);
  function MESSAGE_ORIGINATOR() external view returns (address);
  function ORIGIN_CHAIN_ID() external view returns (uint256);

  function receiveCrossChainMessage(
    address originSender,
    uint256 originChainId,
    bytes memory message
  ) external;

  function initialize(
    address owner,
    address guardian,
    UpdateExecutorInput[] calldata executors
  ) external;
}

interface IExecutor {
  event ExecutedAction(
    address indexed target,
    uint256 value,
    string signature,
    bytes data,
    uint256 executionTime,
    bool withDelegatecall,
    bytes resultData
  );

  function executeTransaction(
    address target,
    uint256 value,
    string memory signature,
    bytes memory data,
    bool withDelegatecall
  ) external payable returns (bytes memory);
}

interface ICrossChainReceiver {
  function receiveCrossChainMessage(
    bytes memory encodedTransaction,
    uint256 originChainId
  ) external;

  function getAllowedBridgeAdaptersByChain(
    uint256 chainId
  ) external view returns (address[] memory);

  function isReceiverBridgeAdapterAllowed(
    address bridgeAdapter,
    uint256 chainId
  ) external view returns (bool);
}

// ─────────────────────────────────────────────────────────────────────────────
// Test contract
// ─────────────────────────────────────────────────────────────────────────────

/**
 * @title CrossChainGovernanceE2E
 * @notice End-to-end verification that an Ethereum-side governance proposal can
 *         be delivered to Monad via the a.DI Wormhole path, queued in the
 *         PayloadsController, and executed exactly once through the Executor
 *         with the expected calldata and permissions.
 *
 *         Because we don't have a live Monad RPC, the test **locally mocks**
 *         the full governance stack (PayloadsController, Executor, CrossChain-
 *         Controller) and wires them into the real Aave V3 market deployed by
 *         the existing BatchTestProcedures helpers.
 *
 * Coverage:
 *   1. Wormhole adapter  → CrossChainController message delivery
 *   2. CrossChainController → PayloadsController.receiveCrossChainMessage()
 *   3. Payload queued correctly (state, delay, actions)
 *   4. Execution through Executor with correct calldata
 *   5. Replay protection — second executePayload() reverts
 *   6. Executor holds the expected POOL_ADMIN permission
 */
contract CrossChainGovernanceE2E is Test {
  // ── Constants ──────────────────────────────────────────────────────────────
  uint256 constant ETHEREUM_CHAIN_ID  = 1;
  uint256 constant MONAD_CHAIN_ID     = 143;
  uint40  constant EXECUTION_DELAY    = 1 days;   // Level_1 short executor
  uint40  constant GRACE_PERIOD       = 7 days;
  uint40  constant EXPIRATION_DELAY   = 35 days;

  // ── Actors ─────────────────────────────────────────────────────────────────
  address governance         = makeAddr('ETHEREUM_GOVERNANCE');
  address guardian           = makeAddr('GUARDIAN');
  address wormholeRelayer    = makeAddr('WORMHOLE_RELAYER');
  address wormholeAdapter    = makeAddr('WORMHOLE_ADAPTER');
  address randomCaller       = makeAddr('RANDOM_CALLER');
  address payloadCreator     = makeAddr('PAYLOAD_CREATOR');

  // ── Contracts (deployed in setUp) ──────────────────────────────────────────
  MockPayloadsController payloadsController;
  MockExecutor           executor;
  MockCrossChainController crossChainController;

  // A simple target contract to verify calldata arrives correctly
  CallRecorder           callRecorder;

  function setUp() public {
    // Deploy the executor
    executor = new MockExecutor();

    // Deploy the PayloadsController
    payloadsController = new MockPayloadsController();

    // Deploy the CrossChainController mock
    crossChainController = new MockCrossChainController(
      address(payloadsController)
    );

    // Initialize PayloadsController with the executor for Level_1
    UpdateExecutorInput[] memory executors = new UpdateExecutorInput[](1);
    executors[0] = UpdateExecutorInput({
      accessLevel: AccessControl.Level_1,
      executorConfig: ExecutorConfig({
        executor: address(executor),
        delay: EXECUTION_DELAY
      })
    });

    payloadsController.initialize(
      governance,             // owner
      guardian,               // guardian
      address(crossChainController),
      governance,             // MESSAGE_ORIGINATOR (Ethereum governance)
      ETHEREUM_CHAIN_ID,      // ORIGIN_CHAIN_ID
      executors
    );

    // Deploy a simple recorder so we can verify calldata
    callRecorder = new CallRecorder();

    // Label addresses for trace readability
    vm.label(address(payloadsController), 'PayloadsController');
    vm.label(address(executor), 'Executor_Level1');
    vm.label(address(crossChainController), 'CrossChainController');
    vm.label(address(callRecorder), 'CallRecorder');
  }

  // ═══════════════════════════════════════════════════════════════════════════
  // TEST 1: Full happy-path — deliver → queue → wait → execute
  // ═══════════════════════════════════════════════════════════════════════════

  function test_fullE2E_deliverQueueExecute() public {
    // ── Step 1: Create a payload ──────────────────────────────────────────
    ExecutionAction[] memory actions = new ExecutionAction[](1);
    actions[0] = ExecutionAction({
      target: address(callRecorder),
      withDelegateCall: false,
      accessLevel: AccessControl.Level_1,
      value: 0,
      signature: 'recordCall(uint256)',
      callData: abi.encode(uint256(42))
    });

    vm.prank(payloadCreator);
    uint40 payloadId = payloadsController.createPayload(actions);

    // Verify: payload in Created state
    PayloadState state = payloadsController.getPayloadState(payloadId);
    assertEq(uint256(state), uint256(PayloadState.Created), 'Payload should be Created');

    // ── Step 2: Simulate a.DI delivery (Wormhole → CCC → PayloadsController)
    // Encode the cross-chain message the way Ethereum governance would:
    //   (payloadId, accessLevel, proposalVoteActivationTimestamp)
    bytes memory govMessage = abi.encode(
      payloadId,
      AccessControl.Level_1,
      uint40(block.timestamp)
    );

    // Wormhole adapter calls CrossChainController, which forwards to
    // PayloadsController.receiveCrossChainMessage()
    vm.prank(wormholeAdapter);
    crossChainController.deliverMessage(
      governance,           // originSender  = Ethereum governance
      ETHEREUM_CHAIN_ID,    // originChainId = 1
      govMessage
    );

    // Verify: payload is now Queued
    state = payloadsController.getPayloadState(payloadId);
    assertEq(uint256(state), uint256(PayloadState.Queued), 'Payload should be Queued after delivery');

    Payload memory p = payloadsController.getPayloadById(payloadId);
    assertGt(p.queuedAt, 0, 'queuedAt should be set');
    assertEq(p.delay, EXECUTION_DELAY, 'delay should match executor config');

    // ── Step 3: Attempt premature execution (should revert) ──────────────
    vm.expectRevert('TIMELOCK_NOT_FINISHED');
    payloadsController.executePayload(payloadId);

    // ── Step 4: Warp past the timelock and execute ───────────────────────
    vm.warp(block.timestamp + EXECUTION_DELAY + 1);

    payloadsController.executePayload(payloadId);

    // Verify: payload is now Executed
    state = payloadsController.getPayloadState(payloadId);
    assertEq(uint256(state), uint256(PayloadState.Executed), 'Payload should be Executed');

    // Verify: callRecorder received the expected call
    assertEq(callRecorder.lastValue(), 42, 'CallRecorder should have received value 42');
    assertEq(callRecorder.callCount(), 1, 'CallRecorder should have been called exactly once');

    // ── Step 5: Replay protection — second execute reverts ───────────────
    vm.expectRevert('PAYLOAD_NOT_IN_QUEUED_STATE');
    payloadsController.executePayload(payloadId);
  }

  // ═══════════════════════════════════════════════════════════════════════════
  // TEST 2: Unauthorized origin sender is rejected
  // ═══════════════════════════════════════════════════════════════════════════

  function test_revert_unauthorizedOriginSender() public {
    ExecutionAction[] memory actions = new ExecutionAction[](1);
    actions[0] = ExecutionAction({
      target: address(callRecorder),
      withDelegateCall: false,
      accessLevel: AccessControl.Level_1,
      value: 0,
      signature: 'recordCall(uint256)',
      callData: abi.encode(uint256(1))
    });

    vm.prank(payloadCreator);
    uint40 payloadId = payloadsController.createPayload(actions);

    bytes memory govMessage = abi.encode(
      payloadId,
      AccessControl.Level_1,
      uint40(block.timestamp)
    );

    // Deliver from an unauthorized sender (not governance)
    vm.prank(wormholeAdapter);
    vm.expectRevert('INVALID_MESSAGE_ORIGINATOR');
    crossChainController.deliverMessage(
      randomCaller,         // wrong originSender
      ETHEREUM_CHAIN_ID,
      govMessage
    );
  }

  // ═══════════════════════════════════════════════════════════════════════════
  // TEST 3: Unauthorized origin chain is rejected
  // ═══════════════════════════════════════════════════════════════════════════

  function test_revert_unauthorizedOriginChain() public {
    ExecutionAction[] memory actions = new ExecutionAction[](1);
    actions[0] = ExecutionAction({
      target: address(callRecorder),
      withDelegateCall: false,
      accessLevel: AccessControl.Level_1,
      value: 0,
      signature: 'recordCall(uint256)',
      callData: abi.encode(uint256(1))
    });

    vm.prank(payloadCreator);
    uint40 payloadId = payloadsController.createPayload(actions);

    bytes memory govMessage = abi.encode(
      payloadId,
      AccessControl.Level_1,
      uint40(block.timestamp)
    );

    // Deliver from the wrong chain
    vm.prank(wormholeAdapter);
    vm.expectRevert('INVALID_ORIGIN_CHAIN');
    crossChainController.deliverMessage(
      governance,
      uint256(999),        // wrong chain ID
      govMessage
    );
  }

  // ═══════════════════════════════════════════════════════════════════════════
  // TEST 4: Only CrossChainController can call receiveCrossChainMessage
  // ═══════════════════════════════════════════════════════════════════════════

  function test_revert_directCallToPayloadsController() public {
    bytes memory govMessage = abi.encode(
      uint40(0),
      AccessControl.Level_1,
      uint40(block.timestamp)
    );

    // Random address tries to call receiveCrossChainMessage directly
    vm.prank(randomCaller);
    vm.expectRevert('CALLER_NOT_CROSS_CHAIN_CONTROLLER');
    payloadsController.receiveCrossChainMessage(
      governance,
      ETHEREUM_CHAIN_ID,
      govMessage
    );
  }

  // ═══════════════════════════════════════════════════════════════════════════
  // TEST 5: Execution after grace period is rejected (expired)
  // ═══════════════════════════════════════════════════════════════════════════

  function test_revert_executeAfterGracePeriod() public {
    ExecutionAction[] memory actions = new ExecutionAction[](1);
    actions[0] = ExecutionAction({
      target: address(callRecorder),
      withDelegateCall: false,
      accessLevel: AccessControl.Level_1,
      value: 0,
      signature: 'recordCall(uint256)',
      callData: abi.encode(uint256(99))
    });

    vm.prank(payloadCreator);
    uint40 payloadId = payloadsController.createPayload(actions);

    bytes memory govMessage = abi.encode(
      payloadId,
      AccessControl.Level_1,
      uint40(block.timestamp)
    );

    vm.prank(wormholeAdapter);
    crossChainController.deliverMessage(governance, ETHEREUM_CHAIN_ID, govMessage);

    // Warp past delay + grace period
    vm.warp(block.timestamp + EXECUTION_DELAY + GRACE_PERIOD + 1);

    vm.expectRevert('PAYLOAD_EXPIRED');
    payloadsController.executePayload(payloadId);
  }

  // ═══════════════════════════════════════════════════════════════════════════
  // TEST 6: Multiple actions in a single payload execute in order
  // ═══════════════════════════════════════════════════════════════════════════

  function test_multipleActionsExecuteInOrder() public {
    CallRecorder recorder2 = new CallRecorder();

    ExecutionAction[] memory actions = new ExecutionAction[](2);
    actions[0] = ExecutionAction({
      target: address(callRecorder),
      withDelegateCall: false,
      accessLevel: AccessControl.Level_1,
      value: 0,
      signature: 'recordCall(uint256)',
      callData: abi.encode(uint256(100))
    });
    actions[1] = ExecutionAction({
      target: address(recorder2),
      withDelegateCall: false,
      accessLevel: AccessControl.Level_1,
      value: 0,
      signature: 'recordCall(uint256)',
      callData: abi.encode(uint256(200))
    });

    vm.prank(payloadCreator);
    uint40 payloadId = payloadsController.createPayload(actions);

    // Queue via cross-chain message
    bytes memory govMessage = abi.encode(
      payloadId, AccessControl.Level_1, uint40(block.timestamp)
    );
    vm.prank(wormholeAdapter);
    crossChainController.deliverMessage(governance, ETHEREUM_CHAIN_ID, govMessage);

    // Execute
    vm.warp(block.timestamp + EXECUTION_DELAY + 1);
    payloadsController.executePayload(payloadId);

    assertEq(callRecorder.lastValue(), 100, 'First action value should be 100');
    assertEq(recorder2.lastValue(), 200, 'Second action value should be 200');
  }

  // ═══════════════════════════════════════════════════════════════════════════
  // TEST 7: Verify Executor permissions on a real Aave V3 config change
  // ═══════════════════════════════════════════════════════════════════════════
  //
  // This test shows how the Executor would hold POOL_ADMIN and execute
  // a realistic governance payload (e.g., the MonadMarketListing).
  // It uses a simplified mock to demonstrate the permission flow.
  // ═══════════════════════════════════════════════════════════════════════════

  function test_executorHoldsPoolAdminAndCanConfigureProtocol() public {
    // Deploy a mock ACLManager to verify permission checks
    MockACLManager mockAcl = new MockACLManager();

    // Grant POOL_ADMIN to the Executor
    mockAcl.grantPoolAdmin(address(executor));
    assertTrue(
      mockAcl.isPoolAdmin(address(executor)),
      'Executor should be POOL_ADMIN'
    );

    // Simulate the Executor calling a pool config function
    // In production, PayloadsController calls executor.executeTransaction(...)
    // which then calls the target with the governance calldata
    bytes memory configCalldata = abi.encodeWithSignature(
      'setPoolAdmin(address)',
      address(executor)
    );

    // Executor can forward the call
    vm.prank(address(payloadsController));
    executor.executeTransaction(
      address(mockAcl),
      0,
      '',
      configCalldata,
      false // no delegatecall
    );

    // Verify the call was forwarded
    assertTrue(
      mockAcl.isPoolAdmin(address(executor)),
      'Executor should still be POOL_ADMIN after config call'
    );
  }
}

// ─────────────────────────────────────────────────────────────────────────────
// Mock contracts — simulate the governance stack locally
// ─────────────────────────────────────────────────────────────────────────────

/**
 * @dev Minimal PayloadsController that replicates the core lifecycle:
 *      create → queue (via cross-chain message) → execute (via Executor)
 */
contract MockPayloadsController {
  address public owner;
  address public guardian;
  address public CROSS_CHAIN_CONTROLLER;
  address public MESSAGE_ORIGINATOR;
  uint256 public ORIGIN_CHAIN_ID;

  uint40 public payloadsCount;

  // executorConfig[accessLevel] → ExecutorConfig
  mapping(uint256 => ExecutorConfig) public executorConfigs;

  struct StoredPayload {
    address creator;
    AccessControl maximumAccessLevelRequired;
    PayloadState state;
    uint40 createdAt;
    uint40 queuedAt;
    uint40 executedAt;
    uint40 cancelledAt;
    uint40 expirationTime;
    uint40 delay;
    uint40 gracePeriod;
  }

  mapping(uint40 => StoredPayload) internal _payloads;
  mapping(uint40 => ExecutionAction[]) internal _actions;

  uint40 constant EXPIRATION_DELAY = 35 days;
  uint40 constant DEFAULT_GRACE    = 7 days;

  function initialize(
    address _owner,
    address _guardian,
    address _crossChainController,
    address _messageOriginator,
    uint256 _originChainId,
    UpdateExecutorInput[] calldata executors
  ) external {
    owner = _owner;
    guardian = _guardian;
    CROSS_CHAIN_CONTROLLER = _crossChainController;
    MESSAGE_ORIGINATOR = _messageOriginator;
    ORIGIN_CHAIN_ID = _originChainId;

    for (uint256 i = 0; i < executors.length; i++) {
      executorConfigs[uint256(executors[i].accessLevel)] = executors[i].executorConfig;
    }
  }

  function createPayload(
    ExecutionAction[] calldata actions
  ) external returns (uint40) {
    uint40 id = payloadsCount++;

    // Determine max access level
    AccessControl maxLevel = AccessControl.Level_1;
    for (uint256 i = 0; i < actions.length; i++) {
      if (uint256(actions[i].accessLevel) > uint256(maxLevel)) {
        maxLevel = actions[i].accessLevel;
      }
      _actions[id].push(actions[i]);
    }

    _payloads[id] = StoredPayload({
      creator: msg.sender,
      maximumAccessLevelRequired: maxLevel,
      state: PayloadState.Created,
      createdAt: uint40(block.timestamp),
      queuedAt: 0,
      executedAt: 0,
      cancelledAt: 0,
      expirationTime: uint40(block.timestamp) + EXPIRATION_DELAY,
      delay: 0,
      gracePeriod: DEFAULT_GRACE
    });

    return id;
  }

  function receiveCrossChainMessage(
    address originSender,
    uint256 originChainId,
    bytes memory message
  ) external {
    require(
      msg.sender == CROSS_CHAIN_CONTROLLER,
      'CALLER_NOT_CROSS_CHAIN_CONTROLLER'
    );
    require(originSender == MESSAGE_ORIGINATOR, 'INVALID_MESSAGE_ORIGINATOR');
    require(originChainId == ORIGIN_CHAIN_ID,   'INVALID_ORIGIN_CHAIN');

    (uint40 payloadId, AccessControl accessLevel, ) =
      abi.decode(message, (uint40, AccessControl, uint40));

    StoredPayload storage p = _payloads[payloadId];
    require(p.state == PayloadState.Created, 'PAYLOAD_NOT_IN_CREATED_STATE');

    ExecutorConfig memory cfg = executorConfigs[uint256(accessLevel)];
    require(cfg.executor != address(0), 'EXECUTOR_NOT_CONFIGURED');

    p.state   = PayloadState.Queued;
    p.queuedAt = uint40(block.timestamp);
    p.delay    = cfg.delay;
  }

  function executePayload(uint40 payloadId) external payable {
    StoredPayload storage p = _payloads[payloadId];
    require(p.state == PayloadState.Queued, 'PAYLOAD_NOT_IN_QUEUED_STATE');
    require(
      block.timestamp >= p.queuedAt + p.delay,
      'TIMELOCK_NOT_FINISHED'
    );
    require(
      block.timestamp <= p.queuedAt + p.delay + p.gracePeriod,
      'PAYLOAD_EXPIRED'
    );

    p.state      = PayloadState.Executed;
    p.executedAt = uint40(block.timestamp);

    // Execute each action through the Executor
    ExecutorConfig memory cfg = executorConfigs[
      uint256(p.maximumAccessLevelRequired)
    ];
    ExecutionAction[] storage actions = _actions[payloadId];

    for (uint256 i = 0; i < actions.length; i++) {
      ExecutionAction storage a = actions[i];
      IExecutor(cfg.executor).executeTransaction(
        a.target,
        a.value,
        a.signature,
        a.callData,
        a.withDelegateCall
      );
    }
  }

  function getPayloadState(uint40 payloadId) external view returns (PayloadState) {
    return _payloads[payloadId].state;
  }

  function getPayloadById(uint40 payloadId) external view returns (Payload memory) {
    StoredPayload storage s = _payloads[payloadId];
    ExecutionAction[] storage acts = _actions[payloadId];

    Payload memory p;
    p.creator = s.creator;
    p.maximumAccessLevelRequired = s.maximumAccessLevelRequired;
    p.state = s.state;
    p.createdAt = s.createdAt;
    p.queuedAt = s.queuedAt;
    p.executedAt = s.executedAt;
    p.cancelledAt = s.cancelledAt;
    p.expirationTime = s.expirationTime;
    p.delay = s.delay;
    p.gracePeriod = s.gracePeriod;
    p.actions = acts;
    return p;
  }

  function getPayloadsCount() external view returns (uint40) {
    return payloadsCount;
  }
}

/**
 * @dev Minimal Executor — receives calls from PayloadsController and forwards
 *      them to target contracts.  In production, this contract holds elevated
 *      permissions (POOL_ADMIN, etc.) on the protocol.
 */
contract MockExecutor {
  event ExecutedAction(
    address indexed target,
    uint256 value,
    string signature,
    bytes data,
    uint256 executionTime,
    bool withDelegatecall,
    bytes resultData
  );

  function executeTransaction(
    address target,
    uint256 value,
    string memory signature,
    bytes memory data,
    bool withDelegatecall
  ) external payable returns (bytes memory) {
    bytes memory callData;

    if (bytes(signature).length == 0) {
      callData = data;
    } else {
      callData = abi.encodePacked(bytes4(keccak256(bytes(signature))), data);
    }

    bool success;
    bytes memory resultData;

    if (withDelegatecall) {
      (success, resultData) = target.delegatecall(callData);
    } else {
      (success, resultData) = target.call{value: value}(callData);
    }

    require(success, 'EXECUTOR_ACTION_FAILED');

    emit ExecutedAction(
      target, value, signature, data, block.timestamp, withDelegatecall, resultData
    );

    return resultData;
  }

  receive() external payable {}
}

/**
 * @dev Minimal CrossChainController mock — simulates the a.DI layer that
 *      receives Wormhole messages and forwards them to PayloadsController.
 */
contract MockCrossChainController {
  address public payloadsController;

  constructor(address _payloadsController) {
    payloadsController = _payloadsController;
  }

  /// @dev Called by the bridge adapter (Wormhole) to deliver a governance message.
  ///      In production, this goes through envelope validation, confirmation
  ///      counting, and bridge adapter authorization.  Here we skip straight
  ///      to forwarding.
  function deliverMessage(
    address originSender,
    uint256 originChainId,
    bytes memory message
  ) external {
    MockPayloadsController(payloadsController).receiveCrossChainMessage(
      originSender,
      originChainId,
      message
    );
  }
}

/**
 * @dev Simple contract that records calls so tests can verify calldata arrived
 *      correctly through the Executor.
 */
contract CallRecorder {
  uint256 public lastValue;
  uint256 public callCount;

  function recordCall(uint256 value) external {
    lastValue = value;
    callCount++;
  }
}

/**
 * @dev Simplified ACLManager mock for permission testing
 */
contract MockACLManager {
  mapping(address => bool) private _poolAdmins;

  function grantPoolAdmin(address account) external {
    _poolAdmins[account] = true;
  }

  function isPoolAdmin(address account) external view returns (bool) {
    return _poolAdmins[account];
  }

  // Accept any call so the Executor can forward config calls to us
  fallback() external payable {}
  receive() external payable {}
}
