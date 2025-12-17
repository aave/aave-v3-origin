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

// AaveOracleMetaData contains all meta data concerning the AaveOracle contract.
var AaveOracleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ADDRESSES_PROVIDER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolAddressesProvider\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BASE_CURRENCY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BASE_CURRENCY_UNIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAssetPrice\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAssetsPrices\",\"inputs\":[{\"name\":\"assets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFallbackOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceOfAsset\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAssetSources\",\"inputs\":[{\"name\":\"assets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"sources\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFallbackOracle\",\"inputs\":[{\"name\":\"fallbackOracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AssetSourceUpdated\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"source\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BaseCurrencySet\",\"inputs\":[{\"name\":\"baseCurrency\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"baseCurrencyUnit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FallbackOracleUpdated\",\"inputs\":[{\"name\":\"fallbackOracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// AaveOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveOracleMetaData.ABI instead.
var AaveOracleABI = AaveOracleMetaData.ABI

// AaveOracle is an auto generated Go binding around an Ethereum contract.
type AaveOracle struct {
	AaveOracleCaller     // Read-only binding to the contract
	AaveOracleTransactor // Write-only binding to the contract
	AaveOracleFilterer   // Log filterer for contract events
}

// AaveOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveOracleSession struct {
	Contract     *AaveOracle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveOracleCallerSession struct {
	Contract *AaveOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AaveOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveOracleTransactorSession struct {
	Contract     *AaveOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AaveOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveOracleRaw struct {
	Contract *AaveOracle // Generic contract binding to access the raw methods on
}

// AaveOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveOracleCallerRaw struct {
	Contract *AaveOracleCaller // Generic read-only contract binding to access the raw methods on
}

// AaveOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveOracleTransactorRaw struct {
	Contract *AaveOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveOracle creates a new instance of AaveOracle, bound to a specific deployed contract.
func NewAaveOracle(address common.Address, backend bind.ContractBackend) (*AaveOracle, error) {
	contract, err := bindAaveOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveOracle{AaveOracleCaller: AaveOracleCaller{contract: contract}, AaveOracleTransactor: AaveOracleTransactor{contract: contract}, AaveOracleFilterer: AaveOracleFilterer{contract: contract}}, nil
}

// NewAaveOracleCaller creates a new read-only instance of AaveOracle, bound to a specific deployed contract.
func NewAaveOracleCaller(address common.Address, caller bind.ContractCaller) (*AaveOracleCaller, error) {
	contract, err := bindAaveOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveOracleCaller{contract: contract}, nil
}

// NewAaveOracleTransactor creates a new write-only instance of AaveOracle, bound to a specific deployed contract.
func NewAaveOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveOracleTransactor, error) {
	contract, err := bindAaveOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveOracleTransactor{contract: contract}, nil
}

// NewAaveOracleFilterer creates a new log filterer instance of AaveOracle, bound to a specific deployed contract.
func NewAaveOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveOracleFilterer, error) {
	contract, err := bindAaveOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveOracleFilterer{contract: contract}, nil
}

// bindAaveOracle binds a generic wrapper to an already deployed contract.
func bindAaveOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveOracle *AaveOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveOracle.Contract.AaveOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveOracle *AaveOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveOracle.Contract.AaveOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveOracle *AaveOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveOracle.Contract.AaveOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveOracle *AaveOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveOracle *AaveOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveOracle *AaveOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveOracle.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveOracle *AaveOracleCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveOracle.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveOracle *AaveOracleSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveOracle.Contract.ADDRESSESPROVIDER(&_AaveOracle.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveOracle *AaveOracleCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveOracle.Contract.ADDRESSESPROVIDER(&_AaveOracle.CallOpts)
}

