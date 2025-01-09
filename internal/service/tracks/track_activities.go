package tracks

import (
	"context"
	"fmt"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/trackactivities"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *service) UpsertTrackActivities(ctx context.Context, userID uint, request trackactivities.TrackActivityRequest) error {
	activity, err := s.trackActivitiesRepo.Get(ctx, userID, request.SpotifyID)

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error get record from database")
		return err
	}

	if err == gorm.ErrRecordNotFound || activity == nil {
		//create user activity
		err = s.trackActivitiesRepo.Create(ctx, trackactivities.TrackActivity{
			UserID:    userID,
			SpotifyID: request.SpotifyID,
			IsLiked:   request.IsLiked,
			CreatedBy: fmt.Sprintf("%d", userID),
			UpdatedBy: fmt.Sprintf("%d", userID),
		})
		if err != nil {
			log.Error().Err(err).Msg("error create record to database")
			return err
		}
		return nil
	}
	activity.IsLiked = request.IsLiked
	err = s.trackActivitiesRepo.Update(ctx, *activity)
	if err != nil {
		log.Error().Err(err).Msg("error update record to database")
		return err
	}
	return nil
}
