package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func displayLocationData(data []byte, cfg *Config) error {
	var locationData locationAreaResponse
	if err := json.Unmarshal(data, &locationData); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	for _, area := range locationData.Results {
		fmt.Println(area.Name)
	}

	cfg.NextLocationURL = locationData.Next
	cfg.PreviousLocationURL = locationData.Previous
	if cfg.NextLocationURL == nil || *cfg.NextLocationURL == "" {
		fmt.Println("No more locations to show.")
	}

	return nil
}

func displayExploreData(data []byte) error {
	var areaData exploreAreaResponse
	if err := json.Unmarshal(data, &areaData); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range areaData.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func fetchData(cfg *Config, url string) ([]byte, error) {
	if data, found := cfg.Cache.Get(url); found {
		return data, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	cfg.Cache.Add(url, body)
	return body, nil
}
