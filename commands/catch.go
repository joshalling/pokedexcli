package commands

import (
	"fmt"
	"math/rand"
)

func Catch(c *Config, args []string) error {
	id := args[0]

	pokemon, err := c.Client.GetPokemon(id)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	baseChance := 800
	result := rand.Intn(baseChance)
	if result > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		c.Pokedex[id] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
