package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"

	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
	"github.com/cjp0421/pokedexcli/internal/pokecache"
)

func CommandCatch(config *cmd_utilities.Config, cache *pokecache.Cache, cliArgument string, pokedex *map[string]Pokemon) error {

	if caughtPokemon, ok := (*pokedex)[cliArgument]; ok {
		fmt.Printf("%s has already been caught!\n", caughtPokemon.Name)
		return nil
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + cliArgument

	pokemon := Pokemon{}

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

	unmarshalErr := json.Unmarshal(data, &pokemon)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	randomNumber := rand.IntN(100)

	var catchingThreshold int

	if pokemon.BaseExperience >= 300 {
		catchingThreshold = 90
	} else if pokemon.BaseExperience >= 200 {
		catchingThreshold = 80
	} else if pokemon.BaseExperience >= 100 {
		catchingThreshold = 60
	} else if pokemon.BaseExperience >= 50 {
		catchingThreshold = 45
	} else if pokemon.BaseExperience >= 0 {
		catchingThreshold = 35
	}

	if randomNumber >= catchingThreshold {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		(*pokedex)[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
