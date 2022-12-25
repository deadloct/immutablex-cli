package cmd

import (
	"context"
	"os"

	"github.com/deadloct/immutablex-cli/lib"
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
	collectionManager := lib.NewCollectionManager()
	if err := collectionManager.Start(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer collectionManager.Stop()

	if shortcut := collectionManager.GetShortcutByName(getCollectionAddress); shortcut != nil {
		getCollectionAddress = shortcut.Addr
	}

	collection, err := collectionManager.GetCollection(context.Background(), getCollectionAddress)
	if err != nil {
		log.Error("failed to retrieve collection %s: %v", getCollectionAddress, err)
		os.Exit(1)
	}

	collectionManager.PrintCollection(collection)
}

func init() {
	rootCmd.AddCommand(getCollectionCmd)
	getCollectionCmd.Flags().StringVarP(&getCollectionAddress, "collection", "c", "", "address of the collection")
	getAssetCmd.MarkFlagRequired("collection")
}
