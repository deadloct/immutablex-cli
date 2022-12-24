package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deadloct/immutablex-cli/lib"
	"github.com/spf13/cobra"
)

var (
	assetAddr string
	assetID   string

	assetCmd = &cobra.Command{
		Use:   "asset",
		Short: "Retrieve BitVerse asset information",
		Long:  `Retrieve a specific BitVerse asset from Immutable`,
		Run:   runAssetCMD,
	}
)

func runAssetCMD(cmd *cobra.Command, args []string) {
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

	if shortcut := collectionManager.GetShortcutByName(assetAddr); shortcut != nil {
		assetAddr = shortcut.Addr
	}

	log.Printf("requesting asset %s from collection %s", assetID, assetAddr)

	asset, err := assetManager.GetAsset(context.Background(), assetAddr, assetID)
	if err != nil {
		fmt.Printf("failed to retrieve asset: %v", err)
		os.Exit(1)
	}

	assetManager.PrintAsset(asset)
}

func init() {
	rootCmd.AddCommand(assetCmd)
	assetCmd.Flags().StringVarP(&assetAddr, "addr", "a", "", "address of the collection or shortcut")
	assetCmd.MarkFlagRequired("addr")
	assetCmd.Flags().StringVarP(&assetID, "id", "i", "", "id of the asset")
	assetCmd.MarkFlagRequired("id")
}
