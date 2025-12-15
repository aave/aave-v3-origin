## Integration Test

### 目标
在base Sepolia Testnet 部署一套完整的aave v3 marketplace合约，用于简单的测试验证。

### 环境信息
部署者地址： 0xd75Fea05fdafD031e2a584Db1d4E6155c862B875
RPC endpoint: https://base-sepolia-public.nodies.app
chain-id: 84532
### 细节要求
1. 部署测试测试token(TestnetERC20.sol)，比如USDC（decimal为6），USDT（decimal为6），WBTC（decimal为8），WETH（decimal为18），每个token的初始supply是1e9个(不考虑decimal)分配给合约部署者。 注意把WETH当成普通的ERC20 合约即可，不需要和ETH之间互换。
2. 分别部署USDC，USDT，WBTC， WETH 四种token的oracle price mock合约（复用MockAggregatorSetPrice.sol）, 最初的价格是：1，1，88700，3078。 这些价格可以被随时调整。
3. 由于base testnet是L2， 所以我们部署的核心协议是L2Pool.sol 而不是Pool.sol。现在需要将整套 aave v3部署到 base testnt上：
    a. 部署必要的核心合约，主要是以下几个：'POOL'，'POOL_CONFIGURATOR'，'PRICE_ORACLE'，'ACL_MANAGER'，'ACL_ADMIN'， 'PRICE_ORACLE_SENTINEL'， 'DATA_PROVIDER'
    b. 部署一个周边合约 LiquidationDataProvider.sol
4. 务必配置和初始化好各个合约使其能够工作，所有需要设置权限的地方都配置为部署者地址。


以上步骤请先充分了解本项目的代码结构，尤其是 tests和scripts文件夹下已经有的部署代码和脚本。

### 输出要求
分为两个阶段：

说明阶段：
1. 在deployment-explain.md中说明是否需要在scripts下构建新的脚本，你准备添加的额外文件和对应的功能。

在我的明确指令下进行执行阶段：
1. 根据deployment-explain.md的内容，完成新的代码/文档的编写和测试
2. 在部署脚本中，增加一个简单的测试用例：
   a. 部署者质押1个WETH并允许作为一个质押品。
   b. 部署者质押1000个USDC但不允许作为一个质押品。
   c. 部署者借出900个USDC。
   d. 部署者归还所需的所有USDC。
   e. 部署者设置 WETH的价格为3100.

