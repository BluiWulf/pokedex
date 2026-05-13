package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
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
		fmt.Printf("Your command was: %v\n", words[0])
	}
}
