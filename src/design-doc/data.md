## 计算依赖数据定义

### 全局配置数据
- `currentTimestamp` (`uint64`): 当前区块时间戳，用于判断 `liquidationGracePeriodUntil` 是否已过。
- `liquidationAllowed` (`bool`): 价格哨兵返回值，决定在 `healthFactor` 低于 0.95 且高于 1.0 的区间内是否允许清算。
- `eModeCategories` (`map[uint8]EModeCategory`): eMode 配置集合，用于确定清算加成、LTV、LT 时的替代参数。

### EModeCategory
- `ltv` (`uint64`): 该类别下的 LTV（bps，1e4 精度）。
- `liquidationThreshold` (`uint64`): 该类别下的清算阈值（bps）。
- `liquidationBonus` (`uint64`): 该类别下的清算奖励（bps）。
- `collateralBitmap` (`uint128`): 允许作为抵押的资产位图。

### 资产级数据（每个 Reserve）
- `asset` (`Address`): 资产地址标识（`common.Address`）。
- `id` (`uint16`): 资产在位图中的索引。
- `decimals` (`uint8`): 资产精度，用于计算 `assetUnit = 10^decimals`。
- `priceInBaseCurrency` (`*big.Int`): 资产价格（基准币计价）。
- `ltv` (`uint64`): 资产的 LTV（bps）。
- `liquidationThreshold` (`uint64`): 资产的清算阈值（bps）。
- `liquidationBonus` (`uint64`): 资产的清算奖励（bps）。
- `liquidationProtocolFee` (`uint64`): 清算协议费（bps）。
- `normalizedIncome` (`*big.Int`): 归一化收益指数（ray，1e27），用于将 aToken scaled balance 转为实际余额。
- `normalizedDebt` (`*big.Int`): 归一化借款指数（ray，1e27），用于将 variable debt scaled balance 转为实际债务。
- `isActive` (`bool`): 资产是否可用。
- `isPaused` (`bool`): 是否暂停。
- `liquidationGracePeriodUntil` (`uint64`): 清算宽限期截止时间戳。
- `aTokenAddress` (`string`): 对应的 aToken。
- `variableDebtTokenAddress` (`string`): 对应的可变债务 Token。

### 用户级数据
- `userAddress` (`Address`): 用户标识（`common.Address`）。
- `userEModeCategory` (`uint8`): 用户当前选择的 eMode 类别。
- `positions` (`[]UserReservePosition`): 用户在各资产上的仓位明细。

### UserReservePosition
- `asset` (`Address`): 资产地址标识（`common.Address`）。
- `useAsCollateral` (`bool`): 是否将该资产用作抵押。
- `isBorrowing` (`bool`): 用户是否在该资产上有借款开关。
- `collateralScaledBalance` (`*big.Int`): aToken 的 scaled balance（保持不变，需乘以 `normalizedIncome` 得到实际余额）。
- `debtScaledBalance` (`*big.Int`): variable debt 的 scaled balance（保持不变，需乘以 `normalizedDebt` 得到实际债务）。
