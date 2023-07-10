package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"pokedex": {
			name:        "inspect {pokemon_name}",
			description: "View all caught Pokemons",
			callback:    callbackPokedex,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Get details of a caught Pokemon",
			callback:    callbackInspect,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Catch the Pokemon",
			callback:    callbackCatch,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "List pokemons in the area",
			callback:    callbackExplore,
		},
		"map": {
			name:        "map",
			description: "List some Location Areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List Previous Location Areas",
			callback:    callbackMapB,
		},
		"help": {
			name:        "help",
			description: "Shows the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Quit Poke cli",
			callback:    callbackExit,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("-> ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		err := command.callback(cfg, args...)

		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