// BASECURRENCY is a free data retrieval call binding the contract method 0xe19f4700.
//
// Solidity: function BASE_CURRENCY() view returns(address)
func (_AaveOracle *AaveOracleCaller) BASECURRENCY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveOracle.contract.Call(opts, &out, "BASE_CURRENCY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASECURRENCY is a free data retrieval call binding the contract method 0xe19f4700.
//
// Solidity: function BASE_CURRENCY() view returns(address)
func (_AaveOracle *AaveOracleSession) BASECURRENCY() (common.Address, error) {
	return _AaveOracle.Contract.BASECURRENCY(&_AaveOracle.CallOpts)
}

// BASECURRENCY is a free data retrieval call binding the contract method 0xe19f4700.
//
// Solidity: function BASE_CURRENCY() view returns(address)
func (_AaveOracle *AaveOracleCallerSession) BASECURRENCY() (common.Address, error) {
	return _AaveOracle.Contract.BASECURRENCY(&_AaveOracle.CallOpts)
}

// BASECURRENCYUNIT is a free data retrieval call binding the contract method 0x8c89b64f.
//
// Solidity: function BASE_CURRENCY_UNIT() view returns(uint256)
func (_AaveOracle *AaveOracleCaller) BASECURRENCYUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveOracle.contract.Call(opts, &out, "BASE_CURRENCY_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASECURRENCYUNIT is a free data retrieval call binding the contract method 0x8c89b64f.
//
// Solidity: function BASE_CURRENCY_UNIT() view returns(uint256)
func (_AaveOracle *AaveOracleSession) BASECURRENCYUNIT() (*big.Int, error) {
	return _AaveOracle.Contract.BASECURRENCYUNIT(&_AaveOracle.CallOpts)
}

// BASECURRENCYUNIT is a free data retrieval call binding the contract method 0x8c89b64f.
//
// Solidity: function BASE_CURRENCY_UNIT() view returns(uint256)
func (_AaveOracle *AaveOracleCallerSession) BASECURRENCYUNIT() (*big.Int, error) {
	return _AaveOracle.Contract.BASECURRENCYUNIT(&_AaveOracle.CallOpts)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_AaveOracle *AaveOracleCaller) GetAssetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveOracle.contract.Call(opts, &out, "getAssetPrice", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_AaveOracle *AaveOracleSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _AaveOracle.Contract.GetAssetPrice(&_AaveOracle.CallOpts, asset)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_AaveOracle *AaveOracleCallerSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _AaveOracle.Contract.GetAssetPrice(&_AaveOracle.CallOpts, asset)
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_AaveOracle *AaveOracleCaller) GetAssetsPrices(opts *bind.CallOpts, assets []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _AaveOracle.contract.Call(opts, &out, "getAssetsPrices", assets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_AaveOracle *AaveOracleSession) GetAssetsPrices(assets []common.Address) ([]*big.Int, error) {
	return _AaveOracle.Contract.GetAssetsPrices(&_AaveOracle.CallOpts, assets)
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_AaveOracle *AaveOracleCallerSession) GetAssetsPrices(assets []common.Address) ([]*big.Int, error) {
	return _AaveOracle.Contract.GetAssetsPrices(&_AaveOracle.CallOpts, assets)
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_AaveOracle *AaveOracleCaller) GetFallbackOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveOracle.contract.Call(opts, &out, "getFallbackOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_AaveOracle *AaveOracleSession) GetFallbackOracle() (common.Address, error) {
	return _AaveOracle.Contract.GetFallbackOracle(&_AaveOracle.CallOpts)
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_AaveOracle *AaveOracleCallerSession) GetFallbackOracle() (common.Address, error) {
	return _AaveOracle.Contract.GetFallbackOracle(&_AaveOracle.CallOpts)
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_AaveOracle *AaveOracleCaller) GetSourceOfAsset(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _AaveOracle.contract.Call(opts, &out, "getSourceOfAsset", asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_AaveOracle *AaveOracleSession) GetSourceOfAsset(asset common.Address) (common.Address, error) {
	return _AaveOracle.Contract.GetSourceOfAsset(&_AaveOracle.CallOpts, asset)
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_AaveOracle *AaveOracleCallerSession) GetSourceOfAsset(asset common.Address) (common.Address, error) {
	return _AaveOracle.Contract.GetSourceOfAsset(&_AaveOracle.CallOpts, asset)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_AaveOracle *AaveOracleTransactor) SetAssetSources(opts *bind.TransactOpts, assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _AaveOracle.contract.Transact(opts, "setAssetSources", assets, sources)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_AaveOracle *AaveOracleSession) SetAssetSources(assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _AaveOracle.Contract.SetAssetSources(&_AaveOracle.TransactOpts, assets, sources)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_AaveOracle *AaveOracleTransactorSession) SetAssetSources(assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _AaveOracle.Contract.SetAssetSources(&_AaveOracle.TransactOpts, assets, sources)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_AaveOracle *AaveOracleTransactor) SetFallbackOracle(opts *bind.TransactOpts, fallbackOracle common.Address) (*types.Transaction, error) {
	return _AaveOracle.contract.Transact(opts, "setFallbackOracle", fallbackOracle)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_AaveOracle *AaveOracleSession) SetFallbackOracle(fallbackOracle common.Address) (*types.Transaction, error) {
	return _AaveOracle.Contract.SetFallbackOracle(&_AaveOracle.TransactOpts, fallbackOracle)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_AaveOracle *AaveOracleTransactorSession) SetFallbackOracle(fallbackOracle common.Address) (*types.Transaction, error) {
	return _AaveOracle.Contract.SetFallbackOracle(&_AaveOracle.TransactOpts, fallbackOracle)
}

// AaveOracleAssetSourceUpdatedIterator is returned from FilterAssetSourceUpdated and is used to iterate over the raw logs and unpacked data for AssetSourceUpdated events raised by the AaveOracle contract.
type AaveOracleAssetSourceUpdatedIterator struct {
	Event *AaveOracleAssetSourceUpdated // Event containing the contract specifics and raw log

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
func (it *AaveOracleAssetSourceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveOracleAssetSourceUpdated)
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
		it.Event = new(AaveOracleAssetSourceUpdated)
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
func (it *AaveOracleAssetSourceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveOracleAssetSourceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveOracleAssetSourceUpdated represents a AssetSourceUpdated event raised by the AaveOracle contract.
type AaveOracleAssetSourceUpdated struct {
	Asset  common.Address
	Source common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAssetSourceUpdated is a free log retrieval operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_AaveOracle *AaveOracleFilterer) FilterAssetSourceUpdated(opts *bind.FilterOpts, asset []common.Address, source []common.Address) (*AaveOracleAssetSourceUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _AaveOracle.contract.FilterLogs(opts, "AssetSourceUpdated", assetRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return &AaveOracleAssetSourceUpdatedIterator{contract: _AaveOracle.contract, event: "AssetSourceUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetSourceUpdated is a free log subscription operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_AaveOracle *AaveOracleFilterer) WatchAssetSourceUpdated(opts *bind.WatchOpts, sink chan<- *AaveOracleAssetSourceUpdated, asset []common.Address, source []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _AaveOracle.contract.WatchLogs(opts, "AssetSourceUpdated", assetRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveOracleAssetSourceUpdated)
				if err := _AaveOracle.contract.UnpackLog(event, "AssetSourceUpdated", log); err != nil {
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

// ParseAssetSourceUpdated is a log parse operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_AaveOracle *AaveOracleFilterer) ParseAssetSourceUpdated(log types.Log) (*AaveOracleAssetSourceUpdated, error) {
	event := new(AaveOracleAssetSourceUpdated)
	if err := _AaveOracle.contract.UnpackLog(event, "AssetSourceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveOracleBaseCurrencySetIterator is returned from FilterBaseCurrencySet and is used to iterate over the raw logs and unpacked data for BaseCurrencySet events raised by the AaveOracle contract.
type AaveOracleBaseCurrencySetIterator struct {
	Event *AaveOracleBaseCurrencySet // Event containing the contract specifics and raw log

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
func (it *AaveOracleBaseCurrencySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveOracleBaseCurrencySet)
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
		it.Event = new(AaveOracleBaseCurrencySet)
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
func (it *AaveOracleBaseCurrencySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveOracleBaseCurrencySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveOracleBaseCurrencySet represents a BaseCurrencySet event raised by the AaveOracle contract.
type AaveOracleBaseCurrencySet struct {
	BaseCurrency     common.Address
	BaseCurrencyUnit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBaseCurrencySet is a free log retrieval operation binding the contract event 0xe27c4c1372396a3d15a9922f74f9dfc7c72b1ad6d63868470787249c356454c1.
//
// Solidity: event BaseCurrencySet(address indexed baseCurrency, uint256 baseCurrencyUnit)
func (_AaveOracle *AaveOracleFilterer) FilterBaseCurrencySet(opts *bind.FilterOpts, baseCurrency []common.Address) (*AaveOracleBaseCurrencySetIterator, error) {

	var baseCurrencyRule []interface{}
	for _, baseCurrencyItem := range baseCurrency {
		baseCurrencyRule = append(baseCurrencyRule, baseCurrencyItem)
	}

	logs, sub, err := _AaveOracle.contract.FilterLogs(opts, "BaseCurrencySet", baseCurrencyRule)
	if err != nil {
		return nil, err
	}
	return &AaveOracleBaseCurrencySetIterator{contract: _AaveOracle.contract, event: "BaseCurrencySet", logs: logs, sub: sub}, nil
}

// WatchBaseCurrencySet is a free log subscription operation binding the contract event 0xe27c4c1372396a3d15a9922f74f9dfc7c72b1ad6d63868470787249c356454c1.
//
// Solidity: event BaseCurrencySet(address indexed baseCurrency, uint256 baseCurrencyUnit)
func (_AaveOracle *AaveOracleFilterer) WatchBaseCurrencySet(opts *bind.WatchOpts, sink chan<- *AaveOracleBaseCurrencySet, baseCurrency []common.Address) (event.Subscription, error) {

	var baseCurrencyRule []interface{}
	for _, baseCurrencyItem := range baseCurrency {
		baseCurrencyRule = append(baseCurrencyRule, baseCurrencyItem)
	}

	logs, sub, err := _AaveOracle.contract.WatchLogs(opts, "BaseCurrencySet", baseCurrencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveOracleBaseCurrencySet)
				if err := _AaveOracle.contract.UnpackLog(event, "BaseCurrencySet", log); err != nil {
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

// ParseBaseCurrencySet is a log parse operation binding the contract event 0xe27c4c1372396a3d15a9922f74f9dfc7c72b1ad6d63868470787249c356454c1.
//
// Solidity: event BaseCurrencySet(address indexed baseCurrency, uint256 baseCurrencyUnit)
func (_AaveOracle *AaveOracleFilterer) ParseBaseCurrencySet(log types.Log) (*AaveOracleBaseCurrencySet, error) {
	event := new(AaveOracleBaseCurrencySet)
	if err := _AaveOracle.contract.UnpackLog(event, "BaseCurrencySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveOracleFallbackOracleUpdatedIterator is returned from FilterFallbackOracleUpdated and is used to iterate over the raw logs and unpacked data for FallbackOracleUpdated events raised by the AaveOracle contract.
type AaveOracleFallbackOracleUpdatedIterator struct {
	Event *AaveOracleFallbackOracleUpdated // Event containing the contract specifics and raw log

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
func (it *AaveOracleFallbackOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveOracleFallbackOracleUpdated)
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
		it.Event = new(AaveOracleFallbackOracleUpdated)
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
func (it *AaveOracleFallbackOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveOracleFallbackOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveOracleFallbackOracleUpdated represents a FallbackOracleUpdated event raised by the AaveOracle contract.
type AaveOracleFallbackOracleUpdated struct {
	FallbackOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFallbackOracleUpdated is a free log retrieval operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_AaveOracle *AaveOracleFilterer) FilterFallbackOracleUpdated(opts *bind.FilterOpts, fallbackOracle []common.Address) (*AaveOracleFallbackOracleUpdatedIterator, error) {

	var fallbackOracleRule []interface{}
	for _, fallbackOracleItem := range fallbackOracle {
		fallbackOracleRule = append(fallbackOracleRule, fallbackOracleItem)
	}

	logs, sub, err := _AaveOracle.contract.FilterLogs(opts, "FallbackOracleUpdated", fallbackOracleRule)
	if err != nil {
		return nil, err
	}
	return &AaveOracleFallbackOracleUpdatedIterator{contract: _AaveOracle.contract, event: "FallbackOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchFallbackOracleUpdated is a free log subscription operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_AaveOracle *AaveOracleFilterer) WatchFallbackOracleUpdated(opts *bind.WatchOpts, sink chan<- *AaveOracleFallbackOracleUpdated, fallbackOracle []common.Address) (event.Subscription, error) {

	var fallbackOracleRule []interface{}
	for _, fallbackOracleItem := range fallbackOracle {
		fallbackOracleRule = append(fallbackOracleRule, fallbackOracleItem)
	}

	logs, sub, err := _AaveOracle.contract.WatchLogs(opts, "FallbackOracleUpdated", fallbackOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveOracleFallbackOracleUpdated)
				if err := _AaveOracle.contract.UnpackLog(event, "FallbackOracleUpdated", log); err != nil {
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

// ParseFallbackOracleUpdated is a log parse operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_AaveOracle *AaveOracleFilterer) ParseFallbackOracleUpdated(log types.Log) (*AaveOracleFallbackOracleUpdated, error) {
	event := new(AaveOracleFallbackOracleUpdated)
	if err := _AaveOracle.contract.UnpackLog(event, "FallbackOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
