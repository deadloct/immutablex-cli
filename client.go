package main

import (
	"context"
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

const MaxAssetsPerReq = 200

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

func (c *Client) GetAssets(
	ctx context.Context,
	addr string,
	before string,
	assets []api.AssetWithOrders,
	cursor string,
) ([]api.AssetWithOrders, error) {

	req := c.imxClient.NewListAssetsRequest(ctx).
		Collection(addr).
		Cursor(cursor).
		PageSize(MaxAssetsPerReq).
		OrderBy("updated_at")

	if before != "" {
		req = req.UpdatedMaxTimestamp(before)
	}

	resp, err := c.imxClient.ListAssets(&req)
	if err != nil {
		return nil, err
	}

	if len(resp.Result) == 0 {
		return assets, nil
	}

	assets = append(assets, resp.Result...)

	first := *resp.Result[0].UpdatedAt.Get()
	last := *resp.Result[len(resp.Result)-1].UpdatedAt.Get()
	log.Printf("fetched %v assets from %v to %v\n", len(resp.Result), first, last)

	if resp.Remaining > 0 {
		return c.GetAssets(ctx, addr, before, assets, resp.Cursor)
	}

	// Attempt to fetch earlier assets
	if len(resp.Result) > 0 {
		return c.GetAssets(ctx, addr, last, assets, resp.Cursor)
	}

	return assets, nil
}
