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

	log.Printf("fetching %d more assets\n", MaxAssetsPerReq)

	req := c.imxClient.NewListAssetsRequest(ctx).
		Collection(addr).
		Cursor(cursor).
		PageSize(MaxAssetsPerReq).
		OrderBy("updated_at")

	if before != "" {
		log.Printf("fetching starting at %v...", before)
		req = req.UpdatedMaxTimestamp(before)
	}

	resp, err := c.imxClient.ListAssets(&req)
	if err != nil {
		log.Printf("failed to get assets for addr %s: %v", addr, err)
		return nil, err
	}

	if len(resp.Result) == 0 {
		log.Println("no assets in this batch")
		return assets, nil
	}

	assets = append(assets, resp.Result...)

	first := *resp.Result[0].UpdatedAt.Get()
	last := *resp.Result[len(resp.Result)-1].UpdatedAt.Get()
	log.Printf("fetched %v assets from %v to %v\n", len(resp.Result), first, last)

	if resp.Remaining > 0 {
		log.Println("fetching more asset pages...")
		return c.GetAssets(ctx, addr, before, assets, resp.Cursor)
	}

	// Attempt to fetch earlier assets
	if len(resp.Result) > 0 {
		log.Printf("attempting to fetch records older than %v...\n", last)
		return c.GetAssets(ctx, addr, last, assets, resp.Cursor)
	}

	return assets, nil
}
