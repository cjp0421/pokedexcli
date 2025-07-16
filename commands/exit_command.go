package commands

import (
	"fmt"
	"os"

	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
	"github.com/cjp0421/pokedexcli/internal/pokecache"
)

func CommandExit(config *cmd_utilities.Config, cache *pokecache.Cache, cliArgument string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
