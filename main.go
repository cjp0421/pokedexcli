package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cjp0421/pokedexcli/commands"
	"github.com/cjp0421/pokedexcli/utilities"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		if !reader.Scan() {
			break
		}
		input := reader.Text()
		cleanedInput := utilities.CleanInput(string(input))

		commandName := cleanedInput[0]

		command, exists := commands.GetCommands()[commandName]
		if exists {
			err := command.Callback()
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
