package tracks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/spotify"
	"github.com/AthThobari/simple_music_catalog_go/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetRecommendation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSvc := NewMockservice(mockCtrl)

	tests := []struct {
		name   string
		expectedStatusCode int
		expectedBody *spotify.RecommendationResponse
		wantErr bool
		mockFn func()
	}{
		{
			name: "success",
			expectedStatusCode: http.StatusOK,
			expectedBody: &spotify.RecommendationResponse{
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
					},
				},
			},
			wantErr: false,
			mockFn: func() {
				mockSvc.EXPECT().GetRecommendation(gomock.Any(), uint(1), 10, "trackID").Return(&spotify.RecommendationResponse{
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
						},
					},
				}, nil)
			},
		},
		{
			name: "failed",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody: nil,
			wantErr: true,
			mockFn: func() {
				mockSvc.EXPECT().GetRecommendation(gomock.Any(), uint(1), 10, "trackID").Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn()
		t.Run(tt.name, func(t *testing.T) {
			api := gin.New()
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()

			endpoint := `/tracks/recommendations?limit=10&trackID=trackID`

			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)
			token, err := jwt.CreateToken(1, "username", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := spotify.RecommendationResponse{}
				_ = json.Unmarshal(w.Body.Bytes(), &response)

				assert.Equal(t, tt.expectedBody, &response)
			}
		})
	}
}
