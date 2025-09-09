package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	//fmt.Println("Hello, World!")
	reader := bufio.NewReader(os.Stdin)
	initCommandList()
	initCache()
	//createApiCall(0)
	for true {
		fmt.Print("Pokedex > ")
		text, _ := reader.ReadString('\n')

		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		handleCommand(commandName)

		//fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {
	var words []string
	words = strings.Fields(strings.ToLower(text))
	return words
}
	
