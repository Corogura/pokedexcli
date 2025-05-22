package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		var input string
		scanner.Scan()
		input = scanner.Text()
		cleanedInput := cleanInput(input)
		fmt.Printf("Your command was: %v\n", cleanedInput[0])
	}
}
