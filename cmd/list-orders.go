package cmd

import (
	"context"
	"os"

	"github.com/deadloct/immutablex-cli/lib/orders"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	listOrdersAuxiliaryFeePercentages string
	listOrdersAuxiliaryFeeRecipients  string
	listOrdersBuyAssetID              string
	listOrdersBuyMaxQuantity          string
	listOrdersBuyMetadata             []string
	listOrdersBuyMinQuantity          string
	listOrdersBuyTokenAddress         string
	listOrdersBuyTokenID              string
	listOrdersBuyTokenName            string
	listOrdersBuyTokenType            string
	listOrdersDirection               string
	listOrdersIncludeFees             bool
	listOrdersMaxTimestamp            string
	listOrdersMinTimestamp            string
	listOrdersOrderBy                 string
	listOrdersPageSize                int
	listOrdersSellAssetID             string
	listOrdersSellMaxQuantity         string
	listOrdersSellMetadata            []string
	listOrdersSellMinQuantity         string
	listOrdersSellTokenAddress        string
	listOrdersSellTokenID             string
	listOrdersSellTokenName           string
	listOrdersSellTokenType           string
	listOrdersStatus                  string
	listOrdersUpdatedMaxTimestamp     string
	listOrdersUpdatedMinTimestamp     string
	listOrdersUser                    string

	listOrdersCmd = &cobra.Command{
		Use:    "list-orders",
		Short:  "List active orders",
		Long:   `Queries the ImmutableX listOrders endpoint for retrieving orders in bulk, see https://docs.x.immutable.com/reference/#/operations/listOrders`,
		PreRun: PreRun,
		Run:    runListOrdersCMD,
	}
)

