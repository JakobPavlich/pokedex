package main

func main() {
	cfg := &config{
		supportedCommands: supportedCommands(),
	}
	startRepl(cfg)
}
