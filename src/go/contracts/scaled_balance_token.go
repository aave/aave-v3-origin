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

// IScaledBalanceTokenMetaData contains all meta data concerning the IScaledBalanceToken contract.
var IScaledBalanceTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getPreviousIndex\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getScaledUserBalanceAndSupply\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"scaledBalanceOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"scaledTotalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Burn\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"balanceIncrease\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Mint\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"balanceIncrease\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// IScaledBalanceTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IScaledBalanceTokenMetaData.ABI instead.
var IScaledBalanceTokenABI = IScaledBalanceTokenMetaData.ABI

// IScaledBalanceToken is an auto generated Go binding around an Ethereum contract.
type IScaledBalanceToken struct {
	IScaledBalanceTokenCaller     // Read-only binding to the contract
	IScaledBalanceTokenTransactor // Write-only binding to the contract
	IScaledBalanceTokenFilterer   // Log filterer for contract events
}

// IScaledBalanceTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IScaledBalanceTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScaledBalanceTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IScaledBalanceTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScaledBalanceTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IScaledBalanceTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScaledBalanceTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IScaledBalanceTokenSession struct {
	Contract     *IScaledBalanceToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IScaledBalanceTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IScaledBalanceTokenCallerSession struct {
	Contract *IScaledBalanceTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IScaledBalanceTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IScaledBalanceTokenTransactorSession struct {
	Contract     *IScaledBalanceTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IScaledBalanceTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IScaledBalanceTokenRaw struct {
	Contract *IScaledBalanceToken // Generic contract binding to access the raw methods on
}

// IScaledBalanceTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IScaledBalanceTokenCallerRaw struct {
	Contract *IScaledBalanceTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IScaledBalanceTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IScaledBalanceTokenTransactorRaw struct {
	Contract *IScaledBalanceTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIScaledBalanceToken creates a new instance of IScaledBalanceToken, bound to a specific deployed contract.
func NewIScaledBalanceToken(address common.Address, backend bind.ContractBackend) (*IScaledBalanceToken, error) {
	contract, err := bindIScaledBalanceToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IScaledBalanceToken{IScaledBalanceTokenCaller: IScaledBalanceTokenCaller{contract: contract}, IScaledBalanceTokenTransactor: IScaledBalanceTokenTransactor{contract: contract}, IScaledBalanceTokenFilterer: IScaledBalanceTokenFilterer{contract: contract}}, nil
}

// NewIScaledBalanceTokenCaller creates a new read-only instance of IScaledBalanceToken, bound to a specific deployed contract.
func NewIScaledBalanceTokenCaller(address common.Address, caller bind.ContractCaller) (*IScaledBalanceTokenCaller, error) {
	contract, err := bindIScaledBalanceToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IScaledBalanceTokenCaller{contract: contract}, nil
}

// NewIScaledBalanceTokenTransactor creates a new write-only instance of IScaledBalanceToken, bound to a specific deployed contract.
func NewIScaledBalanceTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IScaledBalanceTokenTransactor, error) {
	contract, err := bindIScaledBalanceToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IScaledBalanceTokenTransactor{contract: contract}, nil
}

// NewIScaledBalanceTokenFilterer creates a new log filterer instance of IScaledBalanceToken, bound to a specific deployed contract.
func NewIScaledBalanceTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IScaledBalanceTokenFilterer, error) {
	contract, err := bindIScaledBalanceToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IScaledBalanceTokenFilterer{contract: contract}, nil
}

// bindIScaledBalanceToken binds a generic wrapper to an already deployed contract.
func bindIScaledBalanceToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IScaledBalanceTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScaledBalanceToken *IScaledBalanceTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScaledBalanceToken.Contract.IScaledBalanceTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScaledBalanceToken *IScaledBalanceTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScaledBalanceToken.Contract.IScaledBalanceTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScaledBalanceToken *IScaledBalanceTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScaledBalanceToken.Contract.IScaledBalanceTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScaledBalanceToken *IScaledBalanceTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScaledBalanceToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScaledBalanceToken *IScaledBalanceTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScaledBalanceToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScaledBalanceToken *IScaledBalanceTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScaledBalanceToken.Contract.contract.Transact(opts, method, params...)
}

// GetPreviousIndex is a free data retrieval call binding the contract method 0xe0753986.
//
// Solidity: function getPreviousIndex(address user) view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCaller) GetPreviousIndex(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IScaledBalanceToken.contract.Call(opts, &out, "getPreviousIndex", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreviousIndex is a free data retrieval call binding the contract method 0xe0753986.
//
// Solidity: function getPreviousIndex(address user) view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenSession) GetPreviousIndex(user common.Address) (*big.Int, error) {
	return _IScaledBalanceToken.Contract.GetPreviousIndex(&_IScaledBalanceToken.CallOpts, user)
}

