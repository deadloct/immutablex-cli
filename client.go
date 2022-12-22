package main

import (
	"context"
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

type Client struct {
	imxClient *imx.Client
}

func NewClient(alchemyKey string) (*Client, error) {
	cfg := imx.Config{
		AlchemyAPIKey: alchemyKey,
		APIConfig:     api.NewConfiguration(),
		Environment:   imx.Mainnet,
	}

	c, err := imx.NewClient(&cfg)
	if err != nil {
		return nil, err
	}

	return &Client{imxClient: c}, nil
}

func (c *Client) Stop() {
	c.imxClient.EthClient.Close()
}

func (c *Client) PrintAssetCounts(ctx context.Context, addr string) {
	req := c.imxClient.NewListAssetsRequest(ctx).Collection(addr).PageSize(9999)

	resp, err := c.imxClient.ListAssets(&req)
	if err != nil {
		log.Panicf("error when listing assets: %v\n", err)
	}

	rarityCounts := make(map[string]int, 4)
	for _, asset := range resp.Result {
		rarityCounts[asset.Metadata["Rarity"].(string)]++
	}

	log.Printf("Rarities for collection %s: %#v\n", addr, rarityCounts)
}
