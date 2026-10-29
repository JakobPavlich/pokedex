package main

import (
	"github.com/JakobPavlich/pokedex/internal"
	"github.com/JakobPavlich/pokedex/pokeapi"
)

func main() {
	cfg := &config{
		supportedCommands: supportedCommands(),
		cache:             internal.NewCache(5),
		pokedex:           make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
	// pokeapi.GetAreas()
	// next, prev := pokeapi.GetPrevAndNextURL()
	// cfg.nextUrl = next
	// cfg.previousUrl = prev
	//
	// fmt.Println(cfg.nextUrl, cfg.previousUrl)
}
