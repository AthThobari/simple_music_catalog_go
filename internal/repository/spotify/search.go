package spotify

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type SpotifySearchResponse struct {
	Track SpotifyTracks `json:"tracks"`
}

type SpotifyTracks struct {
	HREF     string                `json:"href"`
	Limit    int                   `json:"limit"`
	Next     *string               `json:"next"`
	Previous *string               `json:"previous"`
	Offset int `json:"offset"`
	Items    []SpotifyTracksObject `json:"items"`
}

type SpotifyTracksObject struct {
	Album   SpotifyAlbumObject    `json:"album"`
	Artists []SpotifyArtistObject `json:"artists"`
	Explicit bool                  `json:"explicit"`
	HREF    string                `json:"href"`
	ID      string                `json:"id"`
	Name    string                `json:"name"`
}

type SpotifyAlbumObject struct {
	AlbumType   string              `json:"album_type"`
	TotalTracks int                 `json:"total_tracks"`
	Images      []SpotifyAlbumImage `json:"images"`
	Name        string              `json:"name"`
}

type SpotifyAlbumImage struct {
	URL string `json:"url"`
}

type SpotifyArtistObject struct {
	HREF string `json:"href"`
	Name string `json:"name"`
}

func (o *outbound) Search(ctx context.Context, query string, limit, offset int) (*SpotifySearchResponse, error){
	url := "https://api.spotify.com/v1/search"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error().Err(err).Msg("error create search request for spotify")
		return nil, err
	}
	bearerToken:= ""
	req.Header.Set("Authorization", bearerToken)

	resp, err := o.client.Do(req)
	if resp != nil {
		log.Error().Err(err).Msg("error execute search request for spotify")
		return nil, err
	}
	defer resp.Body.Close()

	var response SpotifySearchResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshal search response from spotify")
		return nil, err
	}
	return &response, nil
}
