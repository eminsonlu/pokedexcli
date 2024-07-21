package main

import (
	"strings"
)

type Command struct {
	name        string
	description string
	callback    func(*config) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"exit": Command{
			name:        "exit",
			description: "Exit the program",
			callback:    Exit,
		},
		"help": Command{
			name:        "help",
			description: "Show all available commands",
			callback:    Help,
		},
		"map": Command{
			name:        "map",
			description: "Show the map",
			callback:    CommandMap,
		},
		"map-back": Command{
			name:        "map-back",
			description: "Show the previous map",
			callback:    CommandMapBack,
		},
	}
}

func cleanInput(input string) []string {
	ni := strings.ToLower(input)
	words := strings.Fields(ni)
	return words
}
