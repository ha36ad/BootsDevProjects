# Pokedex CLI Project

A terminal-based Pokedex application written in Go, using the PokeAPI for data. Users can explore locations, catch Pokemon, inspect caught Pokemon, and list their personal Pokedex.

## Features
- Explore Pokemon world locations (paginated)
- Catch Pokemon with a randomized chance based on base experience
- Inspect caught Pokemon (height, weight, stats, types)
- List all caught Pokemon
- Caching layer to reduce API calls

## Requirements
- Go 1.18 or newer
- Internet connection for PokeAPI requests

## Commands
- `help`: Display usage information
- `exit`: Exit the application
- `map`: View the next 20 location areas
- `mapb`: Go back to the previous 20 location areas
- `explore <area_name>`: List Pokemon available in a location area
- `catch <pokemon_name>`: Attempt to catch a Pokemon
- `inspect <pokemon_name>`: View details of a caught Pokemon
- `pokedex`: List all caught Pokemon

## Setup
1. Clone the repository
2. Run `go mod tidy`
3. Build and run the application:

```bash
go run main.go