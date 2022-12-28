package cmd

import (
	"context"
	"os"

	"github.com/deadloct/immutablex-cli/lib/collections"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	listCollectionsBlacklist string
	listCollectionsDirection string
	listCollectionsKeyword   string
	listCollectionsOrderBy   string
	listCollectionsWhitelist string

	listCollectionsCmd = &cobra.Command{
		Use:    "list-collections",
		Short:  "List NFT collections",
		Long:   `Queries the ImmutableX listCollections endpoint for retrieving collections in bulk, see https://docs.x.immutable.com/reference/#/operations/listCollections`,
		PreRun: PreRun,
		Run:    runListCollectionsCMD,
	}
)

func runListCollectionsCMD(cmd *cobra.Command, args []string) {
	client := collections.NewClient(collections.NewClientConfig(alchemyKey))
	if err := client.Start(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer client.Stop()

	cfg := &collections.ListCollectionsConfig{
		Blacklist: listCollectionsBlacklist,
		Direction: listCollectionsDirection,
		Keyword:   listCollectionsKeyword,
		OrderBy:   listCollectionsOrderBy,
		Whitelist: listCollectionsWhitelist,
	}

	result, err := client.ListCollections(context.Background(), cfg)
	if err != nil {
		log.Error("error retrieving collections: %v\n", err)
		os.Exit(1)
	}

	collections.PrintCollections(result, output)
	log.Debugf("%d total collections returned", len(result))
}

func init() {
	rootCmd.AddCommand(listCollectionsCmd)

	listCollectionsCmd.Flags().StringVarP(&listCollectionsBlacklist, "blacklist", "b", "", "comma-separated collections to exclude")
	listCollectionsCmd.Flags().StringVarP(&listCollectionsDirection, "direction", "d", "desc", "asc|desc")
	listCollectionsCmd.Flags().StringVarP(&listCollectionsKeyword, "keyword", "k", "", "search by name and description")
	listCollectionsCmd.Flags().StringVar(&listCollectionsOrderBy, "order-by", "updated_at", "updated_at|name")
	listCollectionsCmd.Flags().StringVarP(&listCollectionsWhitelist, "whitelist", "w", "", "comma-separated collections to only include")
}
