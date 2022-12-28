# immutablex-cli

> **_NOTE:_**  This is an alpha stage tool that is currently being developed. It only supports the following endpoints at the moment:
>  
> * list-assets
> * get-asset
> * list-collections
> * get-collection
> * list-orders
>  
> All read-only endpoints will be added first followed by write operations afterward.

## Installation

Follow the instructions on [go.dev](https://go.dev/) to install the latest version of Go.

Option 1: `go install` method:

```txt
% go install github.com/deadloct/immutablex-cli@latest

# To remove it later:
% go clean -i github.com/deadloct/immutablex-cli...
```

Option 2: Clone and build method:

```txt
% git clone git@github.com:deadloct/immutablex-cli.git
% cd immutablex-cli
% go build
```

## Examples

For complete usage information, please type `immutablex-cli -h` or `immutablex-cli [subcommand] -h`.

### Get Asset

```txt
% immutablex-cli get-asset --token-address 0x6465ef3009f3c474774f4afb607a5d600ea71d95 --token-id 2578
{
  "collection": {
    "icon_url": "https://thebitverse.io/nft-assets/heroes_cover_image.png",
    "name": "Bitverse Heroes"
  },
  "created_at": "2022-09-27T01:10:55.829259Z",
  "description": "BitHeroes from the Bitverse",
  "id": "0x5bc0692c1b2276f7812438aa32bd0901e4a6915b059ce00061c43cda3458b2b5",
  "image_url": "https://d3n9vm398ay3ts.cloudfront.net/heroes-2022/0/images/hero-2578.gif",
  "metadata": {
    "Background": "Peekaboo Plant",
    "Eye": "Emerald",
    "Frame": "Mythic",
    "Gender": "Female",
    "Generation": 0,
    "Hair": "Aquamarine Ponytail",
    "Halo": "Bad Day",
    "Hat": "Tasseled Wizard",
    "Horn": "Mythic Stabbers",
    "Mask": "Mustache Shield",
    "Outfit": "Edo Yakuza",
    "Rarity": "Mythic",
    "Skin": "Color Tone 6",
    "description": "BitHeroes from the Bitverse",
    "game_meta": "https://d3n9vm398ay3ts.cloudfront.net/heroes-2022/0/game_meta/hero-2578.json",
    "image": "https://d3n9vm398ay3ts.cloudfront.net/heroes-2022/0/images/hero-2578.gif",
    "name": "BitHero #2578"
  },
  "name": "BitHero #2578",
  "status": "imx",
  "token_address": "0x6465ef3009f3c474774f4afb607a5d600ea71d95",
  "token_id": "2578",
  "updated_at": "2022-09-28T23:28:03.015304Z",
  "uri": null,
  "user": "0x1f67800e5aee081b53b7c0f5ac5d33f23e6d1252"
}
```

### List Assets

```txt
% immutablex-cli list-assets --collection 0xe4ac52f4b4a721d1d0ad8c9c689df401c2db7291 --updated-min-timestamp=2022-12-23T00:00:00Z --metadata Generation=0
Portal #969 (Status: burned): (https:/immutascan.io/address/0xe4ac52f4b4a721d1d0ad8c9c689df401c2db7291/969)
Portal #1439 (Status: burned): (https:/immutascan.io/address/0xe4ac52f4b4a721d1d0ad8c9c689df401c2db7291/1439)
2 total assets returned
```

### Get Collection

```txt
% immutablex-cli get-collection -c 0x6465ef3009f3c474774f4afb607a5d600ea71d95
{
  "address": "0x6465ef3009f3c474774f4afb607a5d600ea71d95",
  "collection_image_url": "https://thebitverse.io/nft-assets/heroes_cover_image.png",
  "created_at": "2022-09-17T00:50:57.31643Z",
  "description": "Own your progress. Own your brand. Access exclusive content. Bitverse Heroes are your key to the exciting Bitverse.",
  "icon_url": "https://thebitverse.io/nft-assets/heroes_cover_image.png",
  "metadata_api_url": "https://thebitverse.io/api/heroes/metadata/imx",
  "name": "Bitverse Heroes",
  "project_id": 10014,
  "project_owner_address": "0x771642c8ad544b48308f5e3a49d73da94d62be3f",
  "updated_at": "2022-09-17T00:50:57.31643Z"
}
```

### List Collections

```txt
% immutablex-cli list-collections --blacklist 0x6465ef3009f3c474774f4afb607a5d600ea71d95 --keyword bitverse
Bitverse Portals: https://immutascan.io/address/0xe4ac52f4b4a721d1d0ad8c9c689df401c2db7291
1 total collections returned
```

## Shortcuts

Remembering collection addresses is tedious. If you'd rather use a shortname for a commonly used collection, copy the json data at the top of `lib/collection_manager.go` to some file on your computer, and then set an environment variable `IMX_SHORTCUT_LOCATION` for the full path to that file. After that you can use the shortcut in commands instead of the collection address.

For example, retrieving the specific NFT above with the shortcut `hero`:

```txt
% immutablex-cli get-asset --token-address hero --token-id 2578
{
  "collection": {
    "icon_url": "https://thebitverse.io/nft-assets/heroes_cover_image.png",
    "name": "Bitverse Heroes"
  },
  "created_at": "2022-09-27T01:10:55.829259Z",
  "description": "BitHeroes from the Bitverse",
  "id": "0x5bc0692c1b2276f7812438aa32bd0901e4a6915b059ce00061c43cda3458b2b5",
  "image_url": "https://d3n9vm398ay3ts.cloudfront.net/heroes-2022/0/images/hero-2578.gif",
  "metadata": {
    "Background": "Peekaboo Plant",
    "Eye": "Emerald",
    "Frame": "Mythic",
    "Gender": "Female",
    "Generation": 0,
    "Hair": "Aquamarine Ponytail",
    "Halo": "Bad Day",
    "Hat": "Tasseled Wizard",
    "Horn": "Mythic Stabbers",
    "Mask": "Mustache Shield",
    "Outfit": "Edo Yakuza",
    "Rarity": "Mythic",
    "Skin": "Color Tone 6",
    "description": "BitHeroes from the Bitverse",
    "game_meta": "https://d3n9vm398ay3ts.cloudfront.net/heroes-2022/0/game_meta/hero-2578.json",
    "image": "https://d3n9vm398ay3ts.cloudfront.net/heroes-2022/0/images/hero-2578.gif",
    "name": "BitHero #2578"
  },
  "name": "BitHero #2578",
  "status": "imx",
  "token_address": "0x6465ef3009f3c474774f4afb607a5d600ea71d95",
  "token_id": "2578",
  "updated_at": "2022-09-28T23:28:03.015304Z",
  "uri": null,
  "user": "0x1f67800e5aee081b53b7c0f5ac5d33f23e6d1252"
}
```
