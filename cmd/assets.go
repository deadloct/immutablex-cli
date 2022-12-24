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
	collectionAddr string

	// BitVerse specific, need to make these freeform for other NFT collections
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

	collectionManager := lib.NewCollectionManager()
	if err := collectionManager.Start(); err != nil {
		log.Panic(err)
	}
	defer collectionManager.Stop()

	if shortcut := collectionManager.GetShortcutByName(collectionAddr); shortcut != nil {
		collectionAddr = shortcut.Addr
	}

	filter := &lib.AssetFilter{
		Owner:  owner,
		Rarity: lib.AssetRarity(strings.Title(rarity)),
		Status: lib.AssetStatus(status),
	}

	req := &lib.GetAssetsRequest{CollectionAddr: collectionAddr}
	assets, err := assetManager.GetAssets(context.Background(), req)
	if err != nil {
		fmt.Printf("error retrieving assets for collection %s: %v\n", collectionAddr, err)
	}

	collectionFilter := *filter
	collectionFilter.CollectionAddr = collectionAddr
	filtered := assetManager.FilterAssets(assets, &collectionFilter)

	if verbose {
		assetManager.PrintAssets(collectionAddr, filtered)
	}

	assetManager.PrintAssetCounts(collectionAddr, filtered)
}

func init() {
	rootCmd.AddCommand(assetsCmd)
	assetsCmd.Flags().StringVarP(&collectionAddr, "addr", "a", "", "Address of the collection or shortcut")
	assetCmd.MarkFlagRequired("addr")

	assetsCmd.Flags().StringVarP(&owner, "owner", "o", "", "Filter by owner")
	assetsCmd.Flags().StringVarP(&rarity, "rarity", "r", "", "Filter by rarity")
	assetsCmd.Flags().StringVarP(&status, "status", "s", "", "Filter by status")
}
