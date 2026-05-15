package main

import (
	"github.com/bluiwulf/pokedex/internal/pokeapi"
	"time"
)

func main() {
	apiClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &apiConfig{
		apiClient: apiClient,
	}
	startPokedex(cfg)
}

