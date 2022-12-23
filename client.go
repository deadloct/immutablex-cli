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

func (c *Client) GetAsset(ctx context.Context, collectionAddr, id string) (*api.Asset, error) {
	includeFees := false
	return c.imxClient.GetAsset(ctx, collectionAddr, id, &includeFees)
}

type GetAssetsRequest struct {
	Assets         []api.AssetWithOrders
	Before         string
	CollectionAddr string
	Cursor         string
}

func (c *Client) GetAssets(
	ctx context.Context,
	cfg *GetAssetsRequest,
) ([]api.AssetWithOrders, error) {

	req := c.imxClient.NewListAssetsRequest(ctx).
		Collection(cfg.CollectionAddr).
		Cursor(cfg.Cursor).
		PageSize(MaxAssetsPerReq).
		OrderBy("updated_at")

	if cfg.Before != "" {
		req = req.UpdatedMaxTimestamp(cfg.Before)
	}

	resp, err := c.imxClient.ListAssets(&req)
	if err != nil {
		return nil, err
	}

	if len(resp.Result) == 0 {
		return cfg.Assets, nil
	}

	cfg.Assets = append(cfg.Assets, resp.Result...)
	cfg.Cursor = resp.Cursor

	first := *resp.Result[0].UpdatedAt.Get()
	last := *resp.Result[len(resp.Result)-1].UpdatedAt.Get()
	log.Printf("fetched %v assets from %v to %v\n", len(resp.Result), first, last)

	if resp.Remaining > 0 {
		return c.GetAssets(ctx, cfg)
	}

	// Attempt to fetch earlier assets
	if len(resp.Result) > 0 {
		cfg.Before = last
		return c.GetAssets(ctx, cfg)
	}

	return cfg.Assets, nil
}
