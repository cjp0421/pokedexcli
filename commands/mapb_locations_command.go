package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	apiutilities "github.com/cjp0421/pokedexcli/commands/api_utilities"
	"github.com/cjp0421/pokedexcli/commands/cmd_utilities"
	"github.com/cjp0421/pokedexcli/internal/pokecache"
)

func CommandMapBack(config *cmd_utilities.Config, cache *pokecache.Cache, _ string, _ *map[string]Pokemon) error {
	if config.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	url := config.Previous
	locationAreas := apiutilities.LocationArea{}

	cachedData, ok := cache.Get(url)
	if ok {
		unmarshalErr := json.Unmarshal(cachedData, &locationAreas)
		if unmarshalErr != nil {
			return unmarshalErr
		}
	} else {
		resp, respErr := http.Get(config.Previous)
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

		unmarshalErr := json.Unmarshal(data, &locationAreas)
		if unmarshalErr != nil {
			return unmarshalErr
		}
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous

	for _, result := range locationAreas.Results {

		fmt.Println(result.Name)
	}

	return nil
}
