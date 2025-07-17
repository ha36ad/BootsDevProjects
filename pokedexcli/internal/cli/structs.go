package cli

import "github.com/ha36ad/BootsDevProjects/pokedexcli/internal/pokecache"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}

type Config struct {
	NextLocationURL     *string
	PreviousLocationURL *string
	Cache               *pokecache.Cache
	Pokedex             map[string]pokemonResponse
}
