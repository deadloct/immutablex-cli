package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool
	output  string

	rootCmd = &cobra.Command{
		Use:    "immutablex-cli",
		Short:  "Interacts with the immutablex blockchain",
		Long:   "Helps retrieve information from the ImmutableX layer 2 blockchain",
		PreRun: PreRun,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable debug logging")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Changes output to these options: normal|json")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
