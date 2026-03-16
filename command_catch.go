package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("please provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.caughtPokemon[pokemonName] = pokemon

	return nil
}
