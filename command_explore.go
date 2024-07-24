package main

import "fmt"

func Explore(cfg *config, parameters []string) error {
	location := parameters[0]

	locationResp, err := cfg.pokeapiClient.GetLocationExplore(location)
	if err != nil {
		return err
	}

	fmt.Println("Exploring:", location)
	fmt.Println("Founded Pokemon:")

	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Println(" -", pokemon.Pokemon.Name)
	}

	return nil
}
