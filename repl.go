package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		firstWord := cleanText[0]
		// fmt.Println("Your command was:", firstWord)
		supported := c.supportedCommands
		command, ok := supported[firstWord]
		if ok {
			err := command.callback(c)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	supportedCommands map[string]cliCommand
}

func supportedCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)

	words := strings.Fields(lowerText)

	return words
}

func commandExit(c *config) error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Print("\n")
	supported := c.supportedCommands
	for _, command := range supported {
		fmt.Println(command.name, ":", command.description)
	}
	return nil
}
