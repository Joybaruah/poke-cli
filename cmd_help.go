package main

import "fmt"

func callbackHelp(_ *config, _ ...string) error {
	fmt.Println("Available Commands:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
