package main

import (
	"fmt"
)

func commandExplore(cfg *config, locationNames ...string) error {
	for _, locationName := range locationNames {
		pokemonResp, err := cfg.pokeapiClient.ExploreLocation(locationName)
		if err != nil {
			return err
		}

		fmt.Printf("Exploring %s...\n", locationName)
		fmt.Println("Found Pokemon:")
		for _, enc := range pokemonResp.PokemonEncounters {
			fmt.Printf(" - %s\n", enc.Pokemon.Name)
		}
	}
	return nil
}
