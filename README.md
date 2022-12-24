# immutablex-cli

> **_NOTE:_**  This app is a work in progress. It was previously a BitVerse NFT browser but it's being converted into a general ImmutableX command-line client.

## Installation

Prerequisites:

* Go [go.dev](https://go.dev/)
* An Alchemy API Key from [alchemy.com](https://alchemy.com) (it's free)

First add your Alchemy API key to your environment as `ALCHEMY_API_KEY`.

Next clone the repo and build it:

```bash
% git clone git@github.com:deadloct/immutablex-cli.git
% cd immutablex-cli
% go build
```

## Usage

The app currently has two commands: `asset` and `assets`.

### Asset

```bash
Usage:
  immutablex-cli asset [flags]

Flags:
  -a, --addr string   address of the collection or shortcut
  -h, --help          help for asset
  -i, --id string     id of the asset

Global Flags:
  -v, --verbose   Verbose output
```

Asset will retrive the given asset. The `id` and `addr` fields are required.

For example, to retrieve one of the mythic BitVerse Heroes NFTs:

```bash
% ./immutablex-cli asset -a 0x6465ef3009f3c474774f4afb607a5d600ea71d95 -i 2578
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

```bash
Usage:
  immutablex-cli assets [flags]

Flags:
  -a, --addr string     Address of the collection or shortcut
  -h, --help            help for assets
  -o, --owner string    Filter by owner
  -r, --rarity string   Filter by rarity
  -s, --status string   Filter by status

Global Flags:
  -v, --verbose   Verbose output
```

For example, to retrieve all mythic NFTs with their URLs (`-v`):

```bash
% ./immutablex-cli assets -a 0x6465ef3009f3c474774f4afb607a5d600ea71d95 -r mythic -v
2022/12/24 06:33:03 fetched 200 assets from 2022-12-24T13:18:02.817906Z to 2022-11-30T15:29:52.931155Z
2022/12/24 06:33:03 fetched 200 assets from 2022-11-30T15:02:10.663512Z to 2022-11-03T20:54:10.896841Z
2022/12/24 06:33:04 fetched 200 assets from 2022-11-03T20:51:52.221243Z to 2022-09-27T13:57:14.24508Z
2022/12/24 06:33:04 fetched 172 assets from 2022-09-27T13:57:14.24508Z to 2022-09-21T19:46:01.938354Z
BitHero #7120: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/7120)
BitHero #2546: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2546)
BitHero #2577: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2577)
BitHero #2566: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2566)
BitHero #2588: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2588)
BitHero #2594: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2594)
BitHero #2538: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2538)
BitHero #2571: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2571)
BitHero #2569: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2569)
BitHero #2545: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2545)
BitHero #2560: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2560)
BitHero #2567: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2567)
BitHero #2592: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2592)
BitHero #2581: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2581)
BitHero #2583: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2583)
BitHero #2585: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2585)
BitHero #2578: imx (https:/immutascan.io/address/0x6465ef3009f3c474774f4afb607a5d600ea71d95/2578)

Asset counts for collection 0x6465ef3009f3c474774f4afb607a5d600ea71d95:
- Common: 0
- Rare: 0
- Epic: 0
- Legendary: 0
- Mythic: 17
- Total: 17 
```

## Shortcuts

Remembering collection addresses is tedious. If you'd rather use a shortname for a commonly used collection, copy the json data at the top of `lib/collection_manager.go` to some file on your computer, and then set an environment variable `IMX_SHORTCUT_LOCATION` for the full path to that file. After that you can use the shortcut in commands instead of the collection address.

For example, retrieving the specific NFT above with the shortcut `hero`:

```bash
% ./immutablex-cli asset -a hero -i 2578
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
