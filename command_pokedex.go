package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("pokemon in pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
