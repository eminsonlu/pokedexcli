package main

import (
	"fmt"
	"math/rand"
)

func Catch(cfg *config, parameters []string) error {
	pokemonName := parameters[0]
	pokemon, err := cfg.pokeapiClient.Pokemon(pokemonName)
	if err != nil {
		return err
	}

	num := rand.Intn(pokemon.BaseExperience)
	if num < 25 {
		cfg.caughtPokemon[pokemonName] = pokemon
		fmt.Println("You caught", pokemonName)
	} else {
		fmt.Println("You failed to catch", pokemonName)
	}

	return nil
}
