# immutablex-cli

> **_NOTE:_**  This app is a work in progress. It was previously a BitVerse NFT browser but it's being converted into a general ImmutableX command-line client.

## Installation

Prerequisites:

* Go [go.dev](https://go.dev/)
* An Alchemy API Key from [alchemy.com](https://alchemy.com) (it's free)

First add your Alchemy API key to your environment as `ALCHEMY_API_KEY`.

Next you can either `go install` it, or clone the repo and build it:

```txt
% go install github.com/deadloct/immutablex-cli@latest

# To remove it later:
% go clean -n -i github.com/deadloct/immutablex-cli...
```

Clone and build method:

```txt
% git clone git@github.com:deadloct/immutablex-cli.git
% cd immutablex-cli
% go build
```

## Usage

The app currently has two commands: `asset` and `assets`.

### Asset

Queries the ImmutableX getAsset endpoint for detailed asset information, see [https://docs.x.immutable.com/reference/#/operations/getAsset](https://docs.x.immutable.com/reference/#/operations/getAsset).

```txt
Usage:
  immutablex-cli asset [flags]

Flags:
  -h, --help                   help for asset
  -f, --include-fees           include fees associated with the asset
  -a, --token-address string   address of the collection or shortcut
  -i, --token-id string        id of the asset

Global Flags:
  -v, --verbose   Verbose output
```

Example:

```txt
% immutablex-cli asset --token-address 0x6465ef3009f3c474774f4afb607a5d600ea71d95 --token-id 2578
2022/12/24 06:32:14 requesting asset 2578 from collection 0x6465ef3009f3c474774f4afb607a5d600ea71d95

BitHero #2578:
- Background: Peekaboo Plant
- Eyes: Emerald
- Frame: Mythic
- Gender: Female
- Generation: 0
- Hair: Aquamarine Ponytail
- Hat: Tasseled Wizard
- Outfit: Edo Yakuza
- Rarity: Mythic
- Skin: Color Tone 6
- Description: BitHeroes from the Bitverse
- Image URL: https://d3n9vm398ay3ts.cloudfront.net/heroes-2022/0/images/hero-2578.gif
- Game Meta JSON: https://d3n9vm398ay3ts.cloudfront.net/heroes-2022/0/game_meta/hero-2578.json
- Owner: https://immutascan.io/address/0x1f67800e5aee081b53b7c0f5ac5d33f23e6d1252
```

### Assets

Queries the ImmutableX listAssets API for detailed asset information, see [https://docs.x.immutable.com/reference/#/operations/listAssets](https://docs.x.immutable.com/reference/#/operations/listAssets).

```txt
Usage:
  immutablex-cli assets [flags]

Flags:
  -b, --buy-orders                     Retrieve buy orders for each asset
  -c, --collection string              Address of the collection or shortcut
  -d, --direction string               asc|desc
  -h, --help                           help for assets
  -i, --include-fees                   Retrieves fees for each asset
  -m, --metadata stringArray           Filter by metadata in key=value format (repeatable). For example "immutable-cli assets -m Rarity=Mythic -m Generation=0. Note that metadata keys and values are case sensitive.
  -n, --name string                    Search for this asset name (default "desc")
  -o, --order-by string                updated_at|name (default "updated_at")
  -l, --sell-orders                    Retrieves sell orders for each asset
  -s, --status string                  Filter by the status: eth|imx|preparing_withdrawal|withdrawable|burned
  -x, --updated-max-timestamp string   Include results on or before this time in ISO 8601 UTC format
  -z, --updated-min-timestamp string   Include results on or after this time in ISO 8601 UTC format
  -u, --user string                    Retrieves assets owned by this user/wallet address

Global Flags:
  -v, --verbose   Verbose output
```

Example:

```txt
% immutablex-cli assets --collection 0xe4ac52f4b4a721d1d0ad8c9c689df401c2db7291 --updated-min-timestamp=2022-12-23T00:00:00Z --metadata Generation=0 -v
2022/12/24 13:19:12 fetched 2 assets from 2022-12-23T08:18:19.07428Z to 2022-12-24T13:13:50.647132Z
Portal #969: burned (https:/immutascan.io/address/0xe4ac52f4b4a721d1d0ad8c9c689df401c2db7291/969)
Portal #1439: burned (https:/immutascan.io/address/0xe4ac52f4b4a721d1d0ad8c9c689df401c2db7291/1439)

Asset counts for collection 0xe4ac52f4b4a721d1d0ad8c9c689df401c2db7291:
- Common: 1
- Rare: 1
- Epic: 0
- Legendary: 0
- Mythic: 0
- Total: 2
```

## Shortcuts

Remembering collection addresses is tedious. If you'd rather use a shortname for a commonly used collection, copy the json data at the top of `lib/collection_manager.go` to some file on your computer, and then set an environment variable `IMX_SHORTCUT_LOCATION` for the full path to that file. After that you can use the shortcut in commands instead of the collection address.

For example, retrieving the specific NFT above with the shortcut `hero`:

```txt
% immutablex-cli asset -a hero -i 2578
2022/12/24 06:32:14 requesting asset 2578 from collection 0x6465ef3009f3c474774f4afb607a5d600ea71d95

BitHero #2578:
- Background: Peekaboo Plant
- Eyes: Emerald
- Frame: Mythic
- Gender: Female
- Generation: 0
- Hair: Aquamarine Ponytail
- Hat: Tasseled Wizard
- Outfit: Edo Yakuza
- Rarity: Mythic
- Skin: Color Tone 6
- Description: BitHeroes from the Bitverse
- Image URL: https://d3n9vm398ay3ts.cloudfront.net/heroes-2022/0/images/hero-2578.gif
- Game Meta JSON: https://d3n9vm398ay3ts.cloudfront.net/heroes-2022/0/game_meta/hero-2578.json
- Owner: https://immutascan.io/address/0x1f67800e5aee081b53b7c0f5ac5d33f23e6d1252
```
