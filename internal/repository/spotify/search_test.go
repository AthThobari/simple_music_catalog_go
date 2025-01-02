package spotify

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/AthThobari/simple_music_catalog_go/internal/configs"
	"github.com/AthThobari/simple_music_catalog_go/pkg/httpclient"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func Test_outbound_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockHTTPClient := httpclient.NewMockHTTPClient(mockCtrl)
	next := "https://api.spotify.com/v1/search?offset=10&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9,id;q%3D0.8"
	type args struct {
		query  string
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    *SpotifySearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:  "bohemian rhapsody",
				limit:  10,
				offset: 0,
			},
			want: &SpotifySearchResponse{
				Track: SpotifyTracks{
					HREF:   "https://api.spotify.com/v1/search?offset=0&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9,id;q%3D0.8",
					Limit:  10,
					Next:   &next,
					Offset: 0,
					Items: []SpotifyTracksObject{
						{
							Album: SpotifyAlbumObject{
								AlbumType:   "album",
								TotalTracks: 22,
								Images: []SpotifyAlbumImage{
									{
										URL: "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
									},
									{
										URL: "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
									},
									{
										URL: "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
									},
								},
								Name: "Bohemian Rhapsody (The Original Soundtrack)",
							},
							Artists: []SpotifyArtistObject{
								{
									HREF: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
									Name: "Queen",
								},
							},
							Explicit: false,
							HREF:     "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
							ID:       "3z8h0TU7ReDPLIbEnYhWZb",
							Name:     "Bohemian Rhapsody",
						},
						{
							Album: SpotifyAlbumObject{
								AlbumType:   "compilation",
								TotalTracks: 17,
								Images: []SpotifyAlbumImage{
									{
										URL: "https://i.scdn.co/image/ab67616d0000b273bb19d0c22d5709c9d73c8263",
									},
									{
										URL: "https://i.scdn.co/image/ab67616d00001e02bb19d0c22d5709c9d73c8263",
									},
									{
										URL: "https://i.scdn.co/image/ab67616d00004851bb19d0c22d5709c9d73c8263",
									},
								},
								Name: "Greatest Hits (Remastered)",
							},
							Artists: []SpotifyArtistObject{
								{
									HREF: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
									Name: "Queen",
								},
							},
							Explicit: false,
							HREF:     "https://api.spotify.com/v1/tracks/2OBofMJx94NryV2SK8p8Zf",
							ID:       "2OBofMJx94NryV2SK8p8Zf",
							Name: "Bohemian Rhapsody - Remastered 2011",
						},
					},
					Total: 906,
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				basePath := `https://api.spotify.com/v1/search`
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())

				fmt.Println("Mock URL path:", urlPath)

				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessToken")
				fmt.Println("Mock authorization Header:", req.Header.Get("Authorization"))

				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(searchResponse)),
				}, nil)
			},
		},
		{
			name: "fail",
			args: args{
				query:  "bohemian rhapsody",
				limit:  10,
				offset: 0,
			},
			want: nil,
			wantErr: true,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				basePath := `https://api.spotify.com/v1/search`
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())

				fmt.Println("Mock URL path:", urlPath)

				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessToken")
				fmt.Println("Mock authorization Header:", req.Header.Get("Authorization"))

				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 500,
					Body:       io.NopCloser(bytes.NewBufferString("internal server error")),
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			o := &outbound{
				cfg:         &configs.Config{},
				client:      mockHTTPClient,
				AccessToken: "accessToken",
				TokenType:   "Bearer",
				ExpiredAt:   time.Now().Add(1 * time.Hour),
			}
			got, err := o.Search(context.Background(), tt.args.query, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("outbound.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
