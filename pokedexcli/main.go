package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ha36ad/BootsDevProjects/pokedexcli/internal/cli"
	"github.com/ha36ad/BootsDevProjects/pokedexcli/internal/pokecache"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func main() {
	cache := pokecache.NewCache(5 * time.Second)

	cfg := &cli.Config{
		Cache: cache,
	}

	commands := cli.GetCommands()
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		if command, exists := commands[commandName]; exists {
			if err := command.Callback(cfg, words[1:]); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
