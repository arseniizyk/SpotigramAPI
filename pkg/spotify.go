package spotify

import (
	"SpotigramAPI/internal/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CurrentlyPlaying(access_token string, cfg config.Config) (config.Response, int) {
	endpoint := fmt.Sprintf("%s/player/currently-playing", cfg.SpotifyEndpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println("cant create request")
	}
	req.Header.Add("Authorization", "Bearer "+access_token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return config.Response{
			Error: err,
		}, http.StatusForbidden
	}
	if resp.StatusCode == http.StatusNoContent {
		return config.Response{}, http.StatusNoContent
	}
	decoder := json.NewDecoder(resp.Body)
	var resultResponse config.Response
	err = decoder.Decode(&resultResponse)
	if err != nil {
		log.Println("cant decode json: ", err)
	}

	return resultResponse, http.StatusOK
}
