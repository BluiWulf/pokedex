package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bluiwulf/pokedex/internal/pokecache"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type cliCommand struct {
	name		string
	description string
	callback	func() error
}

func startPokedex() {
	const interval = 30 * time.Second
	scanner := bufio.NewScanner(os.Stdin)
	cfg := apiConfig{
		Next: 	  "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
		PokeCache: pokecache.NewCache(interval),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		usrCmd := words[0]
		
		cmd, valid := cfg.getCommands()[usrCmd]
		if valid {
			err := cmd.callback()
			if err != nil {
				fmt.Println("Error occurred: ", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

// Command functions

func (cfg *apiConfig) getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: 		 	"exit",
			description: 	"Exit the Pokedex",
			callback:		cfg.commandExit,
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		cfg.commandHelp,
		},
		"map": {
			name: 			"map",
			description: 	"Displays the next 20 location areas",
			callback: 		cfg.commandMap,
		},
		"mapb": {
			name: 			"mapb",
			description: 	"Displays the previous 20 location areas",
			callback:		cfg.commandMapb,
		},
	}
}

func (cfg *apiConfig) commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return errors.New("Failed to close Pokedex")
}

func (cfg *apiConfig) commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range cfg.getCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func (cfg *apiConfig) commandMap() error {
	entry, ok := cfg.PokeCache.Get(cfg.Next)
	if !ok {
		res, err := http.Get(cfg.Next)
		if err != nil {
			return errors.New("Failed to get location-areas from PokeAPI")
		}
		defer res.Body.Close()

		entry, err = io.ReadAll(res.Body)
		if err != nil {
			return errors.New("Failed to read response body")
		}
		cfg.PokeCache.Add(cfg.Next, entry)
	}

	resp := apiResp{}
	err := json.Unmarshal(entry, &resp)
	if err != nil {
		return errors.New("Failed to unmarshal json from PokeAPI")
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func (cfg *apiConfig) commandMapb() error {
	if cfg.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	entry, ok := cfg.PokeCache.Get(cfg.Previous)
	if !ok {
		res, err := http.Get(cfg.Previous)
		if err != nil {
			return errors.New("Failed to get location-areas from PokeAPI")
		}
		defer res.Body.Close()

		entry, err = io.ReadAll(res.Body)
		if err != nil {
			return errors.New("Failed to read response body")
		}
		cfg.PokeCache.Add(cfg.Previous, entry)
	}

	resp := apiResp{}
	err := json.Unmarshal(entry, &resp)
	if err != nil {
		return errors.New("Failed to unmarshal json from PokeAPI")
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous
	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}

	return nil
}

// Helper functions

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
