package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JakobPavlich/pokedex/internal"
	"github.com/JakobPavlich/pokedex/pokeapi"
)

func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		firstWord := cleanText[0]
		var parameter string
		if len(cleanText) >= 2 {
			parameter = cleanText[1]
		}

		// fmt.Println("Your command was:", firstWord)
		supported := c.supportedCommands
		command, ok := supported[firstWord]
		if ok {
			err := command.callback(c, parameter)
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
	callback    func(c *config, parameter string) error
}

type config struct {
	supportedCommands map[string]cliCommand
	nextURL           string
	previousURL       *string
	cache             *internal.Cache
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
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "Displays the names of all Pokemon in that area",
			callback:    commandExplore,
		},
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)

	words := strings.Fields(lowerText)

	return words
}

func commandExit(c *config, parameter string) error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}

func commandHelp(c *config, parameter string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Print("\n")
	supported := c.supportedCommands
	for _, command := range supported {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(c *config, parameter string) error {
	if c.nextURL == "" {
		c.nextURL, c.previousURL = pokeapi.GetAreas("", c.cache)
	} else {
		c.nextURL, c.previousURL = pokeapi.GetAreas(c.nextURL, c.cache)
	}
	return nil
}

func commandMapb(c *config, parameter string) error {
	if c.previousURL == nil {
		// c.nextURL, c.previousURL = pokeapi.GetAreas("")
		fmt.Println("you're on the first page")
	} else {
		c.nextURL, c.previousURL = pokeapi.GetAreas(*c.previousURL, c.cache)
	}
	return nil
}

func commandExplore(c *config, parameter string) error {
	// TODO: pokeapi.GetPokemonInArea(area string)
	// če je parameter "" - fmt.print...
	if parameter == "" {
		fmt.Println("Command explore needs a parameter (area)")
	} else {
		pokeapi.GetPokemonsInArea(parameter, c.cache)
	}
	return nil
}
