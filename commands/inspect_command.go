package commands

import (
	"fmt"

	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
	"github.com/cjp0421/pokedexcli/internal/pokecache"
)

func CommandInspect(config *cmd_utilities.Config, cache *pokecache.Cache, cliArgument string, pokedex *map[string]Pokemon) error {

	if caughtPokemon, ok := (*pokedex)[cliArgument]; !ok {
		fmt.Printf("you have not caught %s\n", cliArgument)
		return nil
	} else {
		fmt.Printf("Name: %s\n", caughtPokemon.Name)
		fmt.Printf("Base Experience: %v\n", caughtPokemon.BaseExperience)

		if len(caughtPokemon.Forms) > 1 {
			fmt.Printf("Forms:\n")
			for _, form := range caughtPokemon.Forms {
				fmt.Printf("- %s\n", form.Name)
			}
		}
		fmt.Printf("Height: %v\n", caughtPokemon.Height)
		fmt.Printf("Weight: %v\n", caughtPokemon.Weight)

		fmt.Printf("Moves:\n")
		for _, move := range caughtPokemon.Moves {
			fmt.Printf("- %s\n", move.Move.Name)
		}

	}

	return nil
}
