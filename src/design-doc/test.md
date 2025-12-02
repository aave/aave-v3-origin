## 测试思路

- **基础场景**：构造一个包含抵押资产与债务资产的 `SimulationInput`（价格、精度、LTV、LT、清算奖励、协议费、位图、宽限期等完整配置），确保 `GetUserPositionFullInfo` 返回的总抵押、总债务、平均 LTV/LT、可借额度与健康度符合手工计算。
- **健康仓位不触发清算**：设置 `healthFactor >= 1e18`，调用 `GetLiquidationInfo` 应返回 0 的可清算额度及 `amountToPassToLiquidationCall` 为 0。
- **清算基础路径**：设置 `healthFactor < 1e18`，债务/抵押均大于 `MIN_BASE_MAX_CLOSE_FACTOR_THRESHOLD`，验证 `maxDebtToLiquidate` 与 `_getAvailableCollateralAndDebtToLiquidate` 计算出的债务/抵押和协议费与 Solidity 逻辑一致。
- **Close Factor 分支**：构造 `healthFactor > 0.95e18` 且抵押、债务基准价值均超阈值，确认 `maxDebtToLiquidate` 被限制为总债务的 50%。
- **MIN_LEFTOVER 调整**：人为设置抵押或债务清算后剩余低于 `MIN_LEFTOVER_BASE` 的场景，验证 `_adjustAmountsForGoodLeftovers` 会调整抵押/债务/协议费，结果剩余都不低于阈值。
- **eMode 奖励覆写**：设置用户 eMode 且资产位图包含抵押资产，校验清算奖励使用 eMode 配置值。
- **返回值边界**：当本次清算覆盖全部债务或全部抵押时，`amountToPassToLiquidationCall` 应为 `MaxUint256`，否则等于 `maxDebtToLiquidate`。

> 可用 Go 测试对以上场景进行断言：在 `src/go` 下创建多个测试用例，重建输入数据后分别调用 `GetUserPositionFullInfo` 与 `GetLiquidationInfo`，对比期望数值（含加总、bps 计算与 WAD 健康度）。
