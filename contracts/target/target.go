// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package target

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
)

// TargetMetaData contains all meta data concerning the Target contract.
var TargetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610277806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806340c10f191461004657806370a082311461005b578063983b2d5614610096575b600080fd5b6100596100543660046101cf565b6100c8565b005b6100846100693660046101f9565b6001600160a01b031660009081526001602052604090205490565b60405190815260200160405180910390f35b6100596100a43660046101f9565b6001600160a01b03166000908152602081905260409020805460ff19166001179055565b3360009081526020819052604090205460ff166101445760405162461bcd60e51b815260206004820152603060248201527f4d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766560448201526f20746865204d696e74657220726f6c6560801b606482015260840160405180910390fd5b6001600160a01b0382166000908152600160205260408120805483929061016c90849061021b565b90915550506040518181526001600160a01b038316907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859060200160405180910390a25050565b80356001600160a01b03811681146101ca57600080fd5b919050565b600080604083850312156101e257600080fd5b6101eb836101b3565b946020939093013593505050565b60006020828403121561020b57600080fd5b610214826101b3565b9392505050565b6000821982111561023c57634e487b7160e01b600052601160045260246000fd5b50019056fea264697066735822122010dcb48ba30125f01deb1d1a7723d79343181ce5bd9551deb2e8dae7dc61b13364736f6c634300080d0033",
}

// TargetABI is the input ABI used to generate the binding from.
// Deprecated: Use TargetMetaData.ABI instead.
var TargetABI = TargetMetaData.ABI

// TargetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TargetMetaData.Bin instead.
var TargetBin = TargetMetaData.Bin

// DeployTarget deploys a new Ethereum contract, binding an instance of Target to it.
func DeployTarget(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Target, error) {
	parsed, err := TargetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TargetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Target{TargetCaller: TargetCaller{contract: contract}, TargetTransactor: TargetTransactor{contract: contract}, TargetFilterer: TargetFilterer{contract: contract}}, nil
}

// Target is an auto generated Go binding around an Ethereum contract.
type Target struct {
	TargetCaller     // Read-only binding to the contract
	TargetTransactor // Write-only binding to the contract
	TargetFilterer   // Log filterer for contract events
}

// TargetCaller is an auto generated read-only Go binding around an Ethereum contract.
type TargetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TargetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TargetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TargetSession struct {
	Contract     *Target           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TargetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TargetCallerSession struct {
	Contract *TargetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TargetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TargetTransactorSession struct {
	Contract     *TargetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TargetRaw is an auto generated low-level Go binding around an Ethereum contract.
type TargetRaw struct {
	Contract *Target // Generic contract binding to access the raw methods on
}

// TargetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TargetCallerRaw struct {
	Contract *TargetCaller // Generic read-only contract binding to access the raw methods on
}

// TargetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TargetTransactorRaw struct {
	Contract *TargetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTarget creates a new instance of Target, bound to a specific deployed contract.
func NewTarget(address common.Address, backend bind.ContractBackend) (*Target, error) {
	contract, err := bindTarget(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Target{TargetCaller: TargetCaller{contract: contract}, TargetTransactor: TargetTransactor{contract: contract}, TargetFilterer: TargetFilterer{contract: contract}}, nil
}

// NewTargetCaller creates a new read-only instance of Target, bound to a specific deployed contract.
func NewTargetCaller(address common.Address, caller bind.ContractCaller) (*TargetCaller, error) {
	contract, err := bindTarget(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TargetCaller{contract: contract}, nil
}

// NewTargetTransactor creates a new write-only instance of Target, bound to a specific deployed contract.
func NewTargetTransactor(address common.Address, transactor bind.ContractTransactor) (*TargetTransactor, error) {
	contract, err := bindTarget(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TargetTransactor{contract: contract}, nil
}

// NewTargetFilterer creates a new log filterer instance of Target, bound to a specific deployed contract.
func NewTargetFilterer(address common.Address, filterer bind.ContractFilterer) (*TargetFilterer, error) {
	contract, err := bindTarget(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TargetFilterer{contract: contract}, nil
}

// bindTarget binds a generic wrapper to an already deployed contract.
func bindTarget(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TargetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Target *TargetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Target.Contract.TargetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Target *TargetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Target.Contract.TargetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Target *TargetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Target.Contract.TargetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Target *TargetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Target.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Target *TargetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Target.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Target *TargetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Target.Contract.contract.Transact(opts, method, params...)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter) returns()
func (_Target *TargetTransactor) AddMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _Target.contract.Transact(opts, "addMinter", minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter) returns()
func (_Target *TargetSession) AddMinter(minter common.Address) (*types.Transaction, error) {
	return _Target.Contract.AddMinter(&_Target.TransactOpts, minter)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address minter) returns()
func (_Target *TargetTransactorSession) AddMinter(minter common.Address) (*types.Transaction, error) {
	return _Target.Contract.AddMinter(&_Target.TransactOpts, minter)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address destinationAddress) returns(uint256)
func (_Target *TargetTransactor) BalanceOf(opts *bind.TransactOpts, destinationAddress common.Address) (*types.Transaction, error) {
	return _Target.contract.Transact(opts, "balanceOf", destinationAddress)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address destinationAddress) returns(uint256)
func (_Target *TargetSession) BalanceOf(destinationAddress common.Address) (*types.Transaction, error) {
	return _Target.Contract.BalanceOf(&_Target.TransactOpts, destinationAddress)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address destinationAddress) returns(uint256)
func (_Target *TargetTransactorSession) BalanceOf(destinationAddress common.Address) (*types.Transaction, error) {
	return _Target.Contract.BalanceOf(&_Target.TransactOpts, destinationAddress)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address destinationAddress, uint256 amount) returns()
func (_Target *TargetTransactor) Mint(opts *bind.TransactOpts, destinationAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Target.contract.Transact(opts, "mint", destinationAddress, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address destinationAddress, uint256 amount) returns()
func (_Target *TargetSession) Mint(destinationAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Target.Contract.Mint(&_Target.TransactOpts, destinationAddress, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address destinationAddress, uint256 amount) returns()
func (_Target *TargetTransactorSession) Mint(destinationAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Target.Contract.Mint(&_Target.TransactOpts, destinationAddress, amount)
}

// TargetMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Target contract.
type TargetMintIterator struct {
	Event *TargetMint // Event containing the contract specifics and raw log

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
func (it *TargetMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TargetMint)
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
		it.Event = new(TargetMint)
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
func (it *TargetMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TargetMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TargetMint represents a Mint event raised by the Target contract.
type TargetMint struct {
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed destination, uint256 amount)
func (_Target *TargetFilterer) FilterMint(opts *bind.FilterOpts, destination []common.Address) (*TargetMintIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _Target.contract.FilterLogs(opts, "Mint", destinationRule)
	if err != nil {
		return nil, err
	}
	return &TargetMintIterator{contract: _Target.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed destination, uint256 amount)
func (_Target *TargetFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *TargetMint, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _Target.contract.WatchLogs(opts, "Mint", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TargetMint)
				if err := _Target.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed destination, uint256 amount)
func (_Target *TargetFilterer) ParseMint(log types.Log) (*TargetMint, error) {
	event := new(TargetMint)
	if err := _Target.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
