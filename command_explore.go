package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("please provide a location name")
	}

	locationName := args[0]
	locationDetailResponse, err := cfg.pokeapiClient.GetLocation(locationName)

	if err != nil {
		return err
	}

	for _, pokemon := range locationDetailResponse.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
