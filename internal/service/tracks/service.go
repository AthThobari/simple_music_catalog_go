package tracks

import (
	"context"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/trackactivities"
	"github.com/AthThobari/simple_music_catalog_go/internal/repository/spotify"
)

type spotifyOutboound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
	GetRecommendation(ctx context.Context, limit int, trackID string) (*spotify.SpotifySearchResponse, error)
}

type trackActivitiesRepository interface {
	Create(ctx context.Context, model trackactivities.TrackActivity) error
	Update(ctx context.Context, model trackactivities.TrackActivity) error
	Get(ctx context.Context, userID uint, spotifyID string) (*trackactivities.TrackActivity, error)
	GetBulkSpotifyIDs(ctx context.Context, userID uint, spotifyIDs []string) (map[string]trackactivities.TrackActivity, error)
}

type service struct {
	spotifyOutboound    spotifyOutboound
	trackActivitiesRepo trackActivitiesRepository
}

func NewService(spotifyOutboound spotifyOutboound, trackActivitiesRepo trackActivitiesRepository) *service {
	return &service{spotifyOutboound: spotifyOutboound, trackActivitiesRepo: trackActivitiesRepo}
}
