package commands

import "github.com/joshalling/pokedexcli/pokeapi"

type Config struct {
	Client        pokeapi.Client
	LocationsNext string
	LocationsPrev string
}
