# include .env file and export its env vars
# (-include to ignore error if it does not exist)
-include .env

# deps
update:; forge update

# Build & test
test   :; forge test -vvv --no-match-contract DeploymentsGasLimits
test-contract :; forge test --match-contract ${filter} -vvv
test-watch   :; forge test --watch -vvv --no-match-contract DeploymentsGasLimits

# Coverage
coverage-base :; forge coverage --fuzz-runs 50 --report lcov --no-match-coverage "(scripts|tests|deployments|mocks)"
coverage-clean :; lcov --rc derive_function_end_line=0 --remove ./lcov.info -o ./lcov.info.p \
	'src/contracts/extensions/v3-config-engine/*' \
	'src/contracts/treasury/*' \
	'src/contracts/dependencies/openzeppelin/ReentrancyGuard.sol' \
	'src/contracts/helpers/UiIncentiveDataProviderV3.sol' \
	'src/contracts/helpers/UiPoolDataProviderV3.sol' \
	'src/contracts/helpers/WalletBalanceProvider.sol' \
	'src/contracts/dependencies/*' \
	'src/contracts/helpers/AaveProtocolDataProvider.sol' \
	'src/contracts/protocol/libraries/configuration/*' \
	'src/contracts/protocol/libraries/logic/GenericLogic.sol' \
	'src/contracts/protocol/libraries/logic/ReserveLogic.sol'
coverage-report :; genhtml ./lcov.info.p -o report --branch-coverage --rc derive_function_end_line=0 --parallel
coverage-badge :; coverage=$$(awk -F '[<>]' '/headerCovTableEntryHi/{print $3}' ./report/index.html | sed 's/[^0-9.]//g' | head -n 1); \
	wget -O ./report/coverage.svg "https://img.shields.io/badge/coverage-$${coverage}%25-brightgreen"
coverage :
	make coverage-base
	make coverage-clean
	make coverage-report
	make coverage-badge


# Utilities
download :; cast etherscan-source --chain ${chain} -d src/etherscan/${chain}_${address} ${address}
git-diff :
	@mkdir -p diffs
	# @npx prettier ${before} ${after} --write
	@printf '%s\n%s\n%s\n' "\`\`\`diff" "$$(git diff --no-index --ignore-space-at-eol ${before} ${after})" "\`\`\`" > diffs/${out}.md

# Deploy
deploy-libs-one	:;
	FOUNDRY_PROFILE=${chain} forge script scripts/misc/LibraryPreCompileOne.sol --rpc-url ${chain} --ledger --mnemonic-indexes ${MNEMONIC_INDEX} --sender ${LEDGER_SENDER} --slow --broadcast --verify
deploy-libs-two	:;
	FOUNDRY_PROFILE=${chain} forge script scripts/misc/LibraryPreCompileTwo.sol --rpc-url ${chain} --ledger --mnemonic-indexes ${MNEMONIC_INDEX} --sender ${LEDGER_SENDER} --slow --broadcast --verify

deploy-libs :
	make deploy-libs-one chain=${chain}
	make deploy-libs-two chain=${chain}

deploy-libs-one-pk :;
	FOUNDRY_PROFILE=${chain} forge script scripts/misc/LibraryPreCompileOne.sol --rpc-url ${chain} --private-key ${PRIVATE_KEY} --slow --broadcast
deploy-libs-two-pk :;
	FOUNDRY_PROFILE=${chain} forge script scripts/misc/LibraryPreCompileTwo.sol --rpc-url ${chain} --private-key ${PRIVATE_KEY} --slow --broadcast

deploy-libs-pk :
	make deploy-libs-one-pk chain=${chain}
	make deploy-libs-two-pk chain=${chain}

deploy-monad :
	FOUNDRY_PROFILE=monad forge script scripts/DeployAaveV3MonadMarket.sol \
	  --rpc-url monad \
	  --private-key ${PRIVATE_KEY} \
	  --slow \
	  --broadcast \
	  --gas-estimate-multiplier 200

list-monad :
	FOUNDRY_PROFILE=monad forge script scripts/ExecuteMonadListing.sol \
	  --rpc-url monad \
	  --private-key ${PRIVATE_KEY} \
	  --slow \
	  --broadcast \
	  --gas-estimate-multiplier 200

deploy-monad-ledger :
	FOUNDRY_PROFILE=monad forge script scripts/DeployAaveV3MonadMarket.sol \
	  --rpc-url monad \
	  --ledger \
	  --mnemonic-indexes ${MNEMONIC_INDEX} \
	  --sender ${LEDGER_SENDER} \
	  --gas-estimate-multiplier 200 \
	  --slow \
	  --broadcast \
	  --verify

