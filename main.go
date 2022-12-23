package main

import (
	"fmt"
	"log"
	"os"
)

var Collections = map[string]struct {
	Name string
	Addr string
}{
	"portal": {
		Name: "BitVerse Portals",
		Addr: "0xe4ac52f4b4a721d1d0ad8c9c689df401c2db7291",
	},
	"hero": {
		Name: "BitVerse Heroes",
		Addr: "0x6465ef3009f3c474774f4afb607a5d600ea71d95",
	},
}

func usage() {
	fmt.Printf("Usage: %s [option], where options are:\n\tassets\t(print all assets)\n\tasset portal|hero id\t(display information about an NFT)\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	assetManager := NewAssetManager()
	if err := assetManager.Start(); err != nil {
		log.Panic(err)
	}
	defer assetManager.Stop()

	switch os.Args[1] {
	case "assets":
		for _, collection := range Collections {
			assetManager.PrintAssetCounts(collection.Name, collection.Addr)
		}

	case "asset":
		if len(os.Args) < 4 {
			usage()
		}

		collectionType := os.Args[2]
		if collectionType != "hero" && collectionType != "portal" {
			usage()
		}

		assetManager.PrintAsset(collectionType, os.Args[3])

	default:
		fmt.Printf("invalid command line option '%s'", os.Args[1])
	}
}
