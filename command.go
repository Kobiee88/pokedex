package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{}

var currentOffset = 0

func handleCommand(commandName string) error {
	if cmd, exists := commands[commandName]; exists {
		return cmd.callback()
	} else {
		fmt.Println("Unknown command")
		return fmt.Errorf("Unknown command: %s", commandName)
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap() error {
	createApiCall(currentOffset)
	currentOffset += 20
	return nil
}

func commandMapBack() error {
	if currentOffset <= 20 {
		fmt.Println("you're on the first page")
		return nil
	}
	createApiCall(currentOffset - 40)
	currentOffset -= 20
	return nil
}

func initCommandList() {
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Show this help message",
		callback:    commandHelp,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Show available locations on a map",
		callback:    commandMap,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Show available locations on a map",
		callback:    commandMapBack,
	}
}