package main

import "fmt"

func Help(cfg *config, parameters []string) error {
	for _, command := range GetCommands() {
		fmt.Println(command.name, "-", command.description)
	}
	return nil
}
