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

type Location struct {
	Name 	   string 	  `json:"name"`
	Url	 	   string 	  `json:"url"`
}

type RespAreas struct {
	Count	   int	      `json:"count"`
	Next	   *string    `json:"next"`
	Previous   *string    `json:"previous"`
	Results	   []Location `json:"results"`
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

