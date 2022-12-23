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

func (c *Client) PrintAssetCounts(ctx context.Context, addr string) {
	assets, err := c.GetAssets(ctx, addr, nil, "")
	if err != nil {
		log.Printf("failed to get assets: %v\n", err)
		return
	}

	rarityCounts := make(map[string]int, 4)
	for _, asset := range assets {
		rarity, ok := asset.Metadata["Rarity"].(string)
		if !ok {
			log.Printf("failed to find rarity for asset %#v", asset)
			continue
		}

		if asset.Name.IsSet() {
			rarityCounts[rarity]++
		} else {
			log.Println("asset skipped since it has no name and must be messed up")
		}
	}

	log.Printf("Rarities for collection %s: %#v\n", addr, rarityCounts)
}

func (c *Client) GetAssets(ctx context.Context, addr string, assets []api.AssetWithOrders, cursor string) ([]api.AssetWithOrders, error) {
	log.Printf("fetching %d more assets\n", MaxAssetsPerReq)

	req := c.imxClient.NewListAssetsRequest(ctx).
		Collection(addr).
		Cursor(cursor).
		PageSize(MaxAssetsPerReq).
		OrderBy("updated_at")

	resp, err := c.imxClient.ListAssets(&req)

	if err != nil {
		log.Printf("failed to get assets for addr %s: %v", addr, err)
		return nil, err
	}

	first := *resp.Result[0].UpdatedAt.Get()
	last := *resp.Result[len(resp.Result)-1].UpdatedAt.Get()
	log.Printf("fetched %v assets from %v to %v\n", len(resp.Result), first, last)

	assets = append(assets, resp.Result...)
	cursor = resp.Cursor

	if resp.Remaining > 0 {
		log.Printf("fetching more asset pages (%v)...", resp.Remaining)
		return c.GetAssets(ctx, addr, assets, cursor)
	}

	return assets, nil
}
