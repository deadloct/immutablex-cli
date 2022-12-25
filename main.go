package main

import (
	"os"

	"github.com/deadloct/immutablex-cli/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)

	cmd.Execute()
}
