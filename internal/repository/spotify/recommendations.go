package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rs/zerolog/log"
)
func (o *outbound) GetRecommendation(ctx context.Context, limit int, trackID string) (*SpotifyRecommendationsResponse, error) {
	params := url.Values{}
	params.Set("limit", strconv.Itoa(limit))
	params.Set("seed_tracks", trackID)  // Hanya gunakan seed_tracks

	basePath := `https://api.spotify.com/v1/recommendations`

	urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
	log.Debug().Msgf("Generated URL: %s", urlPath)
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("error reading response body from spotify")
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Error().Int("status", resp.StatusCode).Msgf("non-200 response from spotify: %s", string(body))
		return nil, fmt.Errorf("spotify API error: status %d", resp.StatusCode)
	}

	log.Debug().Msgf("response body: %s", body)

	var response SpotifyRecommendationsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Error().Err(err).Msg("error unmarshal search response from spotify")
		return nil, err
	}

	return &response, nil
}
