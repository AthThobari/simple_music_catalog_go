package spotify

var searchResponse = `
{
  "tracks": {
    "href": "https://api.spotify.com/v1/search?offset=0&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9,id;q%3D0.8",
    "limit": 10,
    "next": "https://api.spotify.com/v1/search?offset=10&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9,id;q%3D0.8",
    "offset": 0,
    "previous": null,
    "total": 906,
    "items": [
      {
        "album": {
          "album_type": "album",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
              },
              "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
              "id": "1dfeR4HaWDbWqFHLkxsg1d",
              "name": "Queen",
              "type": "artist",
              "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/6i6folBtxKV28WX3msQ4FE"
          },
          "href": "https://api.spotify.com/v1/albums/6i6folBtxKV28WX3msQ4FE",
          "id": "6i6folBtxKV28WX3msQ4FE",
          "images": [
            {
              "height": 640,
              "width": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b"
            },
            {
              "height": 300,
              "width": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b"
            },
            {
              "height": 64,
              "width": 64,
              "url": "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"
            }
          ],
          "is_playable": true,
          "name": "Bohemian Rhapsody (The Original Soundtrack)",
          "release_date": "2018-10-19",
          "release_date_precision": "day",
          "total_tracks": 22,
          "type": "album",
          "uri": "spotify:album:6i6folBtxKV28WX3msQ4FE"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
            },
            "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
            "id": "1dfeR4HaWDbWqFHLkxsg1d",
            "name": "Queen",
            "type": "artist",
            "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
          }
        ],
        "disc_number": 1,
        "duration_ms": 354947,
        "explicit": false,
        "external_ids": {
          "isrc": "GBUM71029604"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/3z8h0TU7ReDPLIbEnYhWZb"
        },
        "href": "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
        "id": "3z8h0TU7ReDPLIbEnYhWZb",
        "is_local": false,
        "is_playable": true,
        "name": "Bohemian Rhapsody",
        "popularity": 71,
        "preview_url": null,
        "track_number": 7,
        "type": "track",
        "uri": "spotify:track:3z8h0TU7ReDPLIbEnYhWZb"
      },
      {
        "album": {
          "album_type": "compilation",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
              },
              "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
              "id": "1dfeR4HaWDbWqFHLkxsg1d",
              "name": "Queen",
              "type": "artist",
              "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/3VWrUk4vBznMYXGMPc7dRB"
          },
          "href": "https://.spotify.com/v1/albums/3VWrUk4vBznMYXGMPc7dRB",
          "id": "3VWrUk4vBznMYXGMPc7dRB",
          "images": [
            {
              "height": 640,
              "width": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b273bb19d0c22d5709c9d73c8263"
            },
            {
              "height": 300,
              "width": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e02bb19d0c22d5709c9d73c8263"
            },
            {
              "height": 64,
              "width": 64,
              "url": "https://i.scdn.co/image/ab67616d00004851bb19d0c22d5709c9d73c8263"
            }
          ],
          "is_playable": true,
          "name": "Greatest Hits (Remastered)",
          "release_date": "1981-10-26",
          "release_date_precision": "day",
          "total_tracks": 17,
          "type": "album",
          "uri": "spotify:album:3VWrUk4vBznMYXGMPc7dRB"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
            },
            "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
            "id": "1dfeR4HaWDbWqFHLkxsg1d",
            "name": "Queen",
            "type": "artist",
            "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
          }
        ],
        "disc_number": 1,
        "duration_ms": 355466,
        "explicit": false,
        "external_ids": {
          "isrc": "GBUM71029604"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/2OBofMJx94NryV2SK8p8Zf"
        },
        "href": "https://api.spotify.com/v1/tracks/2OBofMJx94NryV2SK8p8Zf",
        "id": "2OBofMJx94NryV2SK8p8Zf",
        "is_local": false,
        "is_playable": true,
        "name": "Bohemian Rhapsody - Remastered 2011",
        "popularity": 52,
        "preview_url": null,
        "track_number": 1,
        "type": "track",
        "uri": "spotify:track:2OBofMJx94NryV2SK8p8Zf"
      }
    ]
  }
}
`
var recommendationResponse = `
{
  "tracks": {
    "href": "https://api.spotify.com/v1/search?offset=0&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9,id;q%3D0.8",
    "limit": 10,
    "items": [
      {
        "album": {
          "album_type": "album",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
              },
              "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
              "id": "1dfeR4HaWDbWqFHLkxsg1d",
              "name": "Queen",
              "type": "artist",
              "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/6i6folBtxKV28WX3msQ4FE"
          },
          "href": "https://api.spotify.com/v1/albums/6i6folBtxKV28WX3msQ4FE",
          "id": "6i6folBtxKV28WX3msQ4FE",
          "images": [
            {
              "height": 640,
              "width": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b"
            },
            {
              "height": 300,
              "width": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b"
            },
            {
              "height": 64,
              "width": 64,
              "url": "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"
            }
          ],
          "is_playable": true,
          "name": "Bohemian Rhapsody (The Original Soundtrack)",
          "release_date": "2018-10-19",
          "release_date_precision": "day",
          "total_tracks": 22,
          "type": "album",
          "uri": "spotify:album:6i6folBtxKV28WX3msQ4FE"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
            },
            "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
            "id": "1dfeR4HaWDbWqFHLkxsg1d",
            "name": "Queen",
            "type": "artist",
            "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
          }
        ],
        "disc_number": 1,
        "duration_ms": 354947,
        "explicit": false,
        "external_ids": {
          "isrc": "GBUM71029604"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/3z8h0TU7ReDPLIbEnYhWZb"
        },
        "href": "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
        "id": "3z8h0TU7ReDPLIbEnYhWZb",
        "is_local": false,
        "is_playable": true,
        "name": "Bohemian Rhapsody",
        "popularity": 71,
        "preview_url": null,
        "track_number": 7,
        "type": "track",
        "uri": "spotify:track:3z8h0TU7ReDPLIbEnYhWZb"
      },
      {
        "album": {
          "album_type": "compilation",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
              },
              "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
              "id": "1dfeR4HaWDbWqFHLkxsg1d",
              "name": "Queen",
              "type": "artist",
              "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/3VWrUk4vBznMYXGMPc7dRB"
          },
          "href": "https://.spotify.com/v1/albums/3VWrUk4vBznMYXGMPc7dRB",
          "id": "3VWrUk4vBznMYXGMPc7dRB",
          "images": [
            {
              "height": 640,
              "width": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b273bb19d0c22d5709c9d73c8263"
            },
            {
              "height": 300,
              "width": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e02bb19d0c22d5709c9d73c8263"
            },
            {
              "height": 64,
              "width": 64,
              "url": "https://i.scdn.co/image/ab67616d00004851bb19d0c22d5709c9d73c8263"
            }
          ],
          "is_playable": true,
          "name": "Greatest Hits (Remastered)",
          "release_date": "1981-10-26",
          "release_date_precision": "day",
          "total_tracks": 17,
          "type": "album",
          "uri": "spotify:album:3VWrUk4vBznMYXGMPc7dRB"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/1dfeR4HaWDbWqFHLkxsg1d"
            },
            "href": "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
            "id": "1dfeR4HaWDbWqFHLkxsg1d",
            "name": "Queen",
            "type": "artist",
            "uri": "spotify:artist:1dfeR4HaWDbWqFHLkxsg1d"
          }
        ],
        "disc_number": 1,
        "duration_ms": 355466,
        "explicit": false,
        "external_ids": {
          "isrc": "GBUM71029604"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/2OBofMJx94NryV2SK8p8Zf"
        },
        "href": "https://api.spotify.com/v1/tracks/2OBofMJx94NryV2SK8p8Zf",
        "id": "2OBofMJx94NryV2SK8p8Zf",
        "is_local": false,
        "is_playable": true,
        "name": "Bohemian Rhapsody - Remastered 2011",
        "popularity": 52,
        "preview_url": null,
        "track_number": 1,
        "type": "track",
        "uri": "spotify:track:2OBofMJx94NryV2SK8p8Zf"
      }
    ]
  }
}
`
