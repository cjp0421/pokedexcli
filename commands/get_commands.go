package commands

import (
	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
	"github.com/cjp0421/pokedexcli/commands/pokeapicommands"
	"github.com/cjp0421/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*cmd_utilities.Config, *pokecache.Cache, string) error
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
		"explore": {
			Name:        "explore",
			Description: "Show pokemon at location",
			Callback:    pokeapicommands.CommandExplore,
		},
	}
}
