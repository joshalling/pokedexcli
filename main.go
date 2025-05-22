package main

import "strings"

func main() {
	println("Hello, World!")
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
