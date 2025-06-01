package commands

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config, []string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch [pokemon]",
			description: "Catch a Pokémon",
			Callback:    Catch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"explore": {
			name:        "explore [location]",
			description: "Explore a location area",
			Callback:    Explore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    Help,
		},
		"inspect": {
			name:        "inspect [pokemon]",
			description: "Inspect a Pokémon",
			Callback:    Inspect,
		},
		"load": {
			name:        "load",
			description: "Loads game data into config",
			Callback:    Load,
		},
		"map": {
			name:        "map",
			description: "List all location areas",
			Callback:    Mapf,
		},
		"mapb": {
			name:        "mapb",
			description: "List all location areas on previous page",
			Callback:    Mapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all Pokémon in the Pokedex",
			Callback:    Pokedex,
		},
		"print": {
			name:        "print [pokemon]",
			description: "Print a Pokémon to the terminal",
			Callback:    Print,
		},
		"save": {
			name:        "save",
			description: "Save your current state",
			Callback:    Save,
		},
	}
}
