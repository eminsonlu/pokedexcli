package main

import "fmt"

func Pokedex(cfg *config, parameters []string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You don't have any pokemon")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.caughtPokemon {
		fmt.Println(" -", name)
	}

	return nil
}
