package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/deadloct/immutablex-cli/lib"
	libassets "github.com/deadloct/immutablex-cli/lib/assets"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	listAssetsBuyOrders           bool
	listAssetsCollection          string
	listAssetsDirection           string
	listAssetsIncludeFees         bool
	listAssetsName                string
	listAssetsOrderBy             string
	listAssetsSellOrders          bool
	listAssetsStatus              string
	listAssetsUpdatedMaxTimestamp string
	listAssetsUpdatedMinTimestamp string
	listAssetsUser                string

	assetsCmd = &cobra.Command{
		Use:    "list-assets",
		Short:  "List assets (NFTs) in bulk",
		Long:   `Queries the ImmutableX listAssets endpoint for retrieving assets in bulk, see https://docs.x.immutable.com/reference/#/operations/listAssets`,
		PreRun: PreRun,
		Run:    runListAssetsCMD,
	}
)

func runListAssetsCMD(cmd *cobra.Command, args []string) {
	client := libassets.NewClient(libassets.NewClientConfig(alchemyKey))
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

	if shortcut := collectionManager.GetShortcutByName(listAssetsCollection); shortcut != nil {
		listAssetsCollection = shortcut.Addr
	}

	cfg := libassets.ListAssetsConfig{
		BuyOrders:           listAssetsBuyOrders,
		Collection:          listAssetsCollection,
		Direction:           listAssetsDirection,
		IncludeFees:         listAssetsIncludeFees,
		Name:                listAssetsName,
		OrderBy:             listAssetsOrderBy,
		SellOrders:          listAssetsSellOrders,
		Status:              listAssetsStatus,
		UpdatedMaxTimestamp: listAssetsUpdatedMaxTimestamp,
		UpdatedMinTimestamp: listAssetsUpdatedMinTimestamp,
		User:                listAssetsUser,
	}

	assetsMetadata, err := cmd.Flags().GetStringArray("metadata")
	if err != nil {
		log.Debugf("unable to parse metadata: %v\n", err)
	} else {
		cfg.Metadata = assetsMetadata
	}

	assets, err := client.ListAssets(context.Background(), cfg)
	if err != nil {
		log.Error("error retrieving assets for collection %s: %v\n", listAssetsCollection, err)
		os.Exit(1)
	}

	libassets.PrintAssets(listAssetsCollection, assets)
	fmt.Printf("%d total assets returned", len(assets))
}

func init() {
	rootCmd.AddCommand(assetsCmd)

	assetsCmd.Flags().BoolVarP(&listAssetsBuyOrders, "buy-orders", "b", false, "Retrieve buy orders for each asset")
	assetsCmd.Flags().StringVarP(&listAssetsCollection, "collection", "c", "", "Address of the collection or shortcut")
	assetsCmd.Flags().StringVarP(&listAssetsDirection, "direction", "d", "", "asc|desc")
	assetsCmd.Flags().BoolVarP(&listAssetsIncludeFees, "include-fees", "i", false, "Retrieves fees for each asset")
	assetsCmd.Flags().StringArrayP("metadata", "m", nil,
		`Filter by metadata in key=value format (repeatable). For example `+
			`"immutable-cli assets -m Rarity=Mythic -m Generation=0. Note that metadata `+
			`keys and values are case sensitive.`)
	assetsCmd.Flags().StringVarP(&listAssetsName, "name", "n", "desc", "Search for this asset name")
	assetsCmd.Flags().StringVarP(&listAssetsOrderBy, "order-by", "o", "updated_at", "updated_at|name")
	assetsCmd.Flags().BoolVarP(&listAssetsSellOrders, "sell-orders", "l", false, "Retrieves sell orders for each asset")
	assetsCmd.Flags().StringVarP(&listAssetsStatus, "status", "s", "", "Filter by the status: eth|imx|preparing_withdrawal|withdrawable|burned")
	assetsCmd.Flags().StringVarP(&listAssetsUpdatedMaxTimestamp, "updated-max-timestamp", "x", "", "Include results on or before this time in ISO 8601 UTC format")
	assetsCmd.Flags().StringVarP(&listAssetsUpdatedMinTimestamp, "updated-min-timestamp", "z", "", "Include results on or after this time in ISO 8601 UTC format")
	assetsCmd.Flags().StringVarP(&listAssetsUser, "user", "u", "", "Retrieves assets owned by this user/wallet address")

	assetsCmd.MarkFlagRequired("collection")
}
