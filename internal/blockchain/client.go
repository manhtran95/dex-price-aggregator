package blockchain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	client *ethclient.Client
}

func NewClient(rpcURL string) (*Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

func (c *Client) Close() {
	c.client.Close()
}

func (c *Client) GetBlockNumber(ctx context.Context) (*big.Int, error) {
	header, err := c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

func (c *Client) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	return c.client.BalanceAt(ctx, address, nil)
}

func (c *Client) Client() *ethclient.Client {
	return c.client
}