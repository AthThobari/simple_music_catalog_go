package tracks

import (
	"context"
	"reflect"
	"testing"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/spotify"
	"github.com/AthThobari/simple_music_catalog_go/internal/models/trackactivities"
	spotifyRepo "github.com/AthThobari/simple_music_catalog_go/internal/repository/spotify"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_service_GetRecommendation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCtrl.Finish()

	mockspotifyOutbound := NewMockspotifyOutboound(mockCtrl)
	mockTrackActivitiesRepo := NewMocktrackActivitiesRepository(mockCtrl)
	isLikedTrue := true
	isLikedFalse := false

	type args struct {
		userID  uint
		limit   int
		trackID string
	}
	tests := []struct {
		name    string
		args    args
		want    *spotify.RecommendationResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				userID:  1,
				limit:   10,
				trackID: "trackID",
			},
			want: &spotify.RecommendationResponse{
				Items: []spotify.SpotifyTracksObject{
					{
						AlbumType:        "album",
						AlbumTotalTracks: 22,
						AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"},
						AlbumName:        "Bohemian Rhapsody (The Original Soundtrack)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						ID:               "3z8h0TU7ReDPLIbEnYhWZb",
						Name:             "Bohemian Rhapsody",
						IsLiked:          &isLikedTrue,
					},
					{
						AlbumType:        "compilation",
						AlbumTotalTracks: 17,
						AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b273bb19d0c22d5709c9d73c8263", "https://i.scdn.co/image/ab67616d00001e02bb19d0c22d5709c9d73c8263", "https://i.scdn.co/image/ab67616d00004851bb19d0c22d5709c9d73c8263"},
						AlbumName:        "Greatest Hits (Remastered)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						ID:               "2OBofMJx94NryV2SK8p8Zf",
						Name:             "Bohemian Rhapsody - Remastered 2011",
						IsLiked:          &isLikedFalse,
					},
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockspotifyOutbound.EXPECT().GetRecommendation(gomock.Any(), 10, "trackID").Return(&spotifyRepo.SpotifySearchResponse{
					Track: spotifyRepo.SpotifyTracks{
						HREF:  "https://api.spotify.com/v1/search?offset=0&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9,id;q%3D0.8",
						Limit: 10,
						Items: []spotifyRepo.SpotifyTracksObject{
							{
								Album: spotifyRepo.SpotifyAlbumObject{
									AlbumType:   "album",
									TotalTracks: 22,
									Images: []spotifyRepo.SpotifyAlbumImage{
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
								Artists: []spotifyRepo.SpotifyArtistObject{
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
								Album: spotifyRepo.SpotifyAlbumObject{
									AlbumType:   "compilation",
									TotalTracks: 17,
									Images: []spotifyRepo.SpotifyAlbumImage{
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
								Artists: []spotifyRepo.SpotifyArtistObject{
									{
										HREF: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
										Name: "Queen",
									},
								},
								Explicit: false,
								HREF:     "https://api.spotify.com/v1/tracks/2OBofMJx94NryV2SK8p8Zf",
								ID:       "2OBofMJx94NryV2SK8p8Zf",
								Name:     "Bohemian Rhapsody - Remastered 2011",
							},
						},
					},
				}, nil)

				mockTrackActivitiesRepo.EXPECT().GetBulkSpotifyIDs(gomock.Any(), uint(1), []string{"3z8h0TU7ReDPLIbEnYhWZb", "2OBofMJx94NryV2SK8p8Zf"}).Return(map[string]trackactivities.TrackActivity{
					"3z8h0TU7ReDPLIbEnYhWZb": {
						IsLiked: &isLikedTrue,
					},
					"2OBofMJx94NryV2SK8p8Zf": {
						IsLiked: &isLikedFalse,
					},
				}, nil)
			},
		},
		{
			name: "failed: when get bulk spotify id",
			args: args{
				userID:  1,
				limit:   10,
				trackID: "trackID",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockspotifyOutbound.EXPECT().GetRecommendation(gomock.Any(), 10, "trackID").Return(&spotifyRepo.SpotifySearchResponse{
					Track: spotifyRepo.SpotifyTracks{
						HREF:  "https://api.spotify.com/v1/search?offset=0&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9,id;q%3D0.8",
						Limit: 10,
						Items: []spotifyRepo.SpotifyTracksObject{
							{
								Album: spotifyRepo.SpotifyAlbumObject{
									AlbumType:   "album",
									TotalTracks: 22,
									Images: []spotifyRepo.SpotifyAlbumImage{
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
								Artists: []spotifyRepo.SpotifyArtistObject{
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
								Album: spotifyRepo.SpotifyAlbumObject{
									AlbumType:   "compilation",
									TotalTracks: 17,
									Images: []spotifyRepo.SpotifyAlbumImage{
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
								Artists: []spotifyRepo.SpotifyArtistObject{
									{
										HREF: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
										Name: "Queen",
									},
								},
								Explicit: false,
								HREF:     "https://api.spotify.com/v1/tracks/2OBofMJx94NryV2SK8p8Zf",
								ID:       "2OBofMJx94NryV2SK8p8Zf",
								Name:     "Bohemian Rhapsody - Remastered 2011",
							},
						},
					},
				}, nil)

				mockTrackActivitiesRepo.EXPECT().GetBulkSpotifyIDs(gomock.Any(), uint(1), []string{"3z8h0TU7ReDPLIbEnYhWZb", "2OBofMJx94NryV2SK8p8Zf"}).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed: when get recommendation from spotify outbound",
			args: args{
				userID:  1,
				limit:   10,
				trackID: "trackID",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockspotifyOutbound.EXPECT().GetRecommendation(gomock.Any(), 10, "trackID").Return(nil, assert.AnError)

			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				spotifyOutboound:    mockspotifyOutbound,
				trackActivitiesRepo: mockTrackActivitiesRepo,
			}
			got, err := s.GetRecommendation(context.Background(), tt.args.userID, tt.args.limit, tt.args.trackID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetRecommendation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetRecommendation() = %v, want %v", got, tt.want)
			}
		})
	}
}
