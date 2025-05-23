package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) (Config, error)
}

type Config struct {
	Next     string
	Previous string
}

func startRepl() {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List all location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List all location areas on previous page",
			callback:    commandMapb,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	config := Config{}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		words := cleanInput(line)
		if len(words) == 0 {
			continue
		}
		command := words[0]

		cmd, ok := commands[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		var err error
		config, err = cmd.callback(&config)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func commandExit(_ *Config) (Config, error) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return Config{}, nil
}

func commandHelp(_ *Config) (Config, error) {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return Config{}, nil
}

func cleanInput(text string) []string {
	words := strings.Split(text, " ")
	trimmed := make([]string, 0, len(words))
	for _, word := range words {
		trimmedWord := strings.TrimSpace(word)
		if trimmedWord == "" {
			continue
		}

		trimmedWord = strings.ToLower(trimmedWord)
		trimmed = append(trimmed, trimmedWord)
	}
	return trimmed
}
