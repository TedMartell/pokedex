package main

import (
	"fmt"
	"log"

	"github.com/TedMartell/pokedex/internal/pokeapi"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	// Printing each location area name
	for _, area := range resp.Results {
		fmt.Printf("- %v\n", area.Name)
	}
	return nil
}
