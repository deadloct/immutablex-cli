package main

import (
	"context"
	"log"
	"os"
)

const AddrPortals = "0xe4ac52f4b4a721d1d0ad8c9c689df401c2db7291"
const AddrHeroes = "0x6465ef3009f3c474774f4afb607a5d600ea71d95"

var AlchemyKey string

func init() {
	AlchemyKey = os.Getenv("ALCHEMY_API_KEY")
	if AlchemyKey == "" {
		log.Panic("no alchemy api key provided, get one at alchemy.com")
	}
}

func main() {
	c, err := NewClient(AlchemyKey)
	if err != nil {
		log.Panic(err)
	}

	defer c.Stop()

	c.PrintAssetCounts(context.Background(), AddrPortals)
	c.PrintAssetCounts(context.Background(), AddrHeroes)
}
