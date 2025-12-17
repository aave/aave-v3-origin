## golang 程序实现 aave v3 协议部分重要功能的模拟

### 背景

我们的目的是做 aave的清算者，我们需要快速计算出用户的仓位的Ltv，healthy factor等参数以及getLiquidationInfo等信息。
AAVE提供了以下两个我们主要需要的合约函数：
1. LiquidationDataProvider.sol: getUserPositionFullInfo
2. LiquidationDataProvider.sol: getLiquidationInfo

我们的目的是需要用golang来实现上面的两个函数，从而在本地可以进行快速计算而不需要再调用区块链节点来查询。

### 实现路径

1. 请首先仔细阅读本项目（即aave v3）的代码（src/contracts目录下），理解其逻辑和实现方式。重点关注和上面两个函数相关的逻辑和数据。
2. 对两个函数中代码执行过程中需要用到的数据进行梳理，数据分为3类，一类全局配置数据，一类是和资产相关的数据，一类是和用户相关的数据，将数据定义写在design-doc/data.md中，如果还需要定义其他类别的数据，请在design-doc/data.md中一并添加。
3. 根据数据定义，编写golang代码实现两个函数，函数的输入包好步骤2中定义的数据结构。
4. 为了确保代码的正确性，请设计测试文档到design-doc/test.md中。
5. golang代码写在 src/go 文件夹下，适当拆分为2-3个文件。