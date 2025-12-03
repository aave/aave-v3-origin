// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ILiquidationDataProviderCollateralFullInfo is an auto generated low-level Go binding around an user-defined struct.
type ILiquidationDataProviderCollateralFullInfo struct {
	AToken                          common.Address
	CollateralBalance               *big.Int
	CollateralBalanceInBaseCurrency *big.Int
	Price                           *big.Int
	AssetUnit                       *big.Int
}

// ILiquidationDataProviderDebtFullInfo is an auto generated low-level Go binding around an user-defined struct.
type ILiquidationDataProviderDebtFullInfo struct {
	VariableDebtToken         common.Address
	DebtBalance               *big.Int
	DebtBalanceInBaseCurrency *big.Int
	Price                     *big.Int
	AssetUnit                 *big.Int
}

// ILiquidationDataProviderLiquidationInfo is an auto generated low-level Go binding around an user-defined struct.
type ILiquidationDataProviderLiquidationInfo struct {
	UserInfo                      ILiquidationDataProviderUserPositionFullInfo
	CollateralInfo                ILiquidationDataProviderCollateralFullInfo
	DebtInfo                      ILiquidationDataProviderDebtFullInfo
	MaxCollateralToLiquidate      *big.Int
	MaxDebtToLiquidate            *big.Int
	LiquidationProtocolFee        *big.Int
	AmountToPassToLiquidationCall *big.Int
}

// ILiquidationDataProviderUserPositionFullInfo is an auto generated low-level Go binding around an user-defined struct.
type ILiquidationDataProviderUserPositionFullInfo struct {
	TotalCollateralInBaseCurrency  *big.Int
	TotalDebtInBaseCurrency        *big.Int
	AvailableBorrowsInBaseCurrency *big.Int
	CurrentLiquidationThreshold    *big.Int
	Ltv                            *big.Int
	HealthFactor                   *big.Int
}

// LiquidationProviderMetaData contains all meta data concerning the LiquidationProvider contract.
var LiquidationProviderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addressesProvider\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADDRESSES_PROVIDER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolAddressesProvider\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POOL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollateralFullInfo\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralAsset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structILiquidationDataProvider.CollateralFullInfo\",\"components\":[{\"name\":\"aToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralBalanceInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetUnit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDebtFullInfo\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtAsset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structILiquidationDataProvider.DebtFullInfo\",\"components\":[{\"name\":\"variableDebtToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debtBalanceInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetUnit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLiquidationInfo\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralAsset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtAsset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structILiquidationDataProvider.LiquidationInfo\",\"components\":[{\"name\":\"userInfo\",\"type\":\"tuple\",\"internalType\":\"structILiquidationDataProvider.UserPositionFullInfo\",\"components\":[{\"name\":\"totalCollateralInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableBorrowsInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ltv\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"collateralInfo\",\"type\":\"tuple\",\"internalType\":\"structILiquidationDataProvider.CollateralFullInfo\",\"components\":[{\"name\":\"aToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralBalanceInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetUnit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"debtInfo\",\"type\":\"tuple\",\"internalType\":\"structILiquidationDataProvider.DebtFullInfo\",\"components\":[{\"name\":\"variableDebtToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debtBalanceInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetUnit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"maxCollateralToLiquidate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDebtToLiquidate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liquidationProtocolFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountToPassToLiquidationCall\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLiquidationInfo\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralAsset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtAsset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtLiquidationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structILiquidationDataProvider.LiquidationInfo\",\"components\":[{\"name\":\"userInfo\",\"type\":\"tuple\",\"internalType\":\"structILiquidationDataProvider.UserPositionFullInfo\",\"components\":[{\"name\":\"totalCollateralInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableBorrowsInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ltv\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"collateralInfo\",\"type\":\"tuple\",\"internalType\":\"structILiquidationDataProvider.CollateralFullInfo\",\"components\":[{\"name\":\"aToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralBalanceInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetUnit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"debtInfo\",\"type\":\"tuple\",\"internalType\":\"structILiquidationDataProvider.DebtFullInfo\",\"components\":[{\"name\":\"variableDebtToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debtBalanceInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetUnit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"maxCollateralToLiquidate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDebtToLiquidate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liquidationProtocolFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountToPassToLiquidationCall\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserPositionFullInfo\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structILiquidationDataProvider.UserPositionFullInfo\",\"components\":[{\"name\":\"totalCollateralInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableBorrowsInBaseCurrency\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ltv\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidReserveIndex\",\"inputs\":[]}]",
}

// LiquidationProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidationProviderMetaData.ABI instead.
var LiquidationProviderABI = LiquidationProviderMetaData.ABI

// LiquidationProvider is an auto generated Go binding around an Ethereum contract.
type LiquidationProvider struct {
	LiquidationProviderCaller     // Read-only binding to the contract
	LiquidationProviderTransactor // Write-only binding to the contract
	LiquidationProviderFilterer   // Log filterer for contract events
}

// LiquidationProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidationProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidationProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidationProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidationProviderSession struct {
	Contract     *LiquidationProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LiquidationProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidationProviderCallerSession struct {
	Contract *LiquidationProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// LiquidationProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidationProviderTransactorSession struct {
	Contract     *LiquidationProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LiquidationProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidationProviderRaw struct {
	Contract *LiquidationProvider // Generic contract binding to access the raw methods on
}

// LiquidationProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidationProviderCallerRaw struct {
	Contract *LiquidationProviderCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidationProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidationProviderTransactorRaw struct {
	Contract *LiquidationProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidationProvider creates a new instance of LiquidationProvider, bound to a specific deployed contract.
func NewLiquidationProvider(address common.Address, backend bind.ContractBackend) (*LiquidationProvider, error) {
	contract, err := bindLiquidationProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidationProvider{LiquidationProviderCaller: LiquidationProviderCaller{contract: contract}, LiquidationProviderTransactor: LiquidationProviderTransactor{contract: contract}, LiquidationProviderFilterer: LiquidationProviderFilterer{contract: contract}}, nil
}

// NewLiquidationProviderCaller creates a new read-only instance of LiquidationProvider, bound to a specific deployed contract.
func NewLiquidationProviderCaller(address common.Address, caller bind.ContractCaller) (*LiquidationProviderCaller, error) {
	contract, err := bindLiquidationProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidationProviderCaller{contract: contract}, nil
}

// NewLiquidationProviderTransactor creates a new write-only instance of LiquidationProvider, bound to a specific deployed contract.
func NewLiquidationProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidationProviderTransactor, error) {
	contract, err := bindLiquidationProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidationProviderTransactor{contract: contract}, nil
}

// NewLiquidationProviderFilterer creates a new log filterer instance of LiquidationProvider, bound to a specific deployed contract.
func NewLiquidationProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidationProviderFilterer, error) {
	contract, err := bindLiquidationProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidationProviderFilterer{contract: contract}, nil
}

// bindLiquidationProvider binds a generic wrapper to an already deployed contract.
func bindLiquidationProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquidationProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidationProvider *LiquidationProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidationProvider.Contract.LiquidationProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidationProvider *LiquidationProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationProvider.Contract.LiquidationProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidationProvider *LiquidationProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidationProvider.Contract.LiquidationProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidationProvider *LiquidationProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidationProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidationProvider *LiquidationProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidationProvider *LiquidationProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidationProvider.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_LiquidationProvider *LiquidationProviderCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationProvider.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_LiquidationProvider *LiquidationProviderSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _LiquidationProvider.Contract.ADDRESSESPROVIDER(&_LiquidationProvider.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_LiquidationProvider *LiquidationProviderCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _LiquidationProvider.Contract.ADDRESSESPROVIDER(&_LiquidationProvider.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_LiquidationProvider *LiquidationProviderCaller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationProvider.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_LiquidationProvider *LiquidationProviderSession) POOL() (common.Address, error) {
	return _LiquidationProvider.Contract.POOL(&_LiquidationProvider.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_LiquidationProvider *LiquidationProviderCallerSession) POOL() (common.Address, error) {
	return _LiquidationProvider.Contract.POOL(&_LiquidationProvider.CallOpts)
}

// GetCollateralFullInfo is a free data retrieval call binding the contract method 0x0a2d69b4.
//
// Solidity: function getCollateralFullInfo(address user, address collateralAsset) view returns((address,uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderCaller) GetCollateralFullInfo(opts *bind.CallOpts, user common.Address, collateralAsset common.Address) (ILiquidationDataProviderCollateralFullInfo, error) {
	var out []interface{}
	err := _LiquidationProvider.contract.Call(opts, &out, "getCollateralFullInfo", user, collateralAsset)

	if err != nil {
		return *new(ILiquidationDataProviderCollateralFullInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ILiquidationDataProviderCollateralFullInfo)).(*ILiquidationDataProviderCollateralFullInfo)

	return out0, err

}

// GetCollateralFullInfo is a free data retrieval call binding the contract method 0x0a2d69b4.
//
// Solidity: function getCollateralFullInfo(address user, address collateralAsset) view returns((address,uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderSession) GetCollateralFullInfo(user common.Address, collateralAsset common.Address) (ILiquidationDataProviderCollateralFullInfo, error) {
	return _LiquidationProvider.Contract.GetCollateralFullInfo(&_LiquidationProvider.CallOpts, user, collateralAsset)
}

