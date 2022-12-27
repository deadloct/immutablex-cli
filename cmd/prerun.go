package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var alchemyKey string

func PreRun(cmd *cobra.Command, args []string) {
	setupLogging()
	loadAlchemyKey()
}

func setupLogging() {
	log.SetFormatter(&log.TextFormatter{})
	if verbose {
		log.Info("verbose logs enabled")
		log.SetLevel(log.DebugLevel)
	}
}

func loadAlchemyKey() {
	alchemyKey = os.Getenv("ALCHEMY_API_KEY")
}
