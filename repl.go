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
	callback    func() error
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
	}
	scanner := bufio.NewScanner(os.Stdin)
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

		err := cmd.callback()
		if err != nil {
			fmt.Println(err)
		}

	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
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