// GetCollateralFullInfo is a free data retrieval call binding the contract method 0x0a2d69b4.
//
// Solidity: function getCollateralFullInfo(address user, address collateralAsset) view returns((address,uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderCallerSession) GetCollateralFullInfo(user common.Address, collateralAsset common.Address) (ILiquidationDataProviderCollateralFullInfo, error) {
	return _LiquidationProvider.Contract.GetCollateralFullInfo(&_LiquidationProvider.CallOpts, user, collateralAsset)
}

// GetDebtFullInfo is a free data retrieval call binding the contract method 0x2d8c532c.
//
// Solidity: function getDebtFullInfo(address user, address debtAsset) view returns((address,uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderCaller) GetDebtFullInfo(opts *bind.CallOpts, user common.Address, debtAsset common.Address) (ILiquidationDataProviderDebtFullInfo, error) {
	var out []interface{}
	err := _LiquidationProvider.contract.Call(opts, &out, "getDebtFullInfo", user, debtAsset)

	if err != nil {
		return *new(ILiquidationDataProviderDebtFullInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ILiquidationDataProviderDebtFullInfo)).(*ILiquidationDataProviderDebtFullInfo)

	return out0, err

}

// GetDebtFullInfo is a free data retrieval call binding the contract method 0x2d8c532c.
//
// Solidity: function getDebtFullInfo(address user, address debtAsset) view returns((address,uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderSession) GetDebtFullInfo(user common.Address, debtAsset common.Address) (ILiquidationDataProviderDebtFullInfo, error) {
	return _LiquidationProvider.Contract.GetDebtFullInfo(&_LiquidationProvider.CallOpts, user, debtAsset)
}

// GetDebtFullInfo is a free data retrieval call binding the contract method 0x2d8c532c.
//
// Solidity: function getDebtFullInfo(address user, address debtAsset) view returns((address,uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderCallerSession) GetDebtFullInfo(user common.Address, debtAsset common.Address) (ILiquidationDataProviderDebtFullInfo, error) {
	return _LiquidationProvider.Contract.GetDebtFullInfo(&_LiquidationProvider.CallOpts, user, debtAsset)
}

// GetLiquidationInfo is a free data retrieval call binding the contract method 0x9db9ddcf.
//
// Solidity: function getLiquidationInfo(address user, address collateralAsset, address debtAsset) view returns(((uint256,uint256,uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderCaller) GetLiquidationInfo(opts *bind.CallOpts, user common.Address, collateralAsset common.Address, debtAsset common.Address) (ILiquidationDataProviderLiquidationInfo, error) {
	var out []interface{}
	err := _LiquidationProvider.contract.Call(opts, &out, "getLiquidationInfo", user, collateralAsset, debtAsset)

	if err != nil {
		return *new(ILiquidationDataProviderLiquidationInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ILiquidationDataProviderLiquidationInfo)).(*ILiquidationDataProviderLiquidationInfo)

	return out0, err

}

// GetLiquidationInfo is a free data retrieval call binding the contract method 0x9db9ddcf.
//
// Solidity: function getLiquidationInfo(address user, address collateralAsset, address debtAsset) view returns(((uint256,uint256,uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderSession) GetLiquidationInfo(user common.Address, collateralAsset common.Address, debtAsset common.Address) (ILiquidationDataProviderLiquidationInfo, error) {
	return _LiquidationProvider.Contract.GetLiquidationInfo(&_LiquidationProvider.CallOpts, user, collateralAsset, debtAsset)
}

