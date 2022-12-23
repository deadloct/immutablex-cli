package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

type AssetManager struct {
	client IMXClientWrapper
}

func NewAssetManager() *AssetManager {
	return &AssetManager{
		client: NewClient(),
	}
}

func (am *AssetManager) Start() error {
	return am.client.Start()
}

func (am *AssetManager) Stop() {
	am.client.Stop()
}

func (am *AssetManager) GetAsset(ctx context.Context, collectionAddr, id string) (*api.Asset, error) {
	includeFees := false
	return am.client.GetClient().GetAsset(ctx, collectionAddr, id, &includeFees)
}

type GetAssetsRequest struct {
	Assets          []api.AssetWithOrders
	Before          string
	CollectionAddr  string
	Cursor          string
	MetadataFilters map[string]string
}

func (am *AssetManager) GetAssets(
	ctx context.Context,
	cfg *GetAssetsRequest,
) ([]api.AssetWithOrders, error) {

	req := am.client.GetClient().NewListAssetsRequest(ctx).
		Collection(cfg.CollectionAddr).
		PageSize(MaxAssetsPerReq).
		OrderBy("updated_at")

	if cfg.Before != "" {
		req = req.UpdatedMaxTimestamp(cfg.Before)
	}

	if cfg.Cursor != "" {
		req = req.Cursor(cfg.Cursor)
	}

	if cfg.MetadataFilters != nil {
		log.Printf("skipping metadata since it is not currently supported")

		// The api doesn't like this with { "Rarity": "Legendary" }
		// metadata, err := json.Marshal(cfg.MetadataFilters)
		// if err != nil {
		// 	log.Printf("skipping metadata since it cannot be serialized")
		// } else {
		// 	req = req.Metadata(url.QueryEscape(string(metadata)))
		// }
	}

	resp, err := am.client.GetClient().ListAssets(&req)
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
		return am.GetAssets(ctx, cfg)
	}

	// Attempt to fetch earlier assets
	if len(resp.Result) > 0 {
		cfg.Before = last
		return am.GetAssets(ctx, cfg)
	}

	return cfg.Assets, nil
}

func (am *AssetManager) PrintAsset(collectionType, id string) {
	addr := Collections[collectionType].Addr
	asset, err := am.GetAsset(context.Background(), addr, id)
	if err != nil {
		fmt.Printf("failed to retrieve asset: %v", err)
		os.Exit(1)
	}

	fmt.Println(FormatAssetInfo(asset))
}

func (am *AssetManager) PrintAssetCounts(name, collectionAddr string) {
	assets, err := am.GetAssets(context.Background(), &GetAssetsRequest{
		CollectionAddr: collectionAddr,
	})
	if err != nil {
		fmt.Printf("failed to get assets: %v\n", err)
		return
	}

	counts := make(map[string]int, 4)
	for _, asset := range assets {
		rarity, ok := asset.Metadata["Rarity"].(string)
		if !ok {
			log.Printf("asset %s skipped because it doesn't have a rarity\n", asset.TokenId)
			continue
		}

		if !asset.Name.IsSet() {
			log.Printf("asset %s skipped since it has no name and must be messed up\n", asset.TokenId)
		}

		counts[rarity]++
		counts["Total"]++
	}

	fmt.Println(FormatAssetCounts(name, counts))
}
