package cmd

import (
	"context"
	"os"

	"github.com/deadloct/immutablex-cli/lib"
	"github.com/deadloct/immutablex-cli/lib/assets"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	getAssetTokenAddress string
	getAssetTokenID      string
	getAssetIncludeFees  bool

	getAssetCmd = &cobra.Command{
		Use:    "get-asset",
		Short:  "Retrieve asset (NFT) information",
		Long:   `Queries the ImmutableX getAsset endpoint for detailed asset information, see https://docs.x.immutable.com/reference/#/operations/getAsset`,
		PreRun: PreRun,
		Run:    runGetAssetCMD,
	}
)

func runGetAssetCMD(cmd *cobra.Command, args []string) {
	client := assets.NewClient(assets.NewClientConfig(alchemyKey))
	if err := client.Start(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer client.Stop()

	collectionManager := lib.NewCollectionManager()
	if err := collectionManager.Start(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer collectionManager.Stop()

	if shortcut := collectionManager.GetShortcutByName(getAssetTokenAddress); shortcut != nil {
		getAssetTokenAddress = shortcut.Addr
	}

	asset, err := client.GetAsset(context.Background(), getAssetTokenAddress,
		getAssetTokenID, getAssetIncludeFees)
	if err != nil {
		log.Error("failed to retrieve asset: %v", err)
		os.Exit(1)
	}

	assets.PrintAsset(asset)
}

func init() {
	rootCmd.AddCommand(getAssetCmd)
	getAssetCmd.Flags().StringVarP(&getAssetTokenAddress, "token-address", "a", "",
		"address of the collection or shortcut")
	getAssetCmd.Flags().StringVarP(&getAssetTokenID, "token-id", "i", "", "id of the asset")
	getAssetCmd.Flags().BoolVarP(&getAssetIncludeFees, "include-fees", "f", false,
		"include fees associated with the asset")

	getAssetCmd.MarkFlagRequired("token-address")
	getAssetCmd.MarkFlagRequired("token-id")
}
