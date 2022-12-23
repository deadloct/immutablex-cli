# bh-imx-browser

This soon to be rebranded go app connects to Immutable using their API and pulls
information about the BitVerse suite of NFTs by Kongregate.

It's still being actively developed (consider it alpha), but for now, here's how to use it:

## Installation 

First clone the repo and build it:

```bash
% git clone git@github.com:deadloct/bh-imx-browser.git
% cd bh-imx-browser
% go build
```

## Usage

The app currently has two commands: `asset` and `assets`.

### Asset 

```bash
Usage:
  bh-imx-browser asset [flags]

Flags:
  -h, --help          help for asset
  -i, --id string
  -t, --type string   Type (default "hero")

Global Flags:
  -v, --verbose   Verbose output
```

Asset will retrive the given asset. The `id` field is required.

For example, to retrieve one of the mythic heroes:

```bash
% ./bh-imx-browser asset --id 2578

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
  bh-imx-browser assets [flags]

Flags:
  -h, --help            help for assets
  -o, --owner string    Filter by owner
  -r, --rarity string   Filter by rarity
  -s, --status string   Filter by status

Global Flags:
  -v, --verbose   Verbose output
```

For example, to retrieve all mythic NFTs:

```bash
% ./bh-imx-browser assets --rarity mythic -v
2022/12/23 11:26:12 fetched 200 assets from 2022-12-23T19:07:44.478296Z to 2022-11-04T16:29:09.447629Z
2022/12/23 11:26:13 fetched 200 assets from 2022-11-04T16:28:38.023997Z to 2022-10-05T16:52:44.639866Z
2022/12/23 11:26:13 fetched 200 assets from 2022-10-05T16:52:44.631988Z to 2022-09-27T04:19:16.870808Z
2022/12/23 11:26:13 fetched 200 assets from 2022-09-27T04:19:00.166563Z to 2022-09-26T23:43:39.238249Z
2022/12/23 11:26:14 fetched 200 assets from 2022-09-26T23:43:10.482575Z to 2022-09-23T17:15:50.669121Z
2022/12/23 11:26:14 fetched 144 assets from 2022-09-23T17:15:50.628465Z to 2022-09-21T19:45:56.579627Z

BitVerse Portals:
- Common: 0
- Rare: 0
- Epic: 0
- Legendary: 0
- Mythic: 0
- Total: 0

2022/12/23 11:26:15 fetched 200 assets from 2022-12-23T19:08:26.253902Z to 2022-11-30T00:38:32.000008Z
2022/12/23 11:26:15 fetched 200 assets from 2022-11-30T00:38:12.751911Z to 2022-11-03T00:19:29.164616Z
2022/12/23 11:26:15 fetched 200 assets from 2022-11-02T22:42:07.820754Z to 2022-09-27T13:57:14.24508Z
2022/12/23 11:26:16 fetched 166 assets from 2022-09-27T13:57:14.24508Z to 2022-09-21T19:46:01.938354Z
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

BitVerse Heroes:
- Common: 0
- Rare: 0
- Epic: 0
- Legendary: 0
- Mythic: 17
- Total: 17
```
