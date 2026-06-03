package main

import (
	"log"
	"time"

	"github.com/bluiwulf/pokedex/internal/pokeapi"
)

func main() {
	apiClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &apiConfig{
		apiClient: apiClient,
		caught: make(map[string]pokeapi.PokemonInfo),
	}
	err := checkCommands()
	if err != nil {
		log.Fatal(err)
	}
	startPokedex(cfg)
}
