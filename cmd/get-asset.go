package cmd

import (
	"context"
	"os"

	"github.com/deadloct/immutablex-cli/lib"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	getAssetTokenAddress string
	getAssetTokenID      string
	getAssetIncludeFees  bool

	assetCmd = &cobra.Command{
		Use:    "get-asset",
		Short:  "Retrieve asset (NFT) information",
		Long:   `Queries the ImmutableX getAsset endpoint for detailed asset information, see https://docs.x.immutable.com/reference/#/operations/getAsset`,
		PreRun: SetupLogging,
		Run:    runGetAssetCMD,
	}
)

func runGetAssetCMD(cmd *cobra.Command, args []string) {
	assetManager := lib.NewAssetManager()
	if err := assetManager.Start(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer assetManager.Stop()

	collectionManager := lib.NewCollectionManager()
	if err := collectionManager.Start(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer collectionManager.Stop()

	if shortcut := collectionManager.GetShortcutByName(getAssetTokenAddress); shortcut != nil {
		getAssetTokenAddress = shortcut.Addr
	}

	asset, err := assetManager.GetAsset(context.Background(), getAssetTokenAddress,
		getAssetTokenID, getAssetIncludeFees)
	if err != nil {
		log.Error("failed to retrieve asset: %v", err)
		os.Exit(1)
	}

	assetManager.PrintAsset(asset)
}

func init() {
	rootCmd.AddCommand(assetCmd)
	assetCmd.Flags().StringVarP(&getAssetTokenAddress, "token-address", "a", "",
		"address of the collection or shortcut")
	assetCmd.Flags().StringVarP(&getAssetTokenID, "token-id", "i", "", "id of the asset")
	assetCmd.Flags().BoolVarP(&getAssetIncludeFees, "include-fees", "f", false,
		"include fees associated with the asset")

	assetCmd.MarkFlagRequired("token-address")
	assetCmd.MarkFlagRequired("token-id")
}
