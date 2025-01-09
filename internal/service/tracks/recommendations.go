package tracks

import (
	"context"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/spotify"
	"github.com/AthThobari/simple_music_catalog_go/internal/models/trackactivities"
	spotifyRepo "github.com/AthThobari/simple_music_catalog_go/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) GetRecommendation(ctx context.Context, userID uint, limit int, trackID string) (*spotify.RecommendationResponse, error) {
	trackDetails, err := s.spotifyOutboound.GetRecommendation(ctx, limit, trackID)
	if err != nil {
		log.Error().Err(err).Msg("error get recommendation from spotify outbound")
		return nil, err
	}

	trackIDs := make([]string, len(trackDetails.Track.Items))
	for idx, item := range trackDetails.Track.Items {
		trackIDs[idx] = item.ID
	}

	trackactivities, err := s.trackActivitiesRepo.GetBulkSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("error get data activities from database")
		return nil, err
	}
	return modelToRecommendationResponse(trackDetails, trackactivities), nil
}

func modelToRecommendationResponse(data *spotifyRepo.SpotifySearchResponse, mapTrackActivities map[string]trackactivities.TrackActivity) *spotify.RecommendationResponse {
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
			IsLiked:  mapTrackActivities[item.ID].IsLiked,
		})
	}

	return &spotify.RecommendationResponse{
		Items:  items,
	}
}
