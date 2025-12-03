## 测试思路

## 测试环境信息
l2pool_address:  0xA238Dd80C259a72e81d7e4664a9801593F98d1c5
aave_oracle_address: 0x2Cc0Fc26eD4563A5ce5e8bdcfe1A2878676Ae156
liquidation_provider_address: 0xcf03ed89c297e54935bd0bb6ea33d16abca8ed94
fullnode_endpoint: https://open-platform.nodereal.io/5d9c218e356942a6a9c577e2aadd174c/base

上述合约的abigen生成的golang代码在 src/go/contract下，可以生成对应的合约实例来调用。

## 测试总体思路
1. 首先确定wallet w。
2. 查询当前最新区块高度H
3. 在区块链高度H上直接调用fullnode的 liquidation_provider合约的GetUserPositionFullInfo和GetLiquidationInfo方法，得到结果A。
4. 在区块链高度H上准备golang程序的输入：Global，Reserves 和user数据，然后调用GetUserPositionFullInfo 和 GetLiquidationInfo 方法，得到结果B。
5. 对比结果A和结果B。

## 详细测试路径

获取结果B:
1. 测试地址从test-user.txt文件中随机选择N个，这个N可以指定（N值默认是1）
2. 组装GlobalConfig：
   a. LiquidationAllowed 直接设置为true
   b. EModeCategories 的组装需要先确认有多少给category，调用l2 pool的getEModeCategoryData方法，输入的id从0 迭代到127，直到返回的值是全部为空，则不再继续迭代。然后利用l2 pool的关于EmodeCategory的查询方法，组装好EModeCategories这个mao。
   c. 查询fullnode的当前最新区块头，获取得到timestamp, 赋值给GlobalConfig; 记录下高度H(后续要用到)
3. 组装Reserves：
   a. 通过l2 pool的getReservesList方法得到所有的reserve列表。
   b. 通过l2 pool的其他查询reserve的接口将ReserveData的静态数据补充完整
   c. 通过 aave_oracle 合约的getAssetsPrices 则将PriceInBaseCurrency 一次性补足。（查询时指定高度H）
   d. reserve的decimals等数据，则要生成ERC20实例，再去调用decimals() 方法获取
4. 组装UserData
   a. 通过查询l2 pool的getUserEMode获取UserEModeCategory
   b. 通过查询l2 pool的getUserConfiguration 并解析返回值，可以得到用户在每个reserve上是否有borrow，是否enableAsCollateral，将结果赋值给UserReservePosition。
   c. 查询Asset对应的ATokenAddress 和VariableDebtTokenAddress 的scale balance查询出来（利用scaled_balance_token.go）并赋值。

对于GetLiquidationInfo 需要准备collateralAsset, debtAsset两个参数，填入这两个参数时需要保证collateralAsset的scale balance 不为空， 且enable collateral，并且 debtAsset确实有贷款（如果找不到符合要求的两个asset，则报错即可）


获取结果A:
指定高度H 直接调用liquidation_provider的两个函数，用相同的参数。

## 其他要求
集成测试用例写在 src/go 目录下
