package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

type Collection struct {
	Name string
	Addr string
}

var (
	AlchemyKey  string
	Collections = map[string]Collection{
		"portal": {
			Name: "BitVerse Portals",
			Addr: "0xe4ac52f4b4a721d1d0ad8c9c689df401c2db7291",
		},
		"hero": {
			Name: "BitVerse Heroes",
			Addr: "0x6465ef3009f3c474774f4afb607a5d600ea71d95",
		},
	}

	client *Client
)

func init() {
	AlchemyKey = os.Getenv("ALCHEMY_API_KEY")
	if AlchemyKey == "" {
		log.Panic("no alchemy api key provided, get one at alchemy.com")
	}
}

func printAssetCounts() {
	for _, collection := range Collections {
		assets, err := client.GetAssets(context.Background(), &GetAssetsRequest{
			CollectionAddr: collection.Addr,
		})
		if err != nil {
			fmt.Printf("failed to get assets: %v\n", err)
			continue
		}

		counts := make(map[string]int, 4)
		for _, asset := range assets {
			rarity, ok := asset.Metadata["Rarity"].(string)
			if !ok {
				log.Printf("asset %s skipped because it doesn't have a rarity\n", asset.TokenId)
				continue
			}

			if !asset.Name.IsSet() {
				log.Printf("asset %s skipped since it has no name and must be messed up\n", asset.TokenId)
			}

			counts[rarity]++
			counts["Total"]++
		}

		fmt.Println(FormatAssetCounts(collection.Name, counts))
	}
}

func printAssetInformation(collectionType, id string) {
	addr := Collections[collectionType].Addr
	asset, err := client.GetAsset(context.Background(), addr, id)
	if err != nil {
		fmt.Printf("failed to retrieve asset: %v", err)
		os.Exit(1)
	}

	fmt.Println(FormatAssetInfo(asset))
}

func usage() {
	fmt.Printf("Usage: %s [option], where options are:\n\tassets\t(print all assets)\n\tasset portal|hero id\t(display information about an NFT)\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	var err error
	client, err = NewClient(AlchemyKey)
	if err != nil {
		log.Panic(err)
	}

	defer client.Stop()

	switch os.Args[1] {
	case "assets":
		printAssetCounts()

	case "asset":
		if len(os.Args) < 4 {
			usage()
		}

		collectionType := os.Args[2]
		if collectionType != "hero" && collectionType != "portal" {
			usage()
		}

		printAssetInformation(collectionType, os.Args[3])

	default:
		fmt.Printf("invalid command line option '%s'", os.Args[1])
	}
}
