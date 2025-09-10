package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(string) error
}

var commands = map[string]cliCommand{}

var pokedex = map[string]Pokemon{}

var currentOffset = 0

func handleCommand(commandName string, name string) error {
	if cmd, exists := commands[commandName]; exists {
		return cmd.callback(name)
	} else {
		fmt.Println("Unknown command")
		return fmt.Errorf("Unknown command: %s", commandName)
	}
}

func commandExit(name string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(name string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(name string) error {
	createApiCall(currentOffset)
	currentOffset += 20
	return nil
}

func commandMapBack(name string) error {
	if currentOffset <= 20 {
		fmt.Println("you're on the first page")
		return nil
	}
	createApiCall(currentOffset - 40)
	currentOffset -= 20
	return nil
}

func commandExplore(name string) error {
	fmt.Printf("Exploring %s...\n", name)
	printPokemonOfLocationArea(name)
	return nil
}

func commandCatch(name string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if currentPokemon, ok := getPokemon(name); !ok {
		return nil
	} else {
		if rand.Float64() > math.Min(0.9, math.Max(0.3, 0.1+float64(currentPokemon.Base_Experience)/1000)) {
			fmt.Printf("%s escaped!\n", currentPokemon.Name)
			return nil
		}
		pokedex[currentPokemon.Name] = currentPokemon
		fmt.Printf("%s was caught!\n", currentPokemon.Name)
		return nil
	}
}

func commandInspect(name string) error {
	if currentPokemon, ok := pokedex[name]; !ok {
		fmt.Printf("you have not caught %s yet\n", name)
		return nil
	} else {
		printPokemonDetails(currentPokemon)
		return nil
	}
}

func commandPokedex(name string) error {
	if len(pokedex) == 0 {
		fmt.Println("you have not caught any Pokemon yet")
		return nil
	}
	fmt.Println("You have caught the following Pokemon:")
	for _, p := range pokedex {
		fmt.Printf("- %s\n", p.Name)
	}
	return nil
}

func printPokemonDetails(pokemon Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Types: ")
	for i, t := range pokemon.Types {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%s", t.Type.Name)
	}
	fmt.Printf("\n")
	fmt.Printf("Stats:\n")
	for _, s := range pokemon.Stats {
		fmt.Printf("  %s: %d\n", s.Stat.Name, s.Base_Stat)
	}
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
	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Explore a location - provide the location name as an argument",
		callback:    commandExplore,
	}
	commands["catch"] = cliCommand{
		name:        "catch",
		description: "Catch a Pokemon - provide the Pokemon name as an argument",
		callback:    commandCatch,
	}
	commands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Inspect a caught Pokemon - provide the Pokemon name as an argument",
		callback:    commandInspect,
	}
	commands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "List all caught Pokemon",
		callback:    commandPokedex,
	}
}
