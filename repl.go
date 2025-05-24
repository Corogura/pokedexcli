package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	pageConfig := config{
		Next:     "https://pokeapi.co/api/v2/location-area/1",
		Previous: "",
	}
	for {
		fmt.Print("Pokedex > ")
		var input string
		scanner.Scan()
		input = scanner.Text()
		cleanedInput := cleanInput(input)
		c, exists := commands[cleanedInput[0]]
		if exists {
			err := c.callback(&pageConfig)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	slicedString := strings.Fields(text)
	return slicedString
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, elem := range getCommands() {
		fmt.Printf("%v: %v\n", elem.name, elem.description)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas in the Pokemon world. Each subsequent call displays the next 20 location areas.",
			callback:    getLocation,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas that was displayed using map command",
			callback:    getLocationb,
		},
	}
}
