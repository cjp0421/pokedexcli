package commands

import (
	"fmt"

	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
)

func CommandHelp(config *cmd_utilities.Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}
