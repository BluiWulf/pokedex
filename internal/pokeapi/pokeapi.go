package pokeapi

import (
	"encoding/json"
	"github.com/bluiwulf/pokedex/internal/pokecache"
	"io"
	"net/http"
	"time"
)

const (
	apiURL   	= "https://pokeapi.co/api/v2/"
	areasURL 	= apiURL + "location-area/"
	pokemonURL	= apiURL + "pokemon/"
)

type Client struct {
	httpClient http.Client
	PokeCache  pokecache.Cache
}

// This will be used as a generic type for all json objects with simply
// 'name' and 'url'
type BasicData struct {
	Name	   			 string				   `json:"name"`
	Url		   			 string	  			   `json:"url"`
}

func NewClient(timeout, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		PokeCache: pokecache.NewCache(interval),
	}
}

func (c *Client) ListAreas(pageURL *string) (RespAreas, error) {
	page := areasURL
	if pageURL != nil {
		page = *pageURL
	}

	entry, ok := c.PokeCache.Get(page)
	if !ok {
		req, err := http.NewRequest("GET", page, nil)
		if err != nil {
			return RespAreas{}, err
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespAreas{}, err
		}
		defer res.Body.Close()

		entry, err = io.ReadAll(res.Body)
		if err != nil {
			return RespAreas{}, err
		}
		c.PokeCache.Add(page, entry)
	}

	areaResp := RespAreas{}
	err := json.Unmarshal(entry, &areaResp)
	if err != nil {
		return RespAreas{}, err
	}

	return areaResp, nil
}

func (c *Client) GetLocation(area *string) (LocationArea, error) {
	pageURL := areasURL + *area + "/"
	entry, ok := c.PokeCache.Get(pageURL)
	if !ok {
		req, err := http.NewRequest("GET", pageURL, nil)
		if err != nil {
			return LocationArea{}, err
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}
		defer res.Body.Close()

		entry, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, err
		}
		c.PokeCache.Add(pageURL, entry)
	}

	location := LocationArea{}
	err := json.Unmarshal(entry, &location)
	if err != nil {
		return LocationArea{}, err
	}

	return location, nil
}

func (c *Client) GetPokemon(name *string) (PokemonInfo, error) {
	pageURL := pokemonURL + *name + "/"
	entry, ok := c.PokeCache.Get(pageURL)
	if !ok {
		req, err := http.NewRequest("GET", pageURL, nil)
		if err != nil {
			return PokemonInfo{}, err
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonInfo{}, err
		}
		defer res.Body.Close()

		entry, err = io.ReadAll(res.Body)
		if err != nil {
			return PokemonInfo{}, err
		}
		c.PokeCache.Add(pageURL, entry)
	}

	pokemon := PokemonInfo{}
	err := json.Unmarshal(entry, &pokemon)
	if err != nil {
		return PokemonInfo{}, err
	}

	return pokemon, nil
}

