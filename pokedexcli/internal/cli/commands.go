package cli

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func CommandExit(cfg *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(commands map[string]CliCommand) func(cfg *Config, args []string) error {
	return func(cfg *Config, args []string) error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		for _, cmd := range commands {
			fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
		}
		return nil
	}
}

func CommandMap(cfg *Config, args []string) error {
	var url string
	if cfg.NextLocationURL == nil || *cfg.NextLocationURL == "" {
		url = "https://pokeapi.co/api/v2/location-area/?limit=20"
	} else {
		url = *cfg.NextLocationURL
	}

	data, err := fetchData(cfg, url)
	if err != nil {
		return err
	}

	return displayLocationData(data, cfg)
}

func CommandMapBack(cfg *Config, args []string) error {
	if cfg.PreviousLocationURL == nil || *cfg.PreviousLocationURL == "" {
		fmt.Println("You're on the first page.")
		return nil
	}

	cfg.NextLocationURL = cfg.PreviousLocationURL
	return CommandMap(cfg, args)
}

func CommandExplore(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please specify a location area name")
	}
	areaName := strings.ToLower(args[0])
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", areaName)

	fmt.Printf("Exploring %s...\n", areaName)

	data, err := fetchData(cfg, url)
	if err != nil {
		return err
	}

	return displayExploreData(data)
}

func CommandCatch(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please specify a pokemon to catch")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	data, err := fetchData(cfg, url)
	if err != nil {
		return err
	}

	var pokemonData pokemonResponse
	if err := json.Unmarshal(data, &pokemonData); err != nil {
		return fmt.Errorf("error parsing pokemon data: %w", err)
	}

	baseExp := pokemonData.BaseExperience
	catchChance := max(100-baseExp, 10)

	success := false
	for range 10 {
		randomValue := rand.Intn(100) + 1
		if randomValue <= catchChance {
			success = true
			break
		}
	}

	if success {
		fmt.Printf("You caught %s!\n", pokemonName)
		if cfg.Pokedex == nil {
			cfg.Pokedex = make(map[string]pokemonResponse)
		}
		cfg.Pokedex[pokemonName] = pokemonData
	} else {
		fmt.Printf("%s escaped after several tries!\n", pokemonName)
	}
	return nil
}

func CommandInspect(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please specify a pokemon to inspect")
	}
	pokemonName := args[0]
	pokemon, exists := cfg.Pokedex[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

func CommandPokedex(cfg *Config, args []string) error {
	if cfg.Pokedex == nil {
		fmt.Println("You haven't caught any Pokemon yet.")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.Pokedex {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil

}

func GetCommands() map[string]CliCommand {
	commands := make(map[string]CliCommand)
	commands["exit"] = CliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    CommandExit,
	}
	commands["help"] = CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    CommandHelp(commands),
	}
	commands["map"] = CliCommand{
		Name:        "map",
		Description: "Explore the Pokemon world map",
		Callback:    CommandMap,
	}
	commands["mapb"] = CliCommand{
		Name:        "mapb",
		Description: "Go back to the previous Pokemon world map page",
		Callback:    CommandMapBack,
	}
	commands["explore"] = CliCommand{
		Name:        "explore",
		Description: "Explore a location area to see available Pokemon",
		Callback:    CommandExplore,
	}
	commands["catch"] = CliCommand{
		Name:        "catch",
		Description: "Attempt to catch a pokemon",
		Callback:    CommandCatch,
	}
	commands["inspect"] = CliCommand{
		Name:        "inspect",
		Description: "Inspect a caught pokemon",
		Callback:    CommandInspect,
	}
	commands["pokedex"] = CliCommand{
		Name:        "pokedex",
		Description: "List caught pokemon",
		Callback:    CommandPokedex,
	}
	return commands
}
