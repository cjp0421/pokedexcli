package commands

import (
	"fmt"

	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
	"github.com/cjp0421/pokedexcli/internal/pokecache"
)

func PokedexCommand(_ *cmd_utilities.Config, _ *pokecache.Cache, _ string, pokedex *map[string]Pokemon) error {
	if len(*pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range *pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
