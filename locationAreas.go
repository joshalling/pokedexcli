package main

import "fmt"

func commandMap(*Config) (Config, error) {
	areasResponse, err := apiGetAreas("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return Config{}, err
	}
	for _, area := range areasResponse.Results {
		fmt.Println(area.Name)
	}
	return Config{
		Next:     areasResponse.Next,
		Previous: areasResponse.Previous,
	}, nil
}

func commandMapb(config *Config) (Config, error) {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return *config, nil
	}
	areasResponse, err := apiGetAreas(config.Previous)
	if err != nil {
		return *config, err
	}
	for _, area := range areasResponse.Results {
		fmt.Println(area.Name)
	}
	return Config{
		Next:     areasResponse.Next,
		Previous: areasResponse.Previous,
	}, nil
}
