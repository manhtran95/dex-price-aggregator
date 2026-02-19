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

// UniswapV2FactoryMetaData contains all meta data concerning the UniswapV2Factory contract.
var UniswapV2FactoryMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapV2FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV2FactoryMetaData.ABI instead.
var UniswapV2FactoryABI = UniswapV2FactoryMetaData.ABI

// UniswapV2Factory is an auto generated Go binding around an Ethereum contract.
type UniswapV2Factory struct {
	UniswapV2FactoryCaller     // Read-only binding to the contract
	UniswapV2FactoryTransactor // Write-only binding to the contract
	UniswapV2FactoryFilterer   // Log filterer for contract events
}

// UniswapV2FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV2FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV2FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV2FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV2FactorySession struct {
	Contract     *UniswapV2Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV2FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV2FactoryCallerSession struct {
	Contract *UniswapV2FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UniswapV2FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV2FactoryTransactorSession struct {
	Contract     *UniswapV2FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UniswapV2FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV2FactoryRaw struct {
	Contract *UniswapV2Factory // Generic contract binding to access the raw methods on
}

// UniswapV2FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV2FactoryCallerRaw struct {
	Contract *UniswapV2FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV2FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV2FactoryTransactorRaw struct {
	Contract *UniswapV2FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV2Factory creates a new instance of UniswapV2Factory, bound to a specific deployed contract.
func NewUniswapV2Factory(address common.Address, backend bind.ContractBackend) (*UniswapV2Factory, error) {
	contract, err := bindUniswapV2Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Factory{UniswapV2FactoryCaller: UniswapV2FactoryCaller{contract: contract}, UniswapV2FactoryTransactor: UniswapV2FactoryTransactor{contract: contract}, UniswapV2FactoryFilterer: UniswapV2FactoryFilterer{contract: contract}}, nil
}

// NewUniswapV2FactoryCaller creates a new read-only instance of UniswapV2Factory, bound to a specific deployed contract.
func NewUniswapV2FactoryCaller(address common.Address, caller bind.ContractCaller) (*UniswapV2FactoryCaller, error) {
	contract, err := bindUniswapV2Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2FactoryCaller{contract: contract}, nil
}

// NewUniswapV2FactoryTransactor creates a new write-only instance of UniswapV2Factory, bound to a specific deployed contract.
func NewUniswapV2FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV2FactoryTransactor, error) {
	contract, err := bindUniswapV2Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2FactoryTransactor{contract: contract}, nil
}

// NewUniswapV2FactoryFilterer creates a new log filterer instance of UniswapV2Factory, bound to a specific deployed contract.
func NewUniswapV2FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV2FactoryFilterer, error) {
	contract, err := bindUniswapV2Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV2FactoryFilterer{contract: contract}, nil
}

// bindUniswapV2Factory binds a generic wrapper to an already deployed contract.
func bindUniswapV2Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV2FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Factory *UniswapV2FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Factory.Contract.UniswapV2FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Factory *UniswapV2FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Factory.Contract.UniswapV2FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Factory *UniswapV2FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Factory.Contract.UniswapV2FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Factory *UniswapV2FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Factory *UniswapV2FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Factory *UniswapV2FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Factory.Contract.contract.Transact(opts, method, params...)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_UniswapV2Factory *UniswapV2FactoryCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2Factory.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_UniswapV2Factory *UniswapV2FactorySession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _UniswapV2Factory.Contract.GetPair(&_UniswapV2Factory.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_UniswapV2Factory *UniswapV2FactoryCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _UniswapV2Factory.Contract.GetPair(&_UniswapV2Factory.CallOpts, arg0, arg1)
}
