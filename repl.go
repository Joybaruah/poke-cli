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
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		command.callback()

	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
