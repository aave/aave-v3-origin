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

// LiqParams is an auto generated low-level Go binding around an user-defined struct.
type LiqParams struct {
	User            common.Address
	DebtAsset       common.Address
	CollateralAsset common.Address
	DebtToCover     *big.Int
}

// LiquidationWalletMetaData contains all meta data concerning the LiquidationWallet contract.
var LiquidationWalletMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\",\"internalType\":\"contractIL2Pool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"AAVE_V3_POOL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIL2Pool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EOF\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"V3\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"liquidateWithV3\",\"inputs\":[{\"name\":\"liqParams\",\"type\":\"tuple\",\"internalType\":\"structLiqParams\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtAsset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralAsset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtToCover\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"path\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operators\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"uniswapV3SwapCallback\",\"inputs\":[{\"name\":\"amount0Delta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"amount1Delta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawERC20\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawETH\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientProfit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAction\",\"inputs\":[{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidCallbackSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// LiquidationWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidationWalletMetaData.ABI instead.
var LiquidationWalletABI = LiquidationWalletMetaData.ABI

// LiquidationWallet is an auto generated Go binding around an Ethereum contract.
type LiquidationWallet struct {
	LiquidationWalletCaller     // Read-only binding to the contract
	LiquidationWalletTransactor // Write-only binding to the contract
	LiquidationWalletFilterer   // Log filterer for contract events
}

// LiquidationWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidationWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidationWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidationWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidationWalletSession struct {
	Contract     *LiquidationWallet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LiquidationWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidationWalletCallerSession struct {
	Contract *LiquidationWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LiquidationWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidationWalletTransactorSession struct {
	Contract     *LiquidationWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LiquidationWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidationWalletRaw struct {
	Contract *LiquidationWallet // Generic contract binding to access the raw methods on
}

// LiquidationWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidationWalletCallerRaw struct {
	Contract *LiquidationWalletCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidationWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidationWalletTransactorRaw struct {
	Contract *LiquidationWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidationWallet creates a new instance of LiquidationWallet, bound to a specific deployed contract.
func NewLiquidationWallet(address common.Address, backend bind.ContractBackend) (*LiquidationWallet, error) {
	contract, err := bindLiquidationWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidationWallet{LiquidationWalletCaller: LiquidationWalletCaller{contract: contract}, LiquidationWalletTransactor: LiquidationWalletTransactor{contract: contract}, LiquidationWalletFilterer: LiquidationWalletFilterer{contract: contract}}, nil
}

// NewLiquidationWalletCaller creates a new read-only instance of LiquidationWallet, bound to a specific deployed contract.
func NewLiquidationWalletCaller(address common.Address, caller bind.ContractCaller) (*LiquidationWalletCaller, error) {
	contract, err := bindLiquidationWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidationWalletCaller{contract: contract}, nil
}

// NewLiquidationWalletTransactor creates a new write-only instance of LiquidationWallet, bound to a specific deployed contract.
func NewLiquidationWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidationWalletTransactor, error) {
	contract, err := bindLiquidationWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidationWalletTransactor{contract: contract}, nil
}

// NewLiquidationWalletFilterer creates a new log filterer instance of LiquidationWallet, bound to a specific deployed contract.
func NewLiquidationWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidationWalletFilterer, error) {
	contract, err := bindLiquidationWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidationWalletFilterer{contract: contract}, nil
}

// bindLiquidationWallet binds a generic wrapper to an already deployed contract.
func bindLiquidationWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquidationWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidationWallet *LiquidationWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidationWallet.Contract.LiquidationWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidationWallet *LiquidationWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.LiquidationWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidationWallet *LiquidationWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.LiquidationWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidationWallet *LiquidationWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidationWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidationWallet *LiquidationWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidationWallet *LiquidationWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.contract.Transact(opts, method, params...)
}

