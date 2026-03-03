# DEX Price Aggregator

A Go-based service that aggregates prices from multiple decentralized exchanges (DEXs).

## Features

- **Multi-DEX Support** — Uniswap V2, Uniswap V3, SushiSwap, Curve
- **Multi-hop Graph-based Routing** — Finds optimal swap paths across up to 3 hops using a token graph
- **Gas Estimation** — Estimates gas cost per route and factors it into best-route selection
- **Rate Limiting** — Protects the API from abuse with configurable request rate limits
- **Caching** — Caches quotes and token metadata to reduce RPC calls and improve latency
- **Historical Price Tracking** — Records and queries historical swap prices for analytics

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
