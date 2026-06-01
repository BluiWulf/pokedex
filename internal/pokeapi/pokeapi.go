package pokeapi

import (
	"encoding/json"
	"github.com/bluiwulf/pokedex/internal/pokecache"
	"io"
	"net/http"
	"time"

	"fmt"
)

const (
	apiURL   = "https://pokeapi.co/api/v2/"
	areasURL = apiURL + "location-area/"
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

type EncounterVersion struct {
	Rate 				 int				   `json:"rate"`
	Version 			 BasicData			   `json:"version"`
}

type EncounterMethodRate struct {
	EncounterMethod		 BasicData			   `json:"encounter_method"`
	VersionDetails		 []EncounterVersion	   `json:"version_details"`
}

type Name struct {
	Name				 string				   `json:"name"`
	Language			 BasicData			   `json:"language"`
}

type EncounterDetail struct {
	MinLvl				 int				   `json:"min_level"`
	MaxLvl				 int				   `json:"max_level"`
	ConditionVals		 []int				   `json:"condition_values"`
	Chance 				 int				   `json:"chance"`
	Method				 BasicData			   `json:"method"`
}

type PokemonVersion struct {
	Version				 BasicData			   `json:"version"`
	MaxChance			 int				   `json:"max_chance"`
	EncounterDetails	 []EncounterDetail	   `json:"encounter_details"`
}

type PokemonEncounter struct {
	Pokemon				 BasicData			   `json:"pokemon"`
	VersionDetails		 []PokemonVersion	   `json:"version_details"`
}

type LocationArea struct {
	Id		   			 int		  		   `json:"id"`
	Name	   			 string	  			   `json:"name"`
	GameIndex  			 int		  		   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location			 BasicData			   `json:"location"`
	Names				 []Name				   `json:"names"`
}

type RespAreas struct {
	Count	   			 int	      		   `json:"count"`
	Next	   			 *string    		   `json:"next"`
	Previous   			 *string    		   `json:"previous"`
	Results	   			 []BasicData 		   `json:"results"`
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
		fmt.Println()
		fmt.Println("Here's the error")
		fmt.Println()
		return RespAreas{}, err
	}

	return areaResp, nil
}

