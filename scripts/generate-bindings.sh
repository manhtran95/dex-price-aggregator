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

echo "✅ Contract bindings generated!"