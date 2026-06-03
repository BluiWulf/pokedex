package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/bluiwulf/pokedex/internal/pokeapi"
	"math/rand"
	"os"
	"strings"
)

type apiConfig struct {
	apiClient 	pokeapi.Client
	Next 	  	*string
	Previous  	*string
	caught		map[string]pokeapi.PokemonInfo
}

type cliCommand struct {
	name		string
	description string
	callback	func(*apiConfig, ...string) error
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
		args := words[1:]
		
		_, cmds := getCommands()
		cmd, valid := cmds[usrCmd]
		if valid {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println("Error occurred: ", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

// Command functions

func getCommands() ([]string, map[string]cliCommand) {
	keys := []string{
		"map",
		"mapb",
		"explore",
		"catch",
		"inspect",
		"pokedex",
		"help",
		"exit",
	}
	cmds := map[string]cliCommand {
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
		"explore": {
			name:			"explore",
			description:	"Displays a list of all Pokemon in a given area (must provide name of Location Area)",
			callback:		commandExplore,
		},
		"catch": {
			name:			"catch",
			description:	"Attempts to catch the Pokemon (must provide name of Pokemon)",
			callback:		commandCatch,
		},
		"inspect": {
			name:			"inspect",
			description:	"Display details about caught Pokemon (must provide name of Pokemon)",
			callback:		commandInspect,
		},
		"pokedex": {
			name:			"pokedex",
			description:	"Displays a list of all Pokemon caught and registered in the Pokedex",
			callback:		commandPokedex,
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"exit": {
			name: 		 	"exit",
			description: 	"Exit the Pokedex",
			callback:		commandExit,
		},
	}
	return keys, cmds
}

func commandExit(cfg *apiConfig, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return errors.New("failed to close Pokedex")
}

func commandHelp(cfg *apiConfig, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	keys, cmds := getCommands()

	for _, name := range keys {
		cmd := cmds[name]
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *apiConfig, args ...string) error {
	areaResp, err := cfg.apiClient.ListAreas(cfg.Next)
	if err != nil {
		return errors.New("failed to get location areas")
	}

	cfg.Next = areaResp.Next
	cfg.Previous = areaResp.Previous

	for _, result := range areaResp.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(cfg *apiConfig, args ...string) error {
	if cfg.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	areaResp, err := cfg.apiClient.ListAreas(cfg.Previous)
	if err != nil {
		return errors.New("failed to get location areas")
	}

	cfg.Next = areaResp.Next
	cfg.Previous = areaResp.Previous

	for _, result := range areaResp.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandExplore(cfg *apiConfig, args ...string) error {
	if len(args) == 0 {
		return errors.New("location area must be provided")
	}
	if len(args) > 1 {
		return errors.New("only provide one location area")
	}
	area := args[0]

	location, err := cfg.apiClient.GetLocation(&area)
	if err != nil {
		return errors.New("failed to get location area information")
	}

	fmt.Printf("Exploring %v...\n", area)
	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *apiConfig, args ...string) error {
	if len(args) == 0 {
		return errors.New("name of Pokemon must be provided")
	}
	if len(args) > 1 {
		return errors.New("only provide one Pokemon to catch")
	}
	name := args[0]

	pokemon, err := cfg.apiClient.GetPokemon(&name)
	if err != nil {
		return errors.New("failed to get Pokemon information")
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	baseExp := pokemon.BaseXP
	maxExp := 200
	if maxExp < baseExp {
		maxExp += (((baseExp - maxExp) / 100) + 1) * 100
	}

	chance := rand.Intn(maxExp)
	if chance > baseExp {
		fmt.Printf("%v was caught!\n", name)
		cfg.caught[name] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", name)
	}

	return nil
}

func commandInspect(cfg *apiConfig, args ...string) error {
	if len(args) == 0 {
		return errors.New("name of Pokemon must be provided")
	}
	if len(args) > 1 {
		return errors.New("only provide one Pokemon to inspect")
	}
	name := args[0]

	pokemon, ok := cfg.caught[name]
	if !ok {
		fmt.Println("you have not caught that Pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf("  - %v\n", types.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *apiConfig, args ...string) error {
	if len(cfg.caught) == 0 {
		fmt.Println("you have not caught any Pokemon")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caught {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil
}

// Helper functions

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func checkCommands() error {
	keys, cmds := getCommands()
	for _, name := range keys {
		_, valid := cmds[name]
		if !valid {
			return fmt.Errorf("unknown command: %s\n", name)
		}
	}
	return nil
}

