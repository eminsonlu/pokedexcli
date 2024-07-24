package main

import (
	"bufio"
	"fmt"
	"github.com/eminsonlu/pokedexcli/internal/pokeapi"
	"os"
	"time"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	for {
		fmt.Print("pokedex > ")
		inputs := bufio.NewScanner(os.Stdin)
		inputs.Scan()
		cleanedCommand := cleanInput(inputs.Text())

		if len(cleanedCommand) == 0 {
			continue
		}

		command, exist := GetCommands()[cleanedCommand[0]]
		parameters := cleanedCommand[1:]

		if !exist {
			fmt.Println("Command not found")
			continue
		}

		err := command.callback(cfg, parameters)

		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
