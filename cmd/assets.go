package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/deadloct/bh-imx-browser/lib"
	"github.com/spf13/cobra"
)

var (
	owner  string
	rarity string
	status string

	assetsCmd = &cobra.Command{
		Use:   "assets",
		Short: "Retrieve BitVerse assets in bulk",
		Long:  `Retrieve BitVerse assets with filters and in bulk from Immutable`,
		Run:   runAssetsCMD,
	}
)

func runAssetsCMD(cmd *cobra.Command, args []string) {
	assetManager := lib.NewAssetManager()
	if err := assetManager.Start(); err != nil {
		log.Panic(err)
	}
	defer assetManager.Stop()

	filter := &lib.AssetFilter{
		Owner:  owner,
		Rarity: lib.AssetRarity(strings.Title(rarity)),
		Status: lib.AssetStatus(status),
	}

	for _, collection := range lib.Collections {
		req := &lib.GetAssetsRequest{CollectionAddr: collection.Addr}
		assets, err := assetManager.GetAssets(context.Background(), req)
		if err != nil {
			fmt.Printf("error retrieving assets: %v\n", err)
		}

		collectionFilter := *filter
		collectionFilter.CollectionAddr = collection.Addr
		collectionFilter.CollectionName = collection.Name
		filtered := assetManager.FilterAssets(assets, &collectionFilter)

		if verbose {
			assetManager.PrintAssets(collection.Addr, filtered)
		}

		assetManager.PrintAssetCounts(collection.Name, filtered)
	}
}

func init() {
	rootCmd.AddCommand(assetsCmd)
	assetsCmd.Flags().StringVarP(&owner, "owner", "o", "", "Filter by owner")
	assetsCmd.Flags().StringVarP(&rarity, "rarity", "r", "", "Filter by rarity")
	assetsCmd.Flags().StringVarP(&status, "status", "s", "", "Filter by status")
}
