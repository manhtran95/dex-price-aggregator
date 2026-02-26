#!/bin/bash

# Generate Uniswap V2 Factory bindings
abigen --abi=contracts/abis/UniswapV2Factory.json \
       --pkg=contracts \
       --type=UniswapV2Factory \
       --out=internal/contracts/uniswap_v2_factory.go

# Generate Uniswap V2 Pair bindings
abigen --abi=contracts/abis/UniswapV2Pair.json \
       --pkg=contracts \
       --type=UniswapV2Pair \
       --out=internal/contracts/uniswap_v2_pair.go

# Generate ERC20 bindings
abigen --abi=contracts/abis/ERC20.json \
       --pkg=contracts \
       --type=ERC20 \
       --out=internal/contracts/erc20.go

# Generate Uniswap V3 Factory bindings
abigen --abi=contracts/abis/UniswapV3Factory.json \
       --pkg=contracts \
       --type=UniswapV3Factory \
       --out=internal/contracts/uniswap_v3_factory.go

# Generate Uniswap V3 Quoter bindings
abigen --abi=contracts/abis/UniswapV3Quoter.json \
       --pkg=contracts \
       --type=UniswapV3Quoter \
       --out=internal/contracts/uniswap_v3_quoter.go

# Generate Curve Pool bindings
abigen --abi=contracts/abis/CurvePool.json \
       --pkg=contracts \
       --type=CurvePool \
       --out=internal/contracts/curve_pool.go       

echo "✅ Contract bindings generated!"