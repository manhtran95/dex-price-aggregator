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

// UniswapV3QuoterMetaData contains all meta data concerning the UniswapV3Quoter contract.
var UniswapV3QuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniswapV3QuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3QuoterMetaData.ABI instead.
var UniswapV3QuoterABI = UniswapV3QuoterMetaData.ABI

// UniswapV3Quoter is an auto generated Go binding around an Ethereum contract.
type UniswapV3Quoter struct {
	UniswapV3QuoterCaller     // Read-only binding to the contract
	UniswapV3QuoterTransactor // Write-only binding to the contract
	UniswapV3QuoterFilterer   // Log filterer for contract events
}

// UniswapV3QuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3QuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3QuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3QuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3QuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3QuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3QuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3QuoterSession struct {
	Contract     *UniswapV3Quoter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV3QuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3QuoterCallerSession struct {
	Contract *UniswapV3QuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UniswapV3QuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3QuoterTransactorSession struct {
	Contract     *UniswapV3QuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UniswapV3QuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3QuoterRaw struct {
	Contract *UniswapV3Quoter // Generic contract binding to access the raw methods on
}

// UniswapV3QuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3QuoterCallerRaw struct {
	Contract *UniswapV3QuoterCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3QuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3QuoterTransactorRaw struct {
	Contract *UniswapV3QuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3Quoter creates a new instance of UniswapV3Quoter, bound to a specific deployed contract.
func NewUniswapV3Quoter(address common.Address, backend bind.ContractBackend) (*UniswapV3Quoter, error) {
	contract, err := bindUniswapV3Quoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Quoter{UniswapV3QuoterCaller: UniswapV3QuoterCaller{contract: contract}, UniswapV3QuoterTransactor: UniswapV3QuoterTransactor{contract: contract}, UniswapV3QuoterFilterer: UniswapV3QuoterFilterer{contract: contract}}, nil
}

// NewUniswapV3QuoterCaller creates a new read-only instance of UniswapV3Quoter, bound to a specific deployed contract.
func NewUniswapV3QuoterCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3QuoterCaller, error) {
	contract, err := bindUniswapV3Quoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3QuoterCaller{contract: contract}, nil
}

// NewUniswapV3QuoterTransactor creates a new write-only instance of UniswapV3Quoter, bound to a specific deployed contract.
func NewUniswapV3QuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3QuoterTransactor, error) {
	contract, err := bindUniswapV3Quoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3QuoterTransactor{contract: contract}, nil
}

// NewUniswapV3QuoterFilterer creates a new log filterer instance of UniswapV3Quoter, bound to a specific deployed contract.
func NewUniswapV3QuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3QuoterFilterer, error) {
	contract, err := bindUniswapV3Quoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3QuoterFilterer{contract: contract}, nil
}

// bindUniswapV3Quoter binds a generic wrapper to an already deployed contract.
func bindUniswapV3Quoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3QuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Quoter *UniswapV3QuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Quoter.Contract.UniswapV3QuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Quoter *UniswapV3QuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Quoter.Contract.UniswapV3QuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Quoter *UniswapV3QuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Quoter.Contract.UniswapV3QuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Quoter *UniswapV3QuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Quoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Quoter *UniswapV3QuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Quoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Quoter *UniswapV3QuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Quoter.Contract.contract.Transact(opts, method, params...)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) returns(uint256 amountOut)
func (_UniswapV3Quoter *UniswapV3QuoterTransactor) QuoteExactInputSingle(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _UniswapV3Quoter.contract.Transact(opts, "quoteExactInputSingle", tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) returns(uint256 amountOut)
func (_UniswapV3Quoter *UniswapV3QuoterSession) QuoteExactInputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _UniswapV3Quoter.Contract.QuoteExactInputSingle(&_UniswapV3Quoter.TransactOpts, tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) returns(uint256 amountOut)
func (_UniswapV3Quoter *UniswapV3QuoterTransactorSession) QuoteExactInputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _UniswapV3Quoter.Contract.QuoteExactInputSingle(&_UniswapV3Quoter.TransactOpts, tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)
}
