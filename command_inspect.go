package main

import "fmt"

func Inspect(cfg *config, parameters []string) error {
	pokemonName := parameters[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		fmt.Println("You don't have", pokemonName)
		return nil
	}

	fmt.Println("Name:", pokemonName)
	fmt.Println("Base Experience:", pokemon.BaseExperience)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println(" -", t.Type.Name)
	}
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Println(" -", s.Stat.Name, ":", s.BaseStat)
	}

	return nil
}