// GetLiquidationInfo is a free data retrieval call binding the contract method 0x9db9ddcf.
//
// Solidity: function getLiquidationInfo(address user, address collateralAsset, address debtAsset) view returns(((uint256,uint256,uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderCallerSession) GetLiquidationInfo(user common.Address, collateralAsset common.Address, debtAsset common.Address) (ILiquidationDataProviderLiquidationInfo, error) {
	return _LiquidationProvider.Contract.GetLiquidationInfo(&_LiquidationProvider.CallOpts, user, collateralAsset, debtAsset)
}

// GetLiquidationInfo0 is a free data retrieval call binding the contract method 0xce5874c0.
//
// Solidity: function getLiquidationInfo(address user, address collateralAsset, address debtAsset, uint256 debtLiquidationAmount) view returns(((uint256,uint256,uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderCaller) GetLiquidationInfo0(opts *bind.CallOpts, user common.Address, collateralAsset common.Address, debtAsset common.Address, debtLiquidationAmount *big.Int) (ILiquidationDataProviderLiquidationInfo, error) {
	var out []interface{}
	err := _LiquidationProvider.contract.Call(opts, &out, "getLiquidationInfo0", user, collateralAsset, debtAsset, debtLiquidationAmount)

	if err != nil {
		return *new(ILiquidationDataProviderLiquidationInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ILiquidationDataProviderLiquidationInfo)).(*ILiquidationDataProviderLiquidationInfo)

	return out0, err

}

// GetLiquidationInfo0 is a free data retrieval call binding the contract method 0xce5874c0.
//
// Solidity: function getLiquidationInfo(address user, address collateralAsset, address debtAsset, uint256 debtLiquidationAmount) view returns(((uint256,uint256,uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderSession) GetLiquidationInfo0(user common.Address, collateralAsset common.Address, debtAsset common.Address, debtLiquidationAmount *big.Int) (ILiquidationDataProviderLiquidationInfo, error) {
	return _LiquidationProvider.Contract.GetLiquidationInfo0(&_LiquidationProvider.CallOpts, user, collateralAsset, debtAsset, debtLiquidationAmount)
}

// GetLiquidationInfo0 is a free data retrieval call binding the contract method 0xce5874c0.
//
// Solidity: function getLiquidationInfo(address user, address collateralAsset, address debtAsset, uint256 debtLiquidationAmount) view returns(((uint256,uint256,uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256,uint256),uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderCallerSession) GetLiquidationInfo0(user common.Address, collateralAsset common.Address, debtAsset common.Address, debtLiquidationAmount *big.Int) (ILiquidationDataProviderLiquidationInfo, error) {
	return _LiquidationProvider.Contract.GetLiquidationInfo0(&_LiquidationProvider.CallOpts, user, collateralAsset, debtAsset, debtLiquidationAmount)
}

// GetUserPositionFullInfo is a free data retrieval call binding the contract method 0xfa1582d5.
//
// Solidity: function getUserPositionFullInfo(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderCaller) GetUserPositionFullInfo(opts *bind.CallOpts, user common.Address) (ILiquidationDataProviderUserPositionFullInfo, error) {
	var out []interface{}
	err := _LiquidationProvider.contract.Call(opts, &out, "getUserPositionFullInfo", user)

	if err != nil {
		return *new(ILiquidationDataProviderUserPositionFullInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ILiquidationDataProviderUserPositionFullInfo)).(*ILiquidationDataProviderUserPositionFullInfo)

	return out0, err

}

// GetUserPositionFullInfo is a free data retrieval call binding the contract method 0xfa1582d5.
//
// Solidity: function getUserPositionFullInfo(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderSession) GetUserPositionFullInfo(user common.Address) (ILiquidationDataProviderUserPositionFullInfo, error) {
	return _LiquidationProvider.Contract.GetUserPositionFullInfo(&_LiquidationProvider.CallOpts, user)
}

// GetUserPositionFullInfo is a free data retrieval call binding the contract method 0xfa1582d5.
//
// Solidity: function getUserPositionFullInfo(address user) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_LiquidationProvider *LiquidationProviderCallerSession) GetUserPositionFullInfo(user common.Address) (ILiquidationDataProviderUserPositionFullInfo, error) {
	return _LiquidationProvider.Contract.GetUserPositionFullInfo(&_LiquidationProvider.CallOpts, user)
}
