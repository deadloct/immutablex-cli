package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verbose bool

	rootCmd = &cobra.Command{
		Use:    "immutablex-cli",
		Short:  "Interacts with the immutablex blockchain",
		Long:   "Helps retrieve information from the ImmutableX layer 2 blockchain",
		PreRun: SetupLogging,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable debug logging")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func SetupLogging(cmd *cobra.Command, args []string) {
	log.SetFormatter(&log.TextFormatter{})
	if verbose {
		log.Info("verbose logs enabled")
		log.SetLevel(log.DebugLevel)
	}
}
