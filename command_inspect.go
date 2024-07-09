package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("pokemon have not been caught")
	}
	fmt.Printf("name: %s\n", pokemon.Name)
	fmt.Printf("height: %v\n", pokemon.Height)
	fmt.Printf("weight: %v\n", pokemon.Weight)
	fmt.Println("stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s : %v \n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("types:")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s", typ.Type.Name)
	}

	return nil
}
