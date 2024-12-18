package tracks

import (
	"context"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/spotify"
	spotifyRepo "github.com/AthThobari/simple_music_catalog_go/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.spotifyOutboound.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error search track to spotify")
		return nil, err
	}

	return modelToResponse(trackDetails), nil
}

func modelToResponse(data *spotifyRepo.SpotifySearchResponse) *spotify.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTracksObject, 0)

	for _, item := range data.Track.Items {

		artistName := make([]string, len(item.Artists))
		for idx, artist := range item.Artists {
			artistName[idx] = artist.Name
		}

		imageUrls := make([]string, len(item.Album.Images))
		for idx, image := range item.Album.Images {
			imageUrls[idx] = image.URL
		}
		items = append(items, spotify.SpotifyTracksObject{
			//Album related fieds
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesURL:   imageUrls,
			AlbumName:        item.Album.Name,

			//Artist related fields
			ArtistsName: artistName,

			//tracks related fields
			Explicit: item.Explicit,
			ID:       item.ID,
			Name:     item.Name,
		})
	}

	return &spotify.SearchResponse{
		Limit:  data.Track.Limit,
		Offset: data.Track.Offset,
		Items:  items,
		Total:  data.Track.Total,
	}
}
