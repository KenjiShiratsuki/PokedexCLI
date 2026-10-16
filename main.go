package main

func main() {
	cfg := &config{}
	cfg.commandList = getCommands()
	startRepl(cfg)
}
