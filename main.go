package main

import (
	"github.com/bluiwulf/pokedex/internal/pokecache"
)

type apiConfig struct {
	Next 	  string
	Previous  string
	PokeCache *pokecache.Cache
}

type Location struct {
	Name 	  string `json:"name"`
	Url	 	  string `json:"url"`
}

type apiResp struct {
	Count	  int	    `json:"count"`
	Next	  string     `json:"next"`
	Previous  string     `json:"previous"`
	Results	  []Location `json:"results"`
}

func main() {
	startPokedex()
}

