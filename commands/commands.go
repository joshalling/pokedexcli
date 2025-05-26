package commands

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config, []string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "Catch a Pokémon",
			Callback:    Catch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			Callback:    Explore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    Help,
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
	}
}
