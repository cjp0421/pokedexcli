package pokeapicommands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
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

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	fmt.Println(pokemon.Name, pokemon.BaseExperience)

	randomNumber := rand.IntN(100)
	fmt.Println(randomNumber)

	if pokemon.BaseExperience >= 300 {
		if randomNumber >= 90 {
			fmt.Println("Caught 'em!")
		} else {
			fmt.Println(pokemon.Name + " got away!")
		}
	} else if pokemon.BaseExperience >= 200 {
		if randomNumber >= 80 {
			fmt.Println("Caught 'em!")
		} else {
			fmt.Println(pokemon.Name + " got away!")
		}
	} else if pokemon.BaseExperience >= 100 {
		if randomNumber >= 60 {
			fmt.Println("Caught 'em!")
		} else {
			fmt.Println(pokemon.Name + " got away!")
		}
	} else if pokemon.BaseExperience >= 50 {
		if randomNumber >= 45 {
			fmt.Println("Caught 'em!")
		} else {
			fmt.Println(pokemon.Name + " got away!")
		}
	} else if pokemon.BaseExperience >= 0 {
		if randomNumber >= 35 {
			fmt.Println("Caught 'em!")
		} else {
			fmt.Println(pokemon.Name + " got away!")
		}
	}

	return nil
}
