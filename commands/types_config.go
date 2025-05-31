package commands

import "github.com/joshalling/pokedexcli/pokeapi"

type Pokedata struct {
	Pokedex       map[string]pokeapi.Pokemon
	LocationsNext string
	LocationsPrev string
}
type Config struct {
	Client pokeapi.Client
	Data   Pokedata
}
