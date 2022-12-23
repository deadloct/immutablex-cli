package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func usage() {
	fmt.Printf(`
Usage: %s [options], where options are:
	assets [-rarity rarity] [-owner addr] [-status imx-token-status]
	asset -type portal|hero -id id
`, os.Args[0])

	os.Exit(1)
}

func main() {
	assetsListCmd := flag.NewFlagSet("assets", flag.ExitOnError)
	rarity := assetsListCmd.String("rarity", "", "Filter by rarity")
	owner := assetsListCmd.String("owner", "", "Filter by owner")
	status := assetsListCmd.String("status", "", "Filter by status")
	verboseList := assetsListCmd.Bool("v", false, "Print each item in the collection")

	assetCommand := flag.NewFlagSet("asset", flag.ExitOnError)
	assetType := assetCommand.String("type", "", "Either hero or portal")
	assetID := assetCommand.String("id", "", "The NFT ID")

	if len(os.Args) < 2 {
		usage()
	}

	assetManager := NewAssetManager()
	if err := assetManager.Start(); err != nil {
		log.Panic(err)
	}
	defer assetManager.Stop()

	switch os.Args[1] {
	case "assets":
		assetsListCmd.Parse(os.Args[2:])
		filter := &AssetFilter{
			Owner:  *owner,
			Rarity: AssetRarity(strings.Title(*rarity)),
			Status: AssetStatus(*status),
		}

		for _, collection := range Collections {
			req := &GetAssetsRequest{CollectionAddr: collection.Addr}
			assets, err := assetManager.GetAssets(context.Background(), req)
			if err != nil {
				fmt.Printf("error retrieving assets: %v\n", err)
			}

			collectionFilter := *filter
			collectionFilter.CollectionAddr = collection.Addr
			collectionFilter.CollectionName = collection.Name
			filtered := assetManager.FilterAssets(assets, &collectionFilter)

			if *verboseList {
				assetManager.PrintAssets(filtered)
			}

			assetManager.PrintAssetCounts(collection.Name, filtered)
		}

	case "asset":
		assetCommand.Parse(os.Args[2:])
		if *assetType == "" || *assetID == "" {
			usage()
		}

		if *assetType != "hero" && *assetType != "portal" {
			usage()
		}

		id := *assetID
		addr := Collections[*assetType].Addr
		asset, err := assetManager.GetAsset(context.Background(), addr, id)
		if err != nil {
			fmt.Printf("failed to retrieve asset: %v", err)
			os.Exit(1)
		}

		assetManager.PrintAsset(asset)

	default:
		fmt.Printf("invalid command line option '%s'", os.Args[1])
	}
}
