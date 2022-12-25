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
	listCollectionsBlacklist string
	listCollectionsDirection string
	listCollectionsKeyword   string
	listCollectionsOrderBy   string
	listCollectionsWhitelist string

	listCollectionsCmd = &cobra.Command{
		Use:    "list-collections",
		Short:  "List NFT collections",
		Long:   `Queries the ImmutableX listCollections endpoint for retrieving collections in bulk, see https://docs.x.immutable.com/reference/#/operations/listCollections`,
		PreRun: PreCmd,
		Run:    runListCollectionsCMD,
	}
)

func runListCollectionsCMD(cmd *cobra.Command, args []string) {
	collectionManager := lib.NewCollectionManager()
	if err := collectionManager.Start(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer collectionManager.Stop()

	cfg := &lib.ListCollectionsConfig{
		Blacklist: listCollectionsBlacklist,
		Direction: listCollectionsDirection,
		Keyword:   listCollectionsKeyword,
		OrderBy:   listCollectionsOrderBy,
		Whitelist: listCollectionsWhitelist,
	}

	collections, err := collectionManager.ListCollections(context.Background(), cfg)
	if err != nil {
		log.Error("error retrieving collections: %v\n", err)
		os.Exit(1)
	}

	collectionManager.PrintCollections(collections, verbose)
	fmt.Printf("%d total collections returned", len(collections))
}

func init() {
	rootCmd.AddCommand(listCollectionsCmd)

	listCollectionsCmd.Flags().StringVarP(&listCollectionsBlacklist, "blacklist", "b", "", "comma-separated collections to exclude")
	listCollectionsCmd.Flags().StringVarP(&listCollectionsDirection, "direction", "d", "", "asc|desc")
	listCollectionsCmd.Flags().StringVarP(&listCollectionsKeyword, "keyword", "k", "", "search by name and description")
	listCollectionsCmd.Flags().StringVarP(&listCollectionsOrderBy, "order-by", "o", "updated_at", "updated_at|name")
	listCollectionsCmd.Flags().StringVarP(&listCollectionsWhitelist, "whitelist", "w", "", "comma-separated collections to only include")
}
