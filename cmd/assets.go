package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/deadloct/immutablex-cli/lib"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	assetsBuyOrders           bool
	assetsCollection          string
	assetsDirection           string
	assetsIncludeFees         bool
	assetsName                string
	assetsOrderBy             string
	assetsSellOrders          bool
	assetsStatus              string
	assetsUpdatedMaxTimestamp string
	assetsUpdatedMinTimestamp string
	assetsUser                string

	assetsCmd = &cobra.Command{
		Use:    "assets",
		Short:  "List assets (NFTs) in bulk",
		Long:   `Queries the ImmutableX listAssets endpoint for retrieving assets in bulk, see https://docs.x.immutable.com/reference/#/operations/listAssets`,
		PreRun: SetupLogging,
		Run:    runAssetsCMD,
	}
)

func runAssetsCMD(cmd *cobra.Command, args []string) {
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

	if shortcut := collectionManager.GetShortcutByName(assetsCollection); shortcut != nil {
		assetsCollection = shortcut.Addr
	}

	req := &lib.GetAssetsRequest{
		BuyOrders:           assetsBuyOrders,
		Collection:          assetsCollection,
		Direction:           assetsDirection,
		IncludeFees:         assetsIncludeFees,
		Name:                assetsName,
		OrderBy:             assetsOrderBy,
		SellOrders:          assetsSellOrders,
		Status:              assetsStatus,
		UpdatedMaxTimestamp: assetsUpdatedMaxTimestamp,
		UpdatedMinTimestamp: assetsUpdatedMinTimestamp,
		User:                assetsUser,
	}

	assetsMetadata, err := cmd.Flags().GetStringArray("metadata")
	if err != nil {
		log.Debugf("unable to parse metadata: %v\n", err)
	} else {
		req.Metadata = assetsMetadata
	}

	assets, err := assetManager.GetAssets(context.Background(), req)
	if err != nil {
		log.Error("error retrieving assets for collection %s: %v\n", assetsCollection, err)
		os.Exit(1)
	}

	assetManager.PrintAssets(assetsCollection, assets)
	fmt.Printf("%d total assets returned", len(assets))
}

func init() {
	rootCmd.AddCommand(assetsCmd)

	assetsCmd.Flags().BoolVarP(&assetsBuyOrders, "buy-orders", "b", false, "Retrieve buy orders for each asset")
	assetsCmd.Flags().StringVarP(&assetsCollection, "collection", "c", "", "Address of the collection or shortcut")
	assetsCmd.Flags().StringVarP(&assetsDirection, "direction", "d", "", "asc|desc")
	assetsCmd.Flags().BoolVarP(&assetsIncludeFees, "include-fees", "i", false, "Retrieves fees for each asset")
	assetsCmd.Flags().StringArrayP("metadata", "m", nil,
		`Filter by metadata in key=value format (repeatable). For example `+
			`"immutable-cli assets -m Rarity=Mythic -m Generation=0. Note that metadata `+
			`keys and values are case sensitive.`)
	assetsCmd.Flags().StringVarP(&assetsName, "name", "n", "desc", "Search for this asset name")
	assetsCmd.Flags().StringVarP(&assetsOrderBy, "order-by", "o", "updated_at", "updated_at|name")
	assetsCmd.Flags().BoolVarP(&assetsSellOrders, "sell-orders", "l", false, "Retrieves sell orders for each asset")
	assetsCmd.Flags().StringVarP(&assetsStatus, "status", "s", "", "Filter by the status: eth|imx|preparing_withdrawal|withdrawable|burned")
	assetsCmd.Flags().StringVarP(&assetsUpdatedMaxTimestamp, "updated-max-timestamp", "x", "", "Include results on or before this time in ISO 8601 UTC format")
	assetsCmd.Flags().StringVarP(&assetsUpdatedMinTimestamp, "updated-min-timestamp", "z", "", "Include results on or after this time in ISO 8601 UTC format")
	assetsCmd.Flags().StringVarP(&assetsUser, "user", "u", "", "Retrieves assets owned by this user/wallet address")

	assetsCmd.MarkFlagRequired("collection")
}
