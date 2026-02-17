# DEX Price Aggregator

A Go-based service that aggregates prices from multiple decentralized exchanges (DEXs).

## Features

- Compare prices across multiple DEXs (Uniswap V2, Uniswap V3, SushiSwap)
- Find the best swap rates
- RESTful API

## Getting Started

### Prerequisites

- Go 1.21+
- Ethereum RPC endpoint (Infura, Alchemy, or local node)

### Installation

1. Clone the repository
2. Copy `.env.example` to `.env` and add your Ethereum RPC URL
3. Install dependencies: `make install`
4. Run the server: `make run`

### API Endpoints

- `GET /api/v1/health` - Health check
- `POST /api/v1/quote` - Get best quote for a swap
- `POST /api/v1/compare` - Compare prices across all DEXs

## Development
```bash
make run    # Run the server
make build  # Build binary
make test   # Run tests
```