package main

func main() {
	cfg := &config{
		supportedCommands: supportedCommands(),
	}
	startRepl(cfg)
	// pokeapi.GetAreas()
	// next, prev := pokeapi.GetPrevAndNextURL()
	// cfg.nextUrl = next
	// cfg.previousUrl = prev
	//
	// fmt.Println(cfg.nextUrl, cfg.previousUrl)
}