// AAVEV3POOL is a free data retrieval call binding the contract method 0x3b303705.
//
// Solidity: function AAVE_V3_POOL() view returns(address)
func (_LiquidationWallet *LiquidationWalletCaller) AAVEV3POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationWallet.contract.Call(opts, &out, "AAVE_V3_POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AAVEV3POOL is a free data retrieval call binding the contract method 0x3b303705.
//
// Solidity: function AAVE_V3_POOL() view returns(address)
func (_LiquidationWallet *LiquidationWalletSession) AAVEV3POOL() (common.Address, error) {
	return _LiquidationWallet.Contract.AAVEV3POOL(&_LiquidationWallet.CallOpts)
}

// AAVEV3POOL is a free data retrieval call binding the contract method 0x3b303705.
//
// Solidity: function AAVE_V3_POOL() view returns(address)
func (_LiquidationWallet *LiquidationWalletCallerSession) AAVEV3POOL() (common.Address, error) {
	return _LiquidationWallet.Contract.AAVEV3POOL(&_LiquidationWallet.CallOpts)
}

// EOF is a free data retrieval call binding the contract method 0x65d22ffa.
//
// Solidity: function EOF() view returns(uint8)
func (_LiquidationWallet *LiquidationWalletCaller) EOF(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LiquidationWallet.contract.Call(opts, &out, "EOF")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// EOF is a free data retrieval call binding the contract method 0x65d22ffa.
//
// Solidity: function EOF() view returns(uint8)
func (_LiquidationWallet *LiquidationWalletSession) EOF() (uint8, error) {
	return _LiquidationWallet.Contract.EOF(&_LiquidationWallet.CallOpts)
}

// EOF is a free data retrieval call binding the contract method 0x65d22ffa.
//
// Solidity: function EOF() view returns(uint8)
func (_LiquidationWallet *LiquidationWalletCallerSession) EOF() (uint8, error) {
	return _LiquidationWallet.Contract.EOF(&_LiquidationWallet.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LiquidationWallet *LiquidationWalletCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LiquidationWallet.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LiquidationWallet *LiquidationWalletSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LiquidationWallet.Contract.UPGRADEINTERFACEVERSION(&_LiquidationWallet.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_LiquidationWallet *LiquidationWalletCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _LiquidationWallet.Contract.UPGRADEINTERFACEVERSION(&_LiquidationWallet.CallOpts)
}

// V3 is a free data retrieval call binding the contract method 0x2fb42d70.
//
// Solidity: function V3() view returns(uint8)
func (_LiquidationWallet *LiquidationWalletCaller) V3(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LiquidationWallet.contract.Call(opts, &out, "V3")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// V3 is a free data retrieval call binding the contract method 0x2fb42d70.
//
// Solidity: function V3() view returns(uint8)
func (_LiquidationWallet *LiquidationWalletSession) V3() (uint8, error) {
	return _LiquidationWallet.Contract.V3(&_LiquidationWallet.CallOpts)
}

// V3 is a free data retrieval call binding the contract method 0x2fb42d70.
//
// Solidity: function V3() view returns(uint8)
func (_LiquidationWallet *LiquidationWalletCallerSession) V3() (uint8, error) {
	return _LiquidationWallet.Contract.V3(&_LiquidationWallet.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_LiquidationWallet *LiquidationWalletCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LiquidationWallet.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_LiquidationWallet *LiquidationWalletSession) Operators(arg0 common.Address) (bool, error) {
	return _LiquidationWallet.Contract.Operators(&_LiquidationWallet.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_LiquidationWallet *LiquidationWalletCallerSession) Operators(arg0 common.Address) (bool, error) {
	return _LiquidationWallet.Contract.Operators(&_LiquidationWallet.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidationWallet *LiquidationWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidationWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidationWallet *LiquidationWalletSession) Owner() (common.Address, error) {
	return _LiquidationWallet.Contract.Owner(&_LiquidationWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidationWallet *LiquidationWalletCallerSession) Owner() (common.Address, error) {
	return _LiquidationWallet.Contract.Owner(&_LiquidationWallet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LiquidationWallet *LiquidationWalletCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LiquidationWallet.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LiquidationWallet *LiquidationWalletSession) ProxiableUUID() ([32]byte, error) {
	return _LiquidationWallet.Contract.ProxiableUUID(&_LiquidationWallet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LiquidationWallet *LiquidationWalletCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LiquidationWallet.Contract.ProxiableUUID(&_LiquidationWallet.CallOpts)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_LiquidationWallet *LiquidationWalletTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _LiquidationWallet.contract.Transact(opts, "addOperator", operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_LiquidationWallet *LiquidationWalletSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.AddOperator(&_LiquidationWallet.TransactOpts, operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_LiquidationWallet *LiquidationWalletTransactorSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.AddOperator(&_LiquidationWallet.TransactOpts, operator)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address asset, address spender, uint256 amount) returns()
func (_LiquidationWallet *LiquidationWalletTransactor) Approve(opts *bind.TransactOpts, asset common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidationWallet.contract.Transact(opts, "approve", asset, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address asset, address spender, uint256 amount) returns()
func (_LiquidationWallet *LiquidationWalletSession) Approve(asset common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.Approve(&_LiquidationWallet.TransactOpts, asset, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address asset, address spender, uint256 amount) returns()
func (_LiquidationWallet *LiquidationWalletTransactorSession) Approve(asset common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.Approve(&_LiquidationWallet.TransactOpts, asset, spender, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_LiquidationWallet *LiquidationWalletTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _LiquidationWallet.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_LiquidationWallet *LiquidationWalletSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.Initialize(&_LiquidationWallet.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_LiquidationWallet *LiquidationWalletTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.Initialize(&_LiquidationWallet.TransactOpts, _owner)
}

// LiquidateWithV3 is a paid mutator transaction binding the contract method 0xf8341e31.
//
// Solidity: function liquidateWithV3((address,address,address,uint256) liqParams, bytes path) returns()
func (_LiquidationWallet *LiquidationWalletTransactor) LiquidateWithV3(opts *bind.TransactOpts, liqParams LiqParams, path []byte) (*types.Transaction, error) {
	return _LiquidationWallet.contract.Transact(opts, "liquidateWithV3", liqParams, path)
}

// LiquidateWithV3 is a paid mutator transaction binding the contract method 0xf8341e31.
//
// Solidity: function liquidateWithV3((address,address,address,uint256) liqParams, bytes path) returns()
func (_LiquidationWallet *LiquidationWalletSession) LiquidateWithV3(liqParams LiqParams, path []byte) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.LiquidateWithV3(&_LiquidationWallet.TransactOpts, liqParams, path)
}

// LiquidateWithV3 is a paid mutator transaction binding the contract method 0xf8341e31.
//
// Solidity: function liquidateWithV3((address,address,address,uint256) liqParams, bytes path) returns()
func (_LiquidationWallet *LiquidationWalletTransactorSession) LiquidateWithV3(liqParams LiqParams, path []byte) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.LiquidateWithV3(&_LiquidationWallet.TransactOpts, liqParams, path)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_LiquidationWallet *LiquidationWalletTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _LiquidationWallet.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_LiquidationWallet *LiquidationWalletSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.RemoveOperator(&_LiquidationWallet.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_LiquidationWallet *LiquidationWalletTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.RemoveOperator(&_LiquidationWallet.TransactOpts, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidationWallet *LiquidationWalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationWallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidationWallet *LiquidationWalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidationWallet.Contract.RenounceOwnership(&_LiquidationWallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidationWallet *LiquidationWalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidationWallet.Contract.RenounceOwnership(&_LiquidationWallet.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidationWallet *LiquidationWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LiquidationWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidationWallet *LiquidationWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.TransferOwnership(&_LiquidationWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidationWallet *LiquidationWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.TransferOwnership(&_LiquidationWallet.TransactOpts, newOwner)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_LiquidationWallet *LiquidationWalletTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _LiquidationWallet.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_LiquidationWallet *LiquidationWalletSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.UniswapV3SwapCallback(&_LiquidationWallet.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_LiquidationWallet *LiquidationWalletTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.UniswapV3SwapCallback(&_LiquidationWallet.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LiquidationWallet *LiquidationWalletTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LiquidationWallet.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LiquidationWallet *LiquidationWalletSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.UpgradeToAndCall(&_LiquidationWallet.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LiquidationWallet *LiquidationWalletTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.UpgradeToAndCall(&_LiquidationWallet.TransactOpts, newImplementation, data)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address asset, uint256 amount) returns()
func (_LiquidationWallet *LiquidationWalletTransactor) WithdrawERC20(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidationWallet.contract.Transact(opts, "withdrawERC20", asset, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address asset, uint256 amount) returns()
func (_LiquidationWallet *LiquidationWalletSession) WithdrawERC20(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.WithdrawERC20(&_LiquidationWallet.TransactOpts, asset, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address asset, uint256 amount) returns()
func (_LiquidationWallet *LiquidationWalletTransactorSession) WithdrawERC20(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.WithdrawERC20(&_LiquidationWallet.TransactOpts, asset, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_LiquidationWallet *LiquidationWalletTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LiquidationWallet.contract.Transact(opts, "withdrawETH", amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_LiquidationWallet *LiquidationWalletSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.WithdrawETH(&_LiquidationWallet.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_LiquidationWallet *LiquidationWalletTransactorSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _LiquidationWallet.Contract.WithdrawETH(&_LiquidationWallet.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidationWallet *LiquidationWalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationWallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidationWallet *LiquidationWalletSession) Receive() (*types.Transaction, error) {
	return _LiquidationWallet.Contract.Receive(&_LiquidationWallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidationWallet *LiquidationWalletTransactorSession) Receive() (*types.Transaction, error) {
	return _LiquidationWallet.Contract.Receive(&_LiquidationWallet.TransactOpts)
}

// LiquidationWalletInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LiquidationWallet contract.
type LiquidationWalletInitializedIterator struct {
	Event *LiquidationWalletInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidationWalletInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidationWalletInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidationWalletInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidationWalletInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidationWalletInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidationWalletInitialized represents a Initialized event raised by the LiquidationWallet contract.
type LiquidationWalletInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LiquidationWallet *LiquidationWalletFilterer) FilterInitialized(opts *bind.FilterOpts) (*LiquidationWalletInitializedIterator, error) {

	logs, sub, err := _LiquidationWallet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LiquidationWalletInitializedIterator{contract: _LiquidationWallet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LiquidationWallet *LiquidationWalletFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LiquidationWalletInitialized) (event.Subscription, error) {

	logs, sub, err := _LiquidationWallet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidationWalletInitialized)
				if err := _LiquidationWallet.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LiquidationWallet *LiquidationWalletFilterer) ParseInitialized(log types.Log) (*LiquidationWalletInitialized, error) {
	event := new(LiquidationWalletInitialized)
	if err := _LiquidationWallet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidationWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LiquidationWallet contract.
type LiquidationWalletOwnershipTransferredIterator struct {
	Event *LiquidationWalletOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidationWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidationWalletOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidationWalletOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidationWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidationWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidationWalletOwnershipTransferred represents a OwnershipTransferred event raised by the LiquidationWallet contract.
type LiquidationWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidationWallet *LiquidationWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquidationWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidationWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidationWalletOwnershipTransferredIterator{contract: _LiquidationWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidationWallet *LiquidationWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquidationWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidationWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidationWalletOwnershipTransferred)
				if err := _LiquidationWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidationWallet *LiquidationWalletFilterer) ParseOwnershipTransferred(log types.Log) (*LiquidationWalletOwnershipTransferred, error) {
	event := new(LiquidationWalletOwnershipTransferred)
	if err := _LiquidationWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidationWalletUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LiquidationWallet contract.
type LiquidationWalletUpgradedIterator struct {
	Event *LiquidationWalletUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidationWalletUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidationWalletUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidationWalletUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidationWalletUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidationWalletUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidationWalletUpgraded represents a Upgraded event raised by the LiquidationWallet contract.
type LiquidationWalletUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LiquidationWallet *LiquidationWalletFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LiquidationWalletUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LiquidationWallet.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LiquidationWalletUpgradedIterator{contract: _LiquidationWallet.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LiquidationWallet *LiquidationWalletFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LiquidationWalletUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LiquidationWallet.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidationWalletUpgraded)
				if err := _LiquidationWallet.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LiquidationWallet *LiquidationWalletFilterer) ParseUpgraded(log types.Log) (*LiquidationWalletUpgraded, error) {
	event := new(LiquidationWalletUpgraded)
	if err := _LiquidationWallet.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
