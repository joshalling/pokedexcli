package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joshalling/pokedexcli/commands"
)

func startRepl(config *commands.Config) {

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
		args := words[1:]

		cmd, ok := commands.GetCommands()[command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := cmd.Callback(config, args)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
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
