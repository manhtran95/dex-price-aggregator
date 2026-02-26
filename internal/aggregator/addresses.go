package aggregator

import "github.com/ethereum/go-ethereum/common"

// Static pool registry (could be from DB, but hardcoded for simplicity)
var KnownPools = []PoolInfo{

	// ============ Uniswap V3 ============
	// WETH/USDC
	// WETH/USDC 0.05% - 0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640
	// {Address: common.HexToAddress("0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640"), Token0: USDC_ADDRESS, Token1: WETH_ADDRESS, DEX: "UniswapV3", Fee: 500},
	// WETH/USDC 0.3% - 0x8ad599c3A0ff1De082011EFDDc58f1908eb6e6D8
	{Address: common.HexToAddress("0x8ad599c3A0ff1De082011EFDDc58f1908eb6e6D8"), Token0: USDC_ADDRESS, Token1: WETH_ADDRESS, DEX: "UniswapV3", Fee: 3000},
	// WETH/USDC 1% - 0x7BeA39867e4169DBe237d55C8242a8f2fcDcc387 - 700k TVL
	// {Address: common.HexToAddress("0x7BeA39867e4169DBe237d55C8242a8f2fcDcc387"), Token0: USDC_ADDRESS, Token1: WETH_ADDRESS, DEX: "UniswapV3", Fee: 10000},

	// WETH/DAI
	// WETH/DAI 0.05% - 0x60594a405d53811d3BC4766596EFD80fd545A270 - 1.1M
	// {Address: common.HexToAddress("0x60594a405d53811d3BC4766596EFD80fd545A270"), Token0: DAI_ADDRESS, Token1: WETH_ADDRESS, DEX: "UniswapV3", Fee: 500},
	// WETH/DAI 0.3% - 0xC2e9F25Be6257c210d7Adf0D4Cd6E3E881ba25f8 - 2.9M
	{Address: common.HexToAddress("0xC2e9F25Be6257c210d7Adf0D4Cd6E3E881ba25f8"), Token0: DAI_ADDRESS, Token1: WETH_ADDRESS, DEX: "UniswapV3", Fee: 3000},

	// WETH/USDT
	// WETH/USDT 0.05% - 0x11b815efB8f581194ae79006d24E0d814B7697F6
	// {Address: common.HexToAddress("0x11b815efB8f581194ae79006d24E0d814B7697F6"), Token0: WETH_ADDRESS, Token1: USDT_ADDRESS, DEX: "UniswapV3", Fee: 500},
	// WETH/USDT 0.3% - 0x4e68Ccd3E89f51C3074ca5072bbAC773960dFa36
	{Address: common.HexToAddress("0x4e68Ccd3E89f51C3074ca5072bbAC773960dFa36"), Token0: WETH_ADDRESS, Token1: USDT_ADDRESS, DEX: "UniswapV3", Fee: 3000},

	// WETH/WBTC
	// WETH/WBTC 0.05% - 0x4585FE77225b41b697C938B018E2Ac67Ac5a20c0
	// {Address: common.HexToAddress("0x4585FE77225b41b697C938B018E2Ac67Ac5a20c0"), Token0: WBTC_ADDRESS, Token1: WETH_ADDRESS, DEX: "UniswapV3", Fee: 500},
	// WETH/WBTC 0.3% - 0xCBCdF9626bC03E24f779434178A73a0B4bad62eD
	{Address: common.HexToAddress("0xCBCdF9626bC03E24f779434178A73a0B4bad62eD"), Token0: WBTC_ADDRESS, Token1: WETH_ADDRESS, DEX: "UniswapV3", Fee: 3000},

	// DAI/USDC
	// DAI/USDC 0.01% - 0x5777d92f208679DB4b9778590Fa3CAB3aC9e2168
	// {Address: common.HexToAddress("0x5777d92f208679DB4b9778590Fa3CAB3aC9e2168"), Token0: DAI_ADDRESS, Token1: USDC_ADDRESS, DEX: "UniswapV3", Fee: 100},
	// DAI/USDC 0.05% - 0x6c6Bc977E13Df9b0de53b251522280BB72383700
	{Address: common.HexToAddress("0x6c6Bc977E13Df9b0de53b251522280BB72383700"), Token0: DAI_ADDRESS, Token1: USDC_ADDRESS, DEX: "UniswapV3", Fee: 500},

    // ============ Curve ============
    // 3pool: DAI + USDC + USDT
    {
        Address: common.HexToAddress("0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7"),
        Tokens:  []common.Address{DAI_ADDRESS, USDC_ADDRESS, USDT_ADDRESS},
        DEX:     "Curve",
        Fee:     400, // 0.04%
    },
    // FRAX/USDC pool
    {
        Address: common.HexToAddress("0xDcEF968d416a41Cdac0ED8702fAC8128A64241A2"),
        Token0:  FRAX_ADDRESS,
        Token1:  USDC_ADDRESS,
        DEX:     "Curve",
        Fee:     400,
    },	

	/*
		// ============ Uniswap V2 ============
		// WETH/USDC - 0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc
		{Address: common.HexToAddress("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc"), Token0: USDC_ADDRESS, Token1: WETH_ADDRESS, DEX: "UniswapV2", Fee: 3000},
		// WETH/DAI - 0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11
		{Address: common.HexToAddress("0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11"), Token0: DAI_ADDRESS, Token1: WETH_ADDRESS, DEX: "UniswapV2", Fee: 3000},
		// WETH/USDT - 0x0d4a11d5EEaaC28EC3F61d100daF4d40471f1852
		{Address: common.HexToAddress("0x0d4a11d5EEaaC28EC3F61d100daF4d40471f1852"), Token0: WETH_ADDRESS, Token1: USDT_ADDRESS, DEX: "UniswapV2", Fee: 3000},
		// WETH/WBTC - 0xBb2b8038a1640196FbE3e38816F3e67Cba72D940
		{Address: common.HexToAddress("0xBb2b8038a1640196FbE3e38816F3e67Cba72D940"), Token0: WBTC_ADDRESS, Token1: WETH_ADDRESS, DEX: "UniswapV2", Fee: 3000},
		// DAI/USDC - 0xAE461cA67B15dc8dc81CE7615e0320dA1A9aB8D5
		{Address: common.HexToAddress("0xAE461cA67B15dc8dc81CE7615e0320dA1A9aB8D5"), Token0: DAI_ADDRESS, Token1: USDC_ADDRESS, DEX: "UniswapV2", Fee: 3000},
	*/

	// ============ SushiSwap ============
	/*
		// WETH/USDC - 0x397FF1542f962076d0BFE58eA045FfA2d347ACa0
		{Address: common.HexToAddress("0x397FF1542f962076d0BFE58eA045FfA2d347ACa0"), Token0: USDC_ADDRESS, Token1: WETH_ADDRESS, DEX: "SushiSwap", Fee: 3000},
		// WETH/DAI - 0xC3D03e4F041Fd4cD388c549Ee2A29a9E5075882f
		{Address: common.HexToAddress("0xC3D03e4F041Fd4cD388c549Ee2A29a9E5075882f"), Token0: DAI_ADDRESS, Token1: WETH_ADDRESS, DEX: "SushiSwap", Fee: 3000},
		// WETH/USDT - 0x06da0fd433C1A5d7a4faa01111c044910A184553
		{Address: common.HexToAddress("0x06da0fd433C1A5d7a4faa01111c044910A184553"), Token0: WETH_ADDRESS, Token1: USDT_ADDRESS, DEX: "SushiSwap", Fee: 3000},
		// WETH/WBTC - 0xCEfF51756c56CeFFCA006cD410B03FFC46dd3a58
		{Address: common.HexToAddress("0xCEfF51756c56CeFFCA006cD410B03FFC46dd3a58"), Token0: WBTC_ADDRESS, Token1: WETH_ADDRESS, DEX: "SushiSwap", Fee: 3000},
	*/
}

// Common intermediate tokens
var (
	WETH_ADDRESS = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	DAI_ADDRESS  = common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	USDC_ADDRESS = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	USDT_ADDRESS = common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
    FRAX_ADDRESS = common.HexToAddress("0x853d955aCEf822Db058eb8505911ED77F175b99e")
	WBTC_ADDRESS = common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599")
)

var TokenSymbols = map[common.Address]string{
	USDC_ADDRESS: "USDC",
	WETH_ADDRESS: "WETH",
	DAI_ADDRESS:  "DAI",
	USDT_ADDRESS: "USDT",
	WBTC_ADDRESS: "WBTC",
}
