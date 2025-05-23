package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageUrl string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageUrl != "" {
		url = pageUrl
	}

	var data []byte
	areasResponse := LocationAreaResponse{}

	cacheData, ok := c.cache.Get(url)
	if ok {
		data = cacheData
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaResponse{}, fmt.Errorf("error creating request: %w", err)
		}

		res, err := c.http.Do(req)
		if err != nil {
			return LocationAreaResponse{}, fmt.Errorf("error requesting areas: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaResponse{}, fmt.Errorf("error reading areas: %w", err)
		}

		c.cache.Add(url, data)
	}

	err := json.Unmarshal(data, &areasResponse)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error decoding areas: %w", err)
	}

	return areasResponse, nil
}

func (c *Client) GetLocation(id string) (LocationArea, error) {
	url := fmt.Sprintf("%s/location-area/%s", baseURL, id)

	var data []byte
	areasResponse := LocationArea{}

	cacheData, ok := c.cache.Get(url)
	if ok {
		data = cacheData
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationArea{}, fmt.Errorf("error creating area request: %w", err)
		}

		res, err := c.http.Do(req)
		if err != nil {
			return LocationArea{}, fmt.Errorf("error requesting area: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, fmt.Errorf("error reading area: %w", err)
		}

		c.cache.Add(url, data)
	}

	err := json.Unmarshal(data, &areasResponse)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error decoding area: %w", err)
	}

	return areasResponse, nil
}
