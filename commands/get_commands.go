package commands

import (
	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
	"github.com/cjp0421/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*cmd_utilities.Config, *pokecache.Cache, string, *map[string]Pokemon) error
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
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Show previous 20 map locations",
			Callback:    CommandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Show pokemon at location",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempt to catch a Pokemon",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Displays information about a Pokemon you have caught",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Lists all of the Pokemon you have caught",
			Callback:    PokedexCommand,
		},
	}
}
