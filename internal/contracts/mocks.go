package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// MockUniswapV2Factory mocks the factory contract
type MockUniswapV2Factory struct {
	GetPairFunc func(opts *bind.CallOpts, tokenA, tokenB common.Address) (common.Address, error)
}

func (m *MockUniswapV2Factory) GetPair(opts *bind.CallOpts, tokenA, tokenB common.Address) (common.Address, error) {
	if m.GetPairFunc != nil {
		return m.GetPairFunc(opts, tokenA, tokenB)
	}
	// Return a mock pair address
	return common.HexToAddress("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc"), nil
}

// MockUniswapV2Pair mocks the pair contract
type MockUniswapV2Pair struct {
	GetReservesFunc func(opts *bind.CallOpts) (struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	}, error)
	Token0Func func(opts *bind.CallOpts) (common.Address, error)
	Token1Func func(opts *bind.CallOpts) (common.Address, error)
}

func (m *MockUniswapV2Pair) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	if m.GetReservesFunc != nil {
		return m.GetReservesFunc(opts)
	}

	// Return mock reserves: 50,000 WETH and 100,000,000 USDC
	return struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	}{
		Reserve0:           new(big.Int).Mul(big.NewInt(50000), big.NewInt(1e18)),
		Reserve1:           new(big.Int).Mul(big.NewInt(100000000), big.NewInt(1e6)),
		BlockTimestampLast: 0,
	}, nil
}

func (m *MockUniswapV2Pair) Token0(opts *bind.CallOpts) (common.Address, error) {
	if m.Token0Func != nil {
		return m.Token0Func(opts)
	}
	return common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), nil // WETH
}

func (m *MockUniswapV2Pair) Token1(opts *bind.CallOpts) (common.Address, error) {
	if m.Token1Func != nil {
		return m.Token1Func(opts)
	}
	return common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), nil // USDC
}

// MockERC20 mocks the ERC20 contract
type MockERC20 struct {
	SymbolFunc    func(opts *bind.CallOpts) (string, error)
	NameFunc      func(opts *bind.CallOpts) (string, error)
	DecimalsFunc  func(opts *bind.CallOpts) (uint8, error)
	BalanceOfFunc func(opts *bind.CallOpts, account common.Address) (*big.Int, error)
}

func (m *MockERC20) Symbol(opts *bind.CallOpts) (string, error) {
	if m.SymbolFunc != nil {
		return m.SymbolFunc(opts)
	}
	return "MOCK", nil
}

func (m *MockERC20) Name(opts *bind.CallOpts) (string, error) {
	if m.NameFunc != nil {
		return m.NameFunc(opts)
	}
	return "Mock Token", nil
}

func (m *MockERC20) Decimals(opts *bind.CallOpts) (uint8, error) {
	if m.DecimalsFunc != nil {
		return m.DecimalsFunc(opts)
	}
	return 18, nil
}

func (m *MockERC20) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	if m.BalanceOfFunc != nil {
		return m.BalanceOfFunc(opts, account)
	}
	return big.NewInt(1000000000000000000), nil // 1 token
}
