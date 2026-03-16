package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {

	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.Next)

	if err != nil {
		return err
	}

	cfg.Next = locationsResponse.Next
	cfg.Previous = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {

	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}

	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = locationsResponse.Next
	cfg.Previous = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
