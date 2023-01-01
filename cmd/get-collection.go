package cmd

import (
	"context"
	"os"

	"github.com/deadloct/immutablex-go-lib/collections"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	getCollectionAddress string

	getCollectionCmd = &cobra.Command{
		Use:    "get-collection",
		Short:  "Retrieve details collection information",
		Long:   `Queries the ImmutableX getCollection endpoint for detailed collection information, see https://docs.x.immutable.com/reference/#/operations/getCollection`,
		PreRun: PreRun,
		Run:    runGetCollectionCMD,
	}
)

func runGetCollectionCMD(cmd *cobra.Command, args []string) {
	client := collections.NewClient(collections.NewClientConfig(alchemyKey))
	if err := client.Start(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer client.Stop()

	collection, err := client.GetCollection(context.Background(), getCollectionAddress)
	if err != nil {
		log.Error("failed to retrieve collection %s: %v", getCollectionAddress, err)
		os.Exit(1)
	}

	collections.PrintCollection(collection, output)
}

func init() {
	rootCmd.AddCommand(getCollectionCmd)
	getCollectionCmd.Flags().StringVarP(&getCollectionAddress, "collection", "c", "", "address of the collection")
	getAssetCmd.MarkFlagRequired("collection")
}