func runListOrdersCMD(cmd *cobra.Command, args []string) {
	client := orders.NewClient(orders.NewClientConfig(alchemyKey))
	if err := client.Start(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer client.Stop()

	cfg := &orders.ListOrdersConfig{
		AuxiliaryFeePercentages: listOrdersAuxiliaryFeePercentages,
		AuxiliaryFeeRecipients:  listOrdersAuxiliaryFeeRecipients,
		BuyAssetID:              listOrdersBuyAssetID,
		BuyMaxQuantity:          listOrdersBuyMaxQuantity,
		BuyMinQuantity:          listOrdersBuyMinQuantity,
		BuyTokenAddress:         listOrdersBuyTokenAddress,
		BuyTokenID:              listOrdersBuyTokenID,
		BuyTokenName:            listOrdersBuyTokenName,
		BuyTokenType:            listOrdersBuyTokenType,
		Direction:               listOrdersDirection,
		IncludeFees:             listOrdersIncludeFees,
		MaxTimestamp:            listOrdersMaxTimestamp,
		MinTimestamp:            listOrdersMinTimestamp,
		OrderBy:                 listOrdersOrderBy,
		PageSize:                listOrdersPageSize,
		SellAssetID:             listOrdersSellAssetID,
		SellMaxQuantity:         listOrdersSellMaxQuantity,
		SellMinQuantity:         listOrdersSellMinQuantity,
		SellTokenAddress:        listOrdersSellTokenAddress,
		SellTokenID:             listOrdersSellTokenID,
		SellTokenName:           listOrdersSellTokenName,
		SellTokenType:           listOrdersSellTokenType,
		Status:                  listOrdersStatus,
		UpdatedMaxTimestamp:     listOrdersUpdatedMaxTimestamp,
		UpdatedMinTimestamp:     listOrdersUpdatedMinTimestamp,
		User:                    listOrdersUser,
	}

	buyMetadata, err := cmd.Flags().GetStringArray("buy-metadata")
	if err != nil {
		log.Debugf("unable to parse buy metadata: %v\n", err)
	} else {
		cfg.BuyMetadata = jsonEncodeMetadata(buyMetadata)
	}

	sellMetadata, err := cmd.Flags().GetStringArray("sell-metadata")
	if err != nil {
		log.Debugf("unable to parse sell metadata: %v\n", err)
	} else {
		cfg.SellMetadata = jsonEncodeMetadata(sellMetadata)
	}

	result, err := client.ListOrders(context.Background(), cfg)
	if err != nil {
		log.Error("error retrieving orders: %v\n", err)
		os.Exit(1)
	}

	orders.PrintOrders(result, output)
	log.Debugf("%d total orders returned", len(result))
}

func init() {
	rootCmd.AddCommand(listOrdersCmd)

	listOrdersCmd.Flags().StringVar(&listOrdersAuxiliaryFeePercentages, "auxiliary-fee-percentages", "", "Comma separated string of fee percentages that are to be paired with auxiliary_fee_recipients")
	listOrdersCmd.Flags().StringVar(&listOrdersAuxiliaryFeeRecipients, "auxiliary-fee-recipients", "", "Comma separated string of fee recipients that are to be paired with auxiliary_fee_percentages")
	listOrdersCmd.Flags().StringVar(&listOrdersBuyAssetID, "buy-asset-id", "", "Internal IMX ID of the asset this order buys")
	listOrdersCmd.Flags().StringVar(&listOrdersBuyMaxQuantity, "buy-max-quantity", "", "Max quantity for the asset this order buys")
	listOrdersCmd.Flags().StringArray("buy-metadata", nil, "case-sensitive repeatable key=value formatted metadata filters for the asset this order buys")
	listOrdersCmd.Flags().StringVar(&listOrdersBuyMinQuantity, "buy-min-quantity", "", "Min quantity for the asset this order buys")
	listOrdersCmd.Flags().StringVar(&listOrdersBuyTokenAddress, "buy-token-address", "", "Token address of the asset this order buys")
	listOrdersCmd.Flags().StringVar(&listOrdersBuyTokenID, "buy-token-id", "", "ERC721 Token ID of the asset this order buys")
	listOrdersCmd.Flags().StringVar(&listOrdersBuyTokenName, "buy-token-name", "", "Token name of the asset this order buys")
	listOrdersCmd.Flags().StringVar(&listOrdersBuyTokenType, "buy-token-type", "", "Token type of the asset this order buys")
	listOrdersCmd.Flags().StringVar(&listOrdersDirection, "direction", "", "Direction to sort (options: asc|desc)")
	listOrdersCmd.Flags().BoolVar(&listOrdersIncludeFees, "include-fees", false, "Set flag to true to include fee object for orders")
	listOrdersCmd.Flags().StringVar(&listOrdersMaxTimestamp, "max-timestamp", "", "Maximum created at timestamp for this order, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z'")
	listOrdersCmd.Flags().StringVar(&listOrdersMinTimestamp, "min-timestamp", "", "Minimum created at timestamp for this order, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z'")
	listOrdersCmd.Flags().StringVar(&listOrdersOrderBy, "order-by", "", "Property to sort by (options: created_at|expired_at|sell_quantity|buy_quantity|buy_quantity_with_fees|updated_at")
	listOrdersCmd.Flags().IntVar(&listOrdersPageSize, "page-size", 20, "Page size of the result. Unofficial: if page-size is zero, all orders will be returned. Will attempt to retrive more pages up to the desired size.")
	listOrdersCmd.Flags().StringVar(&listOrdersSellAssetID, "sell-asset-id", "", "Internal IMX ID of the asset this order sells")
	listOrdersCmd.Flags().StringVar(&listOrdersSellMaxQuantity, "sell-max_quantity", "", "Max quantity for the asset this order sells")
	listOrdersCmd.Flags().StringArray("sell-metadata", nil, "case-sensitive repeatable key=value formatted metadata filters for the asset this order sells")
	listOrdersCmd.Flags().StringVar(&listOrdersSellMinQuantity, "sell_min-quantity", "", "Min quantity for the asset this order sells")
	listOrdersCmd.Flags().StringVar(&listOrdersSellTokenAddress, "sell-token-address", "", "Token address of the asset this order sells")
	listOrdersCmd.Flags().StringVar(&listOrdersSellTokenID, "sell-token-id", "", "ERC721 Token ID of the asset this order sells")
	listOrdersCmd.Flags().StringVar(&listOrdersSellTokenName, "sell-token-name", "", "Token name of the asset this order sells")
	listOrdersCmd.Flags().StringVar(&listOrdersSellTokenType, "sell-token-type", "", "Token type of the asset this order sells")
	listOrdersCmd.Flags().StringVar(&listOrdersStatus, "status", "", "Status of this order (options: active|filled|cancelled|expired|inactive)")
	listOrdersCmd.Flags().StringVar(&listOrdersUpdatedMaxTimestamp, "updated-max-timestamp", "", "Maximum updated at timestamp for this order, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z'")
	listOrdersCmd.Flags().StringVar(&listOrdersUpdatedMinTimestamp, "updated-min-timestamp", "", "Minimum updated at timestamp for this order, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z'")
	listOrdersCmd.Flags().StringVar(&listOrdersUser, "user", "", "Ethereum address of the user who submitted this order")
}
