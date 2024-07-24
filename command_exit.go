package main

import "os"

func Exit(cfg *config, parameters []string) error {
	os.Exit(0)
	return nil
}
