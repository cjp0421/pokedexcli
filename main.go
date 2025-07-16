package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/cjp0421/pokedexcli/commands"
	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
	"github.com/cjp0421/pokedexcli/internal/pokecache"
	"github.com/cjp0421/pokedexcli/utilities"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	config := cmd_utilities.Config{}
	cache := pokecache.NewCache(5 * time.Second)

	for {
		fmt.Print("Pokedex >")
		if !reader.Scan() {
			break
		}
		input := reader.Text()
		cleanedInput := utilities.CleanInput(string(input))

		commandName := cleanedInput[0]

		var cliArgument string
		if len(cleanedInput) > 1 {
			cliArgument = cleanedInput[1]
		}

		command, exists := commands.GetCommands()[commandName]
		if exists {
			err := command.Callback(&config, cache, cliArgument)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
