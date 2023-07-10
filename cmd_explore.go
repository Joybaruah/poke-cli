package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("location area not provided")
	}

	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.GetLocation(&locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", area.Pokemon.Name)
	}

	return nil
}
