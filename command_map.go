package main

import "fmt"

func CommandMap(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.GetLocationList(cfg.nextLocationsURL)

	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func CommandMapBack(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("No previous locations")
		return nil
	}

	locationResp, err := cfg.pokeapiClient.GetLocationList(cfg.prevLocationsURL)

	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}
