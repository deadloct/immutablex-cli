package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deadloct/bh-imx-browser/imxapi"
	"github.com/spf13/cobra"
)

var (
	assetType string
	id        string

	assetCmd = &cobra.Command{
		Use:   "asset",
		Short: "Retrieve BitVerse asset information",
		Long:  `Retrieve a specific BitVerse asset from Immutable`,
		Run:   runAssetCMD,
	}
)

func runAssetCMD(cmd *cobra.Command, args []string) {
	assetManager := imxapi.NewAssetManager()
	if err := assetManager.Start(); err != nil {
		log.Panic(err)
	}
	defer assetManager.Stop()

	if assetType != "hero" && assetType != "portal" {
		cmd.Help()
		os.Exit(1)
	}

	addr := imxapi.Collections[assetType].Addr
	asset, err := assetManager.GetAsset(context.Background(), addr, id)
	if err != nil {
		fmt.Printf("failed to retrieve asset: %v", err)
		os.Exit(1)
	}

	assetManager.PrintAsset(asset)
}

func init() {
	rootCmd.AddCommand(assetCmd)
	assetCmd.Flags().StringVarP(&assetType, "type", "t", "hero", "Type")
	assetCmd.MarkFlagRequired("type")
	assetCmd.Flags().StringVarP(&id, "id", "i", "", "")
	assetCmd.MarkFlagRequired("id")
}
