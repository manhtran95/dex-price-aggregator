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

// CurvePoolMetaData contains all meta data concerning the CurvePool contract.
var CurvePoolMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"get_dy\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"int128\",\"name\":\"i\"},{\"type\":\"int128\",\"name\":\"j\"},{\"type\":\"uint256\",\"name\":\"dx\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CurvePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvePoolMetaData.ABI instead.
var CurvePoolABI = CurvePoolMetaData.ABI

// CurvePool is an auto generated Go binding around an Ethereum contract.
type CurvePool struct {
	CurvePoolCaller     // Read-only binding to the contract
	CurvePoolTransactor // Write-only binding to the contract
	CurvePoolFilterer   // Log filterer for contract events
}

// CurvePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurvePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvePoolSession struct {
	Contract     *CurvePool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurvePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvePoolCallerSession struct {
	Contract *CurvePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CurvePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvePoolTransactorSession struct {
	Contract     *CurvePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CurvePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurvePoolRaw struct {
	Contract *CurvePool // Generic contract binding to access the raw methods on
}

// CurvePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvePoolCallerRaw struct {
	Contract *CurvePoolCaller // Generic read-only contract binding to access the raw methods on
}

// CurvePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvePoolTransactorRaw struct {
	Contract *CurvePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvePool creates a new instance of CurvePool, bound to a specific deployed contract.
func NewCurvePool(address common.Address, backend bind.ContractBackend) (*CurvePool, error) {
	contract, err := bindCurvePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurvePool{CurvePoolCaller: CurvePoolCaller{contract: contract}, CurvePoolTransactor: CurvePoolTransactor{contract: contract}, CurvePoolFilterer: CurvePoolFilterer{contract: contract}}, nil
}

// NewCurvePoolCaller creates a new read-only instance of CurvePool, bound to a specific deployed contract.
func NewCurvePoolCaller(address common.Address, caller bind.ContractCaller) (*CurvePoolCaller, error) {
	contract, err := bindCurvePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolCaller{contract: contract}, nil
}

// NewCurvePoolTransactor creates a new write-only instance of CurvePool, bound to a specific deployed contract.
func NewCurvePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*CurvePoolTransactor, error) {
	contract, err := bindCurvePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolTransactor{contract: contract}, nil
}

// NewCurvePoolFilterer creates a new log filterer instance of CurvePool, bound to a specific deployed contract.
func NewCurvePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*CurvePoolFilterer, error) {
	contract, err := bindCurvePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvePoolFilterer{contract: contract}, nil
}

// bindCurvePool binds a generic wrapper to an already deployed contract.
func bindCurvePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurvePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePool *CurvePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePool.Contract.CurvePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePool *CurvePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.Contract.CurvePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePool *CurvePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePool.Contract.CurvePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePool *CurvePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePool *CurvePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePool *CurvePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePool.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePool *CurvePoolCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePool *CurvePoolSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.Balances(&_CurvePool.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.Balances(&_CurvePool.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePool *CurvePoolCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePool *CurvePoolSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurvePool.Contract.Coins(&_CurvePool.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_CurvePool *CurvePoolCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _CurvePool.Contract.Coins(&_CurvePool.CallOpts, arg0)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePool *CurvePoolCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurvePool.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePool *CurvePoolSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.GetDy(&_CurvePool.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurvePool *CurvePoolCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurvePool.Contract.GetDy(&_CurvePool.CallOpts, i, j, dx)
}
