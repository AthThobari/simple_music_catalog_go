package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rs/zerolog/log"
)

func (o *outbound) GetRecommendation(ctx context.Context, limit int, trackID string)(*SpotifySearchResponse, error) {
	params := url.Values{}
	params.Set("limit", strconv.Itoa(limit))
	params.Set("type", "track")
	params.Set("seed_tracks", trackID)

	basePath := `https://api.spotify.com/v1/recommendations`

	urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		log.Error().Err(err).Msg("error create recommendation request for spotify")
		return nil, err
	}
	accessToken, tokenType, err := o.GetTokenDetails()
	if err != nil {
		log.Error().Err(err).Msg("error get token details")
		return nil, err
	}

	bearerToken := fmt.Sprintf("%s %s", tokenType, accessToken)
	req.Header.Set("Authorization", bearerToken)

	resp, err := o.client.Do(req)
	if err != nil {
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
