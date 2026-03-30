// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {Script} from 'forge-std/Script.sol';
import {IACLManager} from '../src/contracts/interfaces/IACLManager.sol';
import {IAaveV3ConfigEngine as IEngine} from '../src/contracts/extensions/v3-config-engine/IAaveV3ConfigEngine.sol';
import {MonadMarketListing} from './MonadMarketListing.sol';

/**
 * @title ExecuteMonadListing
 * @notice Forge script that lists WMON and USDC on the deployed Aave V3 Monad market.
 *
 * All addresses match reports/1774818808-market-deployment.json.
 *
 * Usage:
 *   make list-monad
 */
contract ExecuteMonadListing is Script {
  // ── Deployed market addresses (from reports/1774822427-market-deployment.json) ──
  address internal constant CONFIG_ENGINE = 0x521F472C6b1EDBb3B6b7790e83841Ac29223F4fA;
  address internal constant ACL_MANAGER   = 0x56Ae0156ac4A2d4a38B98201183BC917c6851832;
  address internal constant ATOKEN_IMPL   = 0x47fecE01E54Fd5851E65322bB4fc468cD098766C;
  address internal constant VDEBT_IMPL    = 0xCE07bC667B8038333A8B77431cFeAF866cee2129;

  // ── Asset addresses ──────────────────────────────────────────────────────

  // WMON — Wrapped native MON token on Monad mainnet
  address internal constant WMON = 0x3bd359C1119dA7Da1D913D1C4D2B7c461115433A;

  // Circle CCTP v2 USDC on Monad mainnet (proxy)
  address internal constant USDC = 0x754704Bc059F8C67012fEd69BC8A327a5aafb603;

  // MON/USD — confirmed at data.chain.link/feeds/monad
  address internal constant MON_USD_FEED = 0xBcD78f76005B7515837af6b50c7C52BCf73822fb;

  // USDC/USD — confirmed at data.chain.link/feeds/monad
  address internal constant USDC_USD_FEED = 0xf5F15f188AbCB0d165D1Edb7f37F7d6fA2fCebec;

  function run() external {
    vm.startBroadcast();

    MonadMarketListing listing = new MonadMarketListing(
      IEngine(CONFIG_ENGINE),
      WMON,
      MON_USD_FEED,
      USDC,
      USDC_USD_FEED,
      ATOKEN_IMPL,
      VDEBT_IMPL,
      ACL_MANAGER
    );

    IACLManager(ACL_MANAGER).addPoolAdmin(address(listing));
    listing.execute();

    vm.stopBroadcast();
  }
}
