package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(id string) (Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, id)

	var data []byte
	pokemonResponse := Pokemon{}

	cacheData, ok := c.cache.Get(url)
	if ok {
		data = cacheData
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error creating pokemon request: %w", err)
		}

		res, err := c.http.Do(req)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error requesting pokemon: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error reading pokemon: %w", err)
		}

		c.cache.Add(url, data)
	}

	err := json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error decoding pokemon: %w", err)
	}

	return pokemonResponse, nil
}
