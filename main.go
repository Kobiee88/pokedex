package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//fmt.Println("Hello, World!")
	reader := bufio.NewReader(os.Stdin)
	initCommandList()
	initCache()
	//createApiCall(0)
	for {
		fmt.Print("Pokedex > ")
		text, _ := reader.ReadString('\n')

		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		argument := ""
		if len(words) > 1 {
			argument = strings.Join(words[1:], " ")
		}

		//TODO: pass the rest of the words as arguments to the command
		handleCommand(commandName, argument)

		//fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {
	//var words []string
	words := strings.Fields(strings.ToLower(text))
	return words
}
