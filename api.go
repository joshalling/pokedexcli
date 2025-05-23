package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func apiGetAreas(url string) (LocationAreaResponse, error) {
	res, err := http.Get(url)

	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Error requesting areas: %w", err)
	}
	defer res.Body.Close()

	areasResponse := LocationAreaResponse{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&areasResponse)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Error decoding areas: %w", err)
	}

	return areasResponse, nil
}
