package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	http http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		http: http.Client{
			Timeout: timeout,
		},
	}
}

const (
	baseURL = "https://pokeapi.co/api/v2"
)
