package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			line := scanner.Text()
			words := cleanInput(line)
			if len(words) == 0 {
				continue
			}
			first := words[0]
			fmt.Printf("Your command was: %s\n", first)
			break
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
