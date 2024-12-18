package tracks

import (
	"context"

	"github.com/AthThobari/simple_music_catalog_go/internal/repository/spotify"
)


type spotifyOutboound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}

type service struct {
	spotifyOutboound spotifyOutboound
}

func NewService(spotifyOutboound spotifyOutboound) *service {
	return &service{spotifyOutboound: spotifyOutboound}
}