# Tenderly virtual testnet workaround: proxies deployed as sub-CREATEs inside a CALL
# transaction do not have their bytecode served by eth_getCode (Monad/Tenderly bug).
# Storage and events are correct; only the code slot is missing. Fix by injecting
# the bytecode from a known-good proxy of the same type via tenderly_setCode.
# Run once after make deploy-monad and make list-monad, before interacting with the market.
fix-monad-proxies :
	@REPORT=$$(ls -t reports/*-market-deployment.json 2>/dev/null | head -1); \
	if [ -z "$$REPORT" ]; then echo "No market deployment report found"; exit 1; fi; \
	echo "Using report: $$REPORT"; \
	POOL_PROXY=$$(python3 -c "import json; d=json.load(open('$$REPORT')); print(d['poolProxy'])"); \
	CONF_PROXY=$$(python3 -c "import json; d=json.load(open('$$REPORT')); print(d['poolConfiguratorProxy'])"); \
	REWARDS_PROXY=$$(python3 -c "import json; d=json.load(open('$$REPORT')); print(d['rewardsControllerProxy'])"); \
	CODE=$$(cast code $$POOL_PROXY --rpc-url ${RPC_MONAD}); \
	for ADDR in $$CONF_PROXY $$REWARDS_PROXY; do \
	  echo "Patching $$ADDR ..."; \
	  curl -sf -X POST "${RPC_MONAD}" -H "Content-Type: application/json" \
	    -d "{\"jsonrpc\":\"2.0\",\"method\":\"tenderly_setCode\",\"params\":[\"$$ADDR\",\"$$CODE\"],\"id\":1}" \
	    | python3 -c "import json,sys; r=json.load(sys.stdin); print('  ok:', r.get('result','ERROR'))"; \
	done; \
	echo "Patching reserve token proxies (aTokens + vDebtTokens) ..."; \
	python3 -c " \
import subprocess, json, sys; \
assets = subprocess.run(['cast','call','${POOL_PROXY}','getReservesList()(address[])','--rpc-url','${RPC_MONAD}'],capture_output=True,text=True).stdout.strip(); \
addrs = [a.strip() for a in assets.strip('[]').split(',')]; \
for asset in addrs: \
  data = subprocess.run(['cast','call','${POOL_PROXY}','getReserveData(address)((uint256,uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128))',asset,'--rpc-url','${RPC_MONAD}'],capture_output=True,text=True).stdout; \
  proxies = [x.strip() for x in data.strip('()').split(',') if x.strip().startswith('0x') and len(x.strip())==42 and x.strip() != '0x0000000000000000000000000000000000000000']; \
  [print(p) for p in proxies] \
" | while read TADDR; do \
	  SZ=$$(cast code $$TADDR --rpc-url ${RPC_MONAD} | wc -c); \
	  if [ "$$SZ" -le 3 ]; then \
	    echo "  Patching $$TADDR (no code) ..."; \
	    curl -sf -X POST "${RPC_MONAD}" -H "Content-Type: application/json" \
	      -d "{\"jsonrpc\":\"2.0\",\"method\":\"tenderly_setCode\",\"params\":[\"$$TADDR\",\"$$CODE\"],\"id\":1}" \
	      | python3 -c "import json,sys; r=json.load(sys.stdin); print('    ok:', r.get('result','ERROR'))"; \
	  else \
	    echo "  $$TADDR ok (has code)"; \
	  fi; \
	done; \
	echo "Done. All proxy bytecodes restored."

# Invariants
echidna:
	echidna tests/invariants/Tester.t.sol --contract Tester --config ./tests/invariants/_config/echidna_config.yaml --corpus-dir ./tests/invariants/_corpus/echidna/default/_data/corpus

echidna-assert:
	echidna tests/invariants/Tester.t.sol --contract Tester --test-mode assertion --config ./tests/invariants/_config/echidna_config.yaml --corpus-dir ./tests/invariants/_corpus/echidna/default/_data/corpus

echidna-explore:
	echidna tests/invariants/Tester.t.sol --contract Tester --test-mode exploration --config ./tests/invariants/_config/echidna_config.yaml --corpus-dir ./tests/invariants/_corpus/echidna/default/_data/corpus

# Medusa
medusa:
	medusa fuzz --config ./medusa.json

# Echidna Runner

HOST = power-runner
LOCAL_FOLDER = ./
REMOTE_FOLDER = ./echidna-runner
REMOTE_COMMAND = cd $(REMOTE_FOLDER)/aave-v3-origin && make echidna > process_output.log 2>&1
REMOTE_COMMAND_ASSERT = cd $(REMOTE_FOLDER)/aave-v3-origin && make echidna-assert > process_output.log 2>&1

echidna-runner:
	tar --exclude='./tests/invariants/_corpus' -czf - $(LOCAL_FOLDER) | ssh $(HOST) "export PATH=$$PATH:/root/.local/bin:/root/.foundry/bin && mkdir -p $(REMOTE_FOLDER)/aave-v3-origin && tar -xzf - -C $(REMOTE_FOLDER)/aave-v3-origin && $(REMOTE_COMMAND)"

echidna-assert-runner:
	tar --exclude='./tests/invariants/_corpus' -czf - $(LOCAL_FOLDER) | ssh $(HOST) "export PATH=$$PATH:/root/.local/bin:/root/.foundry/bin && mkdir -p $(REMOTE_FOLDER)/aave-v3-origin && tar -xzf - -C $(REMOTE_FOLDER)/aave-v3-origin && $(REMOTE_COMMAND_ASSERT)"
