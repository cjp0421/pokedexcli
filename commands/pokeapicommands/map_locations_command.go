package pokeapicommands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
	pokeapicommands "github.com/cjp0421/pokedexcli/commands/pokeapicommands/api_utilities"
	"github.com/cjp0421/pokedexcli/internal/pokecache"
)

func CommandMap(config *cmd_utilities.Config, cache *pokecache.Cache) error {

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

	locationAreas := pokeapicommands.LocationArea{}

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
