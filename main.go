package main

import "github.com/JakobPavlich/pokedex/internal"

func main() {
	cfg := &config{
		supportedCommands: supportedCommands(),
		cache:             internal.NewCache(5),
	}
	startRepl(cfg)
	// pokeapi.GetAreas()
	// next, prev := pokeapi.GetPrevAndNextURL()
	// cfg.nextUrl = next
	// cfg.previousUrl = prev
	//
	// fmt.Println(cfg.nextUrl, cfg.previousUrl)
}
