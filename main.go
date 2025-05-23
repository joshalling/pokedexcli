package main

import (
	"time"

	"github.com/joshalling/pokedexcli/commands"
	"github.com/joshalling/pokedexcli/pokeapi"
)

func main() {
	client := pokeapi.NewClient(10 * time.Second)

	c := commands.Config{
		Client: client,
	}
	startRepl(&c)
}
