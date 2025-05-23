package commands

import (
	"fmt"
)

func Explore(c *Config, args []string) error {
	id := args[0]
	location, err := c.Client.GetLocation(id)
	if err != nil {
		return err
	}
	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
