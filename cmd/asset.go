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
	assetTokenAddress string
	assetTokenID      string
	assetIncludeFees  bool

	assetCmd = &cobra.Command{
		Use:   "asset",
		Short: "Retrieve asset (NFT) information",
		Long:  `Queries the ImmutableX getAsset endpoint for detailed asset information, see https://docs.x.immutable.com/reference/#/operations/getAsset`,
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

	if shortcut := collectionManager.GetShortcutByName(assetTokenAddress); shortcut != nil {
		assetTokenAddress = shortcut.Addr
	}

	log.Printf("requesting asset %s from collection %s", assetTokenID, assetTokenAddress)

	asset, err := assetManager.GetAsset(context.Background(), assetTokenAddress,
		assetTokenID, assetIncludeFees)
	if err != nil {
		fmt.Printf("failed to retrieve asset: %v", err)
		os.Exit(1)
	}

	assetManager.PrintAsset(asset)
}

func init() {
	rootCmd.AddCommand(assetCmd)
	assetCmd.Flags().StringVarP(&assetTokenAddress, "token-address", "a", "",
		"address of the collection or shortcut")
	assetCmd.Flags().StringVarP(&assetTokenID, "token-id", "i", "", "id of the asset")
	assetCmd.Flags().BoolVarP(&assetIncludeFees, "include-fees", "f", false,
		"include fees associated with the asset")

	assetCmd.MarkFlagRequired("token-address")
	assetCmd.MarkFlagRequired("token-id")
}
