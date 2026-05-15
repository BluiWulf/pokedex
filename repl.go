package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/bluiwulf/pokedex/internal/pokeapi"
	"os"
	"strings"
)

type apiConfig struct {
	apiClient pokeapi.Client
	Next 	  *string
	Previous  *string
}

type cliCommand struct {
	name		string
	description string
	callback	func(*apiConfig) error
}

func startPokedex(cfg *apiConfig) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		usrCmd := words[0]
		
		cmd, valid := getCommands()[usrCmd]
		if valid {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println("Error occurred: ", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

// Command functions

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: 		 	"exit",
			description: 	"Exit the Pokedex",
			callback:		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"map": {
			name: 			"map",
			description: 	"Displays the next 20 location areas",
			callback: 		commandMap,
		},
		"mapb": {
			name: 			"mapb",
			description: 	"Displays the previous 20 location areas",
			callback:		commandMapb,
		},
	}
}

func commandExit(cfg *apiConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return errors.New("Failed to close Pokedex")
}

func commandHelp(cfg *apiConfig) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *apiConfig) error {
	areaResp, err := cfg.apiClient.ListAreas(cfg.Next)
	if err != nil {
		return errors.New("Failed to get location areas")
	}

	cfg.Next = areaResp.Next
	cfg.Previous = areaResp.Previous

	for _, result := range areaResp.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(cfg *apiConfig) error {
	if cfg.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	areaResp, err := cfg.apiClient.ListAreas(cfg.Previous)
	if err != nil {
		return errors.New("Failed to get location areas")
	}

	cfg.Next = areaResp.Next
	cfg.Previous = areaResp.Previous

	for _, result := range areaResp.Results {
		fmt.Println(result.Name)
	}

	return nil
}

// Helper functions

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
