package pokeapi

import (
	"net/http"
	"time"

	"github.com/joshalling/pokedexcli/pokecache"
)

type Client struct {
	http  http.Client
	cache *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		http: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(10),
	}
}

const (
	baseURL = "https://pokeapi.co/api/v2"
)
