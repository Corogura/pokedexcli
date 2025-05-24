package main

import "fmt"

func commandInspect(cfg *config, names ...string) error {
	for _, name := range names {
		pkmn, exists := cfg.pokedex[name]
		if !exists {
			return fmt.Errorf("%s is not found in the Pokedex", name)
		} else {
			fmt.Printf("Name: %s\n", pkmn.Name)
			fmt.Printf("Height: %v\n", pkmn.Height)
			fmt.Printf("Weight: %v\n", pkmn.Weight)
			fmt.Println("Stats:")
			for _, stat := range pkmn.Stats {
				fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
			}
			fmt.Println("Types:")
			for _, t := range pkmn.Types {
				fmt.Printf(" - %s\n", t.Type.Name)
			}
			fmt.Println("Abilities:")
			for _, ability := range pkmn.Abilities {
				fmt.Printf(" - %s\n", ability.Ability.Name)
			}
		}
	}
	return nil
}

func commandPokedex(cfg *config, s ...string) error {
	if cfg.pokedex == nil {
		return fmt.Errorf("no Pokemon found in the Pokedex")
	}
	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.pokedex {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
