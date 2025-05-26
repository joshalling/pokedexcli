package commands

import "github.com/joshalling/pokedexcli/pokeapi"

type Config struct {
	Client        pokeapi.Client
	Pokedex       map[string]pokeapi.Pokemon
	LocationsNext string
	LocationsPrev string
}
