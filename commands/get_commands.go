package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
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
			Description: "Show map locations",
			Callback:    commandMap,
		},
	}
}

type LocationArea struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap() error {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"

	resp, respErr := http.Get(baseUrl)
	if respErr != nil {
		fmt.Println("Error making HTTP request:", respErr)
		return respErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %d", err)
		return nil
	}

	// fmt.Println(string(body))

	locationAreas := LocationArea{}
	// fmt.Println(body)

	unmarshalErr := json.Unmarshal(data, &locationAreas)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	for _, result := range locationAreas.Results {

		fmt.Println(result.Name)
	}

	return nil
}
