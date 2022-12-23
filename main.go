package main

import (
	"context"
	"log"
	"os"
)

var AlchemyKey string

func init() {
	AlchemyKey = os.Getenv("ALCHEMY_API_KEY")
	if AlchemyKey == "" {
		log.Panic("no alchemy api key provided, get one at alchemy.com")
	}
}

func printAssetCounts(client *Client, name, addr string) {
	assets, err := client.GetAssets(context.Background(), addr, "", nil, "")
	if err != nil {
		log.Printf("failed to get assets: %v\n", err)
		return
	}

	rarityCounts := make(map[string]int, 4)
	for _, asset := range assets {
		rarity, ok := asset.Metadata["Rarity"].(string)
		if !ok {
			log.Printf("failed to find rarity for asset %#v", asset)
			continue
		}

		if asset.Name.IsSet() {
			rarityCounts[rarity]++
			rarityCounts["Total"]++
		} else {
			log.Println("asset skipped since it has no name and must be messed up")
		}
	}

	log.Printf("Counts for '%s': %#v\n", name, rarityCounts)
}

func main() {
	c, err := NewClient(AlchemyKey)
	if err != nil {
		log.Panic(err)
	}

	defer c.Stop()

	addrs := map[string]string{
		"BitVerse Portals": "0xe4ac52f4b4a721d1d0ad8c9c689df401c2db7291",
		"BitVerse Heroes":  "0x6465ef3009f3c474774f4afb607a5d600ea71d95",
	}

	for name, addr := range addrs {
		printAssetCounts(c, name, addr)
	}
}
