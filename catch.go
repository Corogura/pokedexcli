package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, names ...string) error {
	for _, name := range names {
		infoResp, err := cfg.pokeapiClient.GetPokemonInfo(name)
		if err != nil {
			return err
		}

		fmt.Printf("Throwing a Pokeball at %s...\n", name)
		baseExp := infoResp.BaseExperience
		chance := 5000 / baseExp
		if rand.Intn(100) < chance {
			fmt.Printf("%s was caught!\n", name)
			cfg.pokedex[name] = infoResp
		} else {
			fmt.Printf("%s escaped!\n", name)
		}
	}
	return nil
}
