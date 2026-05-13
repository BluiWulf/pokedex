package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name		string
	description string
	callback	func() error
}

func startPokedex() {
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
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return errors.New("Failed to close Pokedex")
}

func commandHelp() error {
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

// Helper functions

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
