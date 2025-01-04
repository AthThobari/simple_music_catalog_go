package spotify

type SearchResponse struct {
	Items  []SpotifyTracksObject `json:"items"`
	Limit  int                   `json:"limit"`
	Offset int                   `json:"offset"`
	Total  int                   `json:"total"`
}

type SpotifyTracksObject struct {
	//Album related fieds
	AlbumType        string   `json:"albumType"`
	AlbumTotalTracks int      `json:"totalTracks"`
	AlbumImagesURL   []string `json:"albumImagesURL"`
	AlbumName        string   `json:"albumName"`

	//Artist related fields
	ArtistsName []string `json:"artistsName"`

	//tracks related fields
	Explicit     bool     `json:"explicit"`
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	IsLiked *bool `json:"isLiked"`
}


