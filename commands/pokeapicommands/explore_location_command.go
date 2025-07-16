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
		Url  string `json:"url"`
	} `json:"location"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func CommandExplore(config *cmd_utilities.Config, cache *pokecache.Cache, name string) error {

	// if config.Next == "" {
	// 	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	// 	offsetAndLimit := "?offset=0&limit=20"
	// 	config.Next = baseUrl + offsetAndLimit
	// }

	// url := config.Next
	url := "https://pokeapi.co/api/v2/location-area/" + name

	locationPokemon := LocationPokemon{}

	cachedData, ok := cache.Get(url)
	if ok {
		unmarshalErr := json.Unmarshal(cachedData, &locationPokemon)
		if unmarshalErr != nil {
			return unmarshalErr
		}
	} else {

		// resp, respErr := http.Get(config.Next)
		resp, respErr := http.Get(url)
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

		// cache.Add(url, data)

		unmarshalErr := json.Unmarshal(data, &locationPokemon)
		if unmarshalErr != nil {
			return unmarshalErr
		}
	}

	// config.Next = locationAreas.Next
	// config.Previous = locationAreas.Previous

	for _, pokemon := range locationPokemon.PokemonEncounters {

		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
