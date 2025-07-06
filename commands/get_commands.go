package commands

import (
	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
	"github.com/cjp0421/pokedexcli/commands/pokeapicommands"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*cmd_utilities.Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Show 20 map locations",
			Callback:    pokeapicommands.CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Show previous 20 map locations",
			Callback:    pokeapicommands.CommandMapBack,
		},
	}
}
