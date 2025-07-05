package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	Next     string
	Previous *string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
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
		"mapb": {
			Name:        "mapb",
			Description: "Show previous 20 map locations",
			Callback:    commandMapBack,
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

func commandMap(config *Config) error {

	if config.Next == "" {
		baseUrl := "https://pokeapi.co/api/v2/location-area/"
		offsetAndLimit := "?offset=0&limit=20"
		config.Next = baseUrl + offsetAndLimit
	}

	resp, respErr := http.Get(config.Next)
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

	locationAreas := LocationArea{}

	unmarshalErr := json.Unmarshal(data, &locationAreas)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous

	for _, result := range locationAreas.Results {

		fmt.Println(result.Name)
	}

	return nil
}

func commandMapBack(config *Config) error {
	if config.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	resp, respErr := http.Get(*config.Previous)
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

	locationAreas := LocationArea{}

	unmarshalErr := json.Unmarshal(data, &locationAreas)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous

	for _, result := range locationAreas.Results {

		fmt.Println(result.Name)
	}

	return nil
}
