package pokeapicommands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
	"github.com/cjp0421/pokedexcli/internal/pokecache"
)

type LocationPokemon struct {
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	PokemonEncounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	BaseExperience int    `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
	} `json:"moves"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

func CommandExplore(config *cmd_utilities.Config, cache *pokecache.Cache, cliArgument string, _ *map[string]Pokemon) error {

	if config.LocationUrl == "" {
		baseUrl := "https://pokeapi.co/api/v2/location-area/"
		config.LocationUrl = baseUrl + cliArgument
	}

	url := config.LocationUrl

	locationPokemon := LocationPokemon{}

	cachedData, ok := cache.Get(url)
	if ok {
		unmarshalErr := json.Unmarshal(cachedData, &locationPokemon)
		if unmarshalErr != nil {
			return unmarshalErr
		}
	} else {

		resp, respErr := http.Get(config.LocationUrl)
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

		cache.Add(url, data)

		unmarshalErr := json.Unmarshal(data, &locationPokemon)
		if unmarshalErr != nil {
			return unmarshalErr
		}
	}

	for _, pokemon := range locationPokemon.PokemonEncounters {

		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
