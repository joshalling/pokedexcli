package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetLocations(pageUrl string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageUrl != "" {
		url = pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := c.http.Do(req)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error requesting areas: %w", err)
	}
	defer res.Body.Close()

	areasResponse := LocationAreaResponse{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&areasResponse)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error decoding areas: %w", err)
	}

	return areasResponse, nil
}
