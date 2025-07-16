package pokeapicommands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
	"github.com/cjp0421/pokedexcli/internal/pokecache"
)

func CommandCatch(config *cmd_utilities.Config, cache *pokecache.Cache, cliArgument string) error {

	// if config.LocationUrl == "" {
	// 	baseUrl := "https://pokeapi.co/api/v2/pokemon/"
	// 	config.LocationUrl = baseUrl + cliArgument
	// }

	// url := config.LocationUrl
	url := "https://pokeapi.co/api/v2/pokemon/" + cliArgument

	pokemon := Pokemon{}

	cachedData, ok := cache.Get(url)
	if ok {
		unmarshalErr := json.Unmarshal(cachedData, &pokemon)
		if unmarshalErr != nil {
			return unmarshalErr
		}
	} else {

		// resp, respErr := http.Get(config.LocationUrl)
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

		unmarshalErr := json.Unmarshal(data, &pokemon)
		if unmarshalErr != nil {
			return unmarshalErr
		}
	}

	fmt.Println(pokemon.Name, pokemon.BaseExperience)

	return nil
}
