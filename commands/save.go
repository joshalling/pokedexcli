package commands

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save(c *Config, _ []string) error {
	file, err := os.OpenFile("pokesave.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err = encoder.Encode(c.Data); err != nil {
		return err
	}
	fmt.Println("Game Saved!")
	return nil
}

func Load(c *Config, _ []string) error {
	file, err := os.OpenFile("pokesave.json", os.O_RDONLY, 0)
	if err != nil {
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&c.Data); err != nil {
		return err
	}
	fmt.Println("Game Loaded!")
	return nil
}
