package tracks

import (
	"context"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/spotify"
	"github.com/AthThobari/simple_music_catalog_go/internal/models/trackactivities"
	spotifyRepo "github.com/AthThobari/simple_music_catalog_go/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.spotifyOutboound.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error search track to spotify")
		return nil, err
	}

	trackIDs := make([]string, len(trackDetails.Track.Items))
	for idx, item := range trackDetails.Track.Items {
		trackIDs[idx] = item.ID
	}

	trackActivities, err := s.trackActivitiesRepo.GetBulkSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("error get track activities from database")
		return nil, err
	}

	return modelToResponse(trackDetails, trackActivities), nil
}

func modelToResponse(data *spotifyRepo.SpotifySearchResponse, mapTrackActivities map[string]trackactivities.TrackActivity) *spotify.SearchResponse {
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
			IsLiked: mapTrackActivities[item.ID].IsLiked,
		})
	}

	return &spotify.SearchResponse{
		Limit:  data.Track.Limit,
		Offset: data.Track.Offset,
		Items:  items,
		Total:  data.Track.Total,
	}
}
