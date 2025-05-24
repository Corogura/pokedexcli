package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Corogura/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		var input string
		scanner.Scan()
		input = scanner.Text()
		cleanedInput := cleanInput(input)
		c, exists := commands[cleanedInput[0]]
		if exists {
			if len(cleanedInput) > 1 {
				for i := range len(cleanedInput) - 1 {
					err := c.callback(cfg, cleanedInput[i+1])
					if err != nil {
						fmt.Println(err)
					}
				}
			} else {
				err := c.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
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

func commandExit(cfg *config, s ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, s ...string) error {
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
	callback    func(cfg *config, arg ...string) error
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
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas that was displayed using map command",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of all the Pokemon located in a specified area",
			callback:    commandExplore,
		},
	}
}