// GetPreviousIndex is a free data retrieval call binding the contract method 0xe0753986.
//
// Solidity: function getPreviousIndex(address user) view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCallerSession) GetPreviousIndex(user common.Address) (*big.Int, error) {
	return _IScaledBalanceToken.Contract.GetPreviousIndex(&_IScaledBalanceToken.CallOpts, user)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCaller) GetScaledUserBalanceAndSupply(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IScaledBalanceToken.contract.Call(opts, &out, "getScaledUserBalanceAndSupply", user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _IScaledBalanceToken.Contract.GetScaledUserBalanceAndSupply(&_IScaledBalanceToken.CallOpts, user)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCallerSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _IScaledBalanceToken.Contract.GetScaledUserBalanceAndSupply(&_IScaledBalanceToken.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCaller) ScaledBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IScaledBalanceToken.contract.Call(opts, &out, "scaledBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _IScaledBalanceToken.Contract.ScaledBalanceOf(&_IScaledBalanceToken.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCallerSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _IScaledBalanceToken.Contract.ScaledBalanceOf(&_IScaledBalanceToken.CallOpts, user)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCaller) ScaledTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IScaledBalanceToken.contract.Call(opts, &out, "scaledTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenSession) ScaledTotalSupply() (*big.Int, error) {
	return _IScaledBalanceToken.Contract.ScaledTotalSupply(&_IScaledBalanceToken.CallOpts)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCallerSession) ScaledTotalSupply() (*big.Int, error) {
	return _IScaledBalanceToken.Contract.ScaledTotalSupply(&_IScaledBalanceToken.CallOpts)
}

// IScaledBalanceTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the IScaledBalanceToken contract.
type IScaledBalanceTokenBurnIterator struct {
	Event *IScaledBalanceTokenBurn // Event containing the contract specifics and raw log

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
func (it *IScaledBalanceTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScaledBalanceTokenBurn)
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
		it.Event = new(IScaledBalanceTokenBurn)
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
func (it *IScaledBalanceTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScaledBalanceTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScaledBalanceTokenBurn represents a Burn event raised by the IScaledBalanceToken contract.
type IScaledBalanceTokenBurn struct {
	From            common.Address
	Target          common.Address
	Value           *big.Int
	BalanceIncrease *big.Int
	Index           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x4cf25bc1d991c17529c25213d3cc0cda295eeaad5f13f361969b12ea48015f90.
//
// Solidity: event Burn(address indexed from, address indexed target, uint256 value, uint256 balanceIncrease, uint256 index)
func (_IScaledBalanceToken *IScaledBalanceTokenFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address, target []common.Address) (*IScaledBalanceTokenBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IScaledBalanceToken.contract.FilterLogs(opts, "Burn", fromRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &IScaledBalanceTokenBurnIterator{contract: _IScaledBalanceToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x4cf25bc1d991c17529c25213d3cc0cda295eeaad5f13f361969b12ea48015f90.
//
// Solidity: event Burn(address indexed from, address indexed target, uint256 value, uint256 balanceIncrease, uint256 index)
func (_IScaledBalanceToken *IScaledBalanceTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IScaledBalanceTokenBurn, from []common.Address, target []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IScaledBalanceToken.contract.WatchLogs(opts, "Burn", fromRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScaledBalanceTokenBurn)
				if err := _IScaledBalanceToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x4cf25bc1d991c17529c25213d3cc0cda295eeaad5f13f361969b12ea48015f90.
//
// Solidity: event Burn(address indexed from, address indexed target, uint256 value, uint256 balanceIncrease, uint256 index)
func (_IScaledBalanceToken *IScaledBalanceTokenFilterer) ParseBurn(log types.Log) (*IScaledBalanceTokenBurn, error) {
	event := new(IScaledBalanceTokenBurn)
	if err := _IScaledBalanceToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScaledBalanceTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IScaledBalanceToken contract.
type IScaledBalanceTokenMintIterator struct {
	Event *IScaledBalanceTokenMint // Event containing the contract specifics and raw log

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
func (it *IScaledBalanceTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScaledBalanceTokenMint)
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
		it.Event = new(IScaledBalanceTokenMint)
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
func (it *IScaledBalanceTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScaledBalanceTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScaledBalanceTokenMint represents a Mint event raised by the IScaledBalanceToken contract.
type IScaledBalanceTokenMint struct {
	Caller          common.Address
	OnBehalfOf      common.Address
	Value           *big.Int
	BalanceIncrease *big.Int
	Index           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196.
//
// Solidity: event Mint(address indexed caller, address indexed onBehalfOf, uint256 value, uint256 balanceIncrease, uint256 index)
func (_IScaledBalanceToken *IScaledBalanceTokenFilterer) FilterMint(opts *bind.FilterOpts, caller []common.Address, onBehalfOf []common.Address) (*IScaledBalanceTokenMintIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _IScaledBalanceToken.contract.FilterLogs(opts, "Mint", callerRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &IScaledBalanceTokenMintIterator{contract: _IScaledBalanceToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196.
//
// Solidity: event Mint(address indexed caller, address indexed onBehalfOf, uint256 value, uint256 balanceIncrease, uint256 index)
func (_IScaledBalanceToken *IScaledBalanceTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IScaledBalanceTokenMint, caller []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _IScaledBalanceToken.contract.WatchLogs(opts, "Mint", callerRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScaledBalanceTokenMint)
				if err := _IScaledBalanceToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196.
//
// Solidity: event Mint(address indexed caller, address indexed onBehalfOf, uint256 value, uint256 balanceIncrease, uint256 index)
func (_IScaledBalanceToken *IScaledBalanceTokenFilterer) ParseMint(log types.Log) (*IScaledBalanceTokenMint, error) {
	event := new(IScaledBalanceTokenMint)
	if err := _IScaledBalanceToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
