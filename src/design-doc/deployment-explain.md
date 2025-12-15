## 是否需要新增脚本
- 需要。现有 `scripts/DeployAaveV3MarketBatched*.sol` 偏向通用批量部署/预编译库，不直接覆盖“本地测试 + base sepolia L2Pool + 自定义 mock token/oracle”的需求。

## 计划新增/修改的文件与功能
- `scripts/DeployAaveV3BaseSepolia.sol`（新）：一站式 forge script，完成：
  - 部署四个测试代币（USDC 6d、USDT 6d、WBTC 8d、WETH 18d，初始供应分配给部署者）。
  - 部署四个 `MockAggregatorSetPrice`，初始价格 1/1/88700/3078，可后续手动调价。
  - 部署并初始化 L2 核心组件：`POOL`(L2Pool)、`POOL_CONFIGURATOR`、`PRICE_ORACLE`（注入上述 sources）、`ACL_MANAGER`、`ACL_ADMIN`、`PRICE_ORACLE_SENTINEL`、`DATA_PROVIDER`，以及外围 `LiquidationDataProvider`，把需要的角色都授予部署者。
  - 将代币上市到池子（配置利率策略/ltv 等可用固定参数），确保可正常存借清算。
  - 输出部署产物（JSON/console）便于后续测试复用。
- `scripts/misc/DeployAaveV3MarketBatchedBase.sol`（可选小改）：如果复用其部分初始化逻辑，可能添加帮助函数暴露更细粒度的部署步骤，否则不改。
- `src/design-doc/deployment.md`（仅参考，不改代码）：作为需求来源。

## 复用思路
- 复用 `TestnetERC20` 与 `MockAggregatorSetPrice` 作为资产与喂价。
- 参考 `DeployAaveV3MarketBatched.sol` 的结构组织脚本，但针对 base-sepolia chainId、L2Pool 构造和自定义 token/oracle。
- 使用 forge standard-json 或 broadcast 方式，脚本中支持从环境变量读取私钥/RPC，避免硬编码。

## 其他说明
- 说明阶段不改代码，执行阶段再按上面文件清单落地。

## 运行指引（执行阶段给你的操作步骤）
- 依赖：`forge`，RPC 指向 `https://base-sepolia-public.nodies.app`（或自选 base sepolia 节点）。
- 环境变量：`PRIVATE_KEY`（部署者私钥，对应文档中的 0xd75Fe... 地址）。
- 执行命令示例：
  - `forge script scripts/DeployAaveV3BaseSepolia.sol:DeployAaveV3BaseSepolia --rpc-url https://base-sepolia-public.nodies.app --broadcast --verify=false`
- 结果：控制台打印四个 token 地址、对应 oracle、SequencerOracle、Pool、ConfigEngine、AaveOracle。价格可通过 `MockAggregatorSetPrice.setLatestAnswer(int256)` 调整，Sequencer 状态可通过 `SequencerOracle.setAnswer(bool,uint256)` 手动改。
