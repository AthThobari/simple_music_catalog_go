package trackactivities

import (
	"context"

	"github.com/AthThobari/simple_music_catalog_go/internal/models/trackactivities"
)

func (r *repository) Create(ctx context.Context, model trackactivities.TrackActivity) error {
	return r.db.Create(&model).Error
}

func (r *repository) Update(ctx context.Context, model trackactivities.TrackActivity) error {
	return r.db.Save(&model).Error
}

func (r *repository) Get(ctx context.Context, userID uint, spotifyID string) (*trackactivities.TrackActivity, error) {
	activity := trackactivities.TrackActivity{}
	res := r.db.Where("user_id = ?", userID).Where("spotify_id = ?", spotifyID ).First(&activity)
	if res.Error != nil {
		return nil, res.Error
	}
	return &activity, nil
}

func (r *repository) GetBulkSpotifyIDs(ctx context.Context, userID uint, spotifyIDs []string) (map[string]trackactivities.TrackActivity, error) {
	activities := make([]trackactivities.TrackActivity, 0)
	res := r.db.Where("user_id = ?", userID).Where("spotify_id IN ?", spotifyIDs).Find(&activities)

	if res.Error != nil {
		return nil, res.Error
	}

	result := make(map[string]trackactivities.TrackActivity, 0)
	for _, activity := range activities {
		result[activity.SpotifyID] = activity
	}

	return result, nil
}
