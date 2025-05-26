package commands

import (
	"fmt"
)

func Inspect(c *Config, args []string) error {
	id := args[0]
	pokemon, exists := c.Pokedex[id]
	if !exists {
		fmt.Printf("Pokemon with %s not found in Pokedex\n", id)
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("  -%s\n", t.Type.Name)
		}
	}

	return nil
}

func Pokedex(c *Config, _ []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.Pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
