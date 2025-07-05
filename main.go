package main

import (
	"bufio"
	"fmt"
	"os"

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

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
