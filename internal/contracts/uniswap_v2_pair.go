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

// UniswapV2PairMetaData contains all meta data concerning the UniswapV2Pair contract.
var UniswapV2PairMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniswapV2PairABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV2PairMetaData.ABI instead.
var UniswapV2PairABI = UniswapV2PairMetaData.ABI

// UniswapV2Pair is an auto generated Go binding around an Ethereum contract.
type UniswapV2Pair struct {
	UniswapV2PairCaller     // Read-only binding to the contract
	UniswapV2PairTransactor // Write-only binding to the contract
	UniswapV2PairFilterer   // Log filterer for contract events
}

// UniswapV2PairCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV2PairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2PairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV2PairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2PairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV2PairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2PairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV2PairSession struct {
	Contract     *UniswapV2Pair    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV2PairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV2PairCallerSession struct {
	Contract *UniswapV2PairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UniswapV2PairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV2PairTransactorSession struct {
	Contract     *UniswapV2PairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UniswapV2PairRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV2PairRaw struct {
	Contract *UniswapV2Pair // Generic contract binding to access the raw methods on
}

// UniswapV2PairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV2PairCallerRaw struct {
	Contract *UniswapV2PairCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV2PairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV2PairTransactorRaw struct {
	Contract *UniswapV2PairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV2Pair creates a new instance of UniswapV2Pair, bound to a specific deployed contract.
func NewUniswapV2Pair(address common.Address, backend bind.ContractBackend) (*UniswapV2Pair, error) {
	contract, err := bindUniswapV2Pair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Pair{UniswapV2PairCaller: UniswapV2PairCaller{contract: contract}, UniswapV2PairTransactor: UniswapV2PairTransactor{contract: contract}, UniswapV2PairFilterer: UniswapV2PairFilterer{contract: contract}}, nil
}

// NewUniswapV2PairCaller creates a new read-only instance of UniswapV2Pair, bound to a specific deployed contract.
func NewUniswapV2PairCaller(address common.Address, caller bind.ContractCaller) (*UniswapV2PairCaller, error) {
	contract, err := bindUniswapV2Pair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2PairCaller{contract: contract}, nil
}

// NewUniswapV2PairTransactor creates a new write-only instance of UniswapV2Pair, bound to a specific deployed contract.
func NewUniswapV2PairTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV2PairTransactor, error) {
	contract, err := bindUniswapV2Pair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2PairTransactor{contract: contract}, nil
}

// NewUniswapV2PairFilterer creates a new log filterer instance of UniswapV2Pair, bound to a specific deployed contract.
func NewUniswapV2PairFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV2PairFilterer, error) {
	contract, err := bindUniswapV2Pair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV2PairFilterer{contract: contract}, nil
}

// bindUniswapV2Pair binds a generic wrapper to an already deployed contract.
func bindUniswapV2Pair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV2PairMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Pair *UniswapV2PairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Pair.Contract.UniswapV2PairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Pair *UniswapV2PairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Pair.Contract.UniswapV2PairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Pair *UniswapV2PairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Pair.Contract.UniswapV2PairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Pair *UniswapV2PairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Pair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Pair *UniswapV2PairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Pair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Pair *UniswapV2PairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Pair.Contract.contract.Transact(opts, method, params...)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_UniswapV2Pair *UniswapV2PairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _UniswapV2Pair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_UniswapV2Pair *UniswapV2PairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _UniswapV2Pair.Contract.GetReserves(&_UniswapV2Pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_UniswapV2Pair *UniswapV2PairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _UniswapV2Pair.Contract.GetReserves(&_UniswapV2Pair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_UniswapV2Pair *UniswapV2PairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV2Pair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_UniswapV2Pair *UniswapV2PairSession) Price0CumulativeLast() (*big.Int, error) {
	return _UniswapV2Pair.Contract.Price0CumulativeLast(&_UniswapV2Pair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_UniswapV2Pair *UniswapV2PairCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _UniswapV2Pair.Contract.Price0CumulativeLast(&_UniswapV2Pair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_UniswapV2Pair *UniswapV2PairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapV2Pair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_UniswapV2Pair *UniswapV2PairSession) Price1CumulativeLast() (*big.Int, error) {
	return _UniswapV2Pair.Contract.Price1CumulativeLast(&_UniswapV2Pair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_UniswapV2Pair *UniswapV2PairCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _UniswapV2Pair.Contract.Price1CumulativeLast(&_UniswapV2Pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_UniswapV2Pair *UniswapV2PairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2Pair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_UniswapV2Pair *UniswapV2PairSession) Token0() (common.Address, error) {
	return _UniswapV2Pair.Contract.Token0(&_UniswapV2Pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_UniswapV2Pair *UniswapV2PairCallerSession) Token0() (common.Address, error) {
	return _UniswapV2Pair.Contract.Token0(&_UniswapV2Pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_UniswapV2Pair *UniswapV2PairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2Pair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_UniswapV2Pair *UniswapV2PairSession) Token1() (common.Address, error) {
	return _UniswapV2Pair.Contract.Token1(&_UniswapV2Pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_UniswapV2Pair *UniswapV2PairCallerSession) Token1() (common.Address, error) {
	return _UniswapV2Pair.Contract.Token1(&_UniswapV2Pair.CallOpts)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_UniswapV2Pair *UniswapV2PairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _UniswapV2Pair.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_UniswapV2Pair *UniswapV2PairSession) Skim(to common.Address) (*types.Transaction, error) {
	return _UniswapV2Pair.Contract.Skim(&_UniswapV2Pair.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_UniswapV2Pair *UniswapV2PairTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _UniswapV2Pair.Contract.Skim(&_UniswapV2Pair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_UniswapV2Pair *UniswapV2PairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _UniswapV2Pair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_UniswapV2Pair *UniswapV2PairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _UniswapV2Pair.Contract.Swap(&_UniswapV2Pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_UniswapV2Pair *UniswapV2PairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _UniswapV2Pair.Contract.Swap(&_UniswapV2Pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_UniswapV2Pair *UniswapV2PairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Pair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_UniswapV2Pair *UniswapV2PairSession) Sync() (*types.Transaction, error) {
	return _UniswapV2Pair.Contract.Sync(&_UniswapV2Pair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_UniswapV2Pair *UniswapV2PairTransactorSession) Sync() (*types.Transaction, error) {
	return _UniswapV2Pair.Contract.Sync(&_UniswapV2Pair.TransactOpts)
}
