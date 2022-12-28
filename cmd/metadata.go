package cmd

import (
	"encoding/json"
	"strings"

	log "github.com/sirupsen/logrus"
)

func jsonEncodeMetadata(metadata []string) string {
	metamap := make(map[string][]string, len(metadata))
	for _, item := range metadata {
		parts := strings.SplitN(item, "=", 2)
		if len(parts) != 2 {
			log.Debugf("could not parse metadata item %s into a key=value pair", item)
			continue
		}

		metamap[parts[0]] = append(metamap[parts[0]], parts[1])
	}

	data, err := json.Marshal(metamap)
	if err != nil {
		log.Debugf("skipping metamata completely because it could not be converted to json: %v", err)
	}

	return string(data)
}
