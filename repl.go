package main


import (
	"bufio"
	"fmt"
	"os"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	cyan  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
)

func startRepl() {
	commands := getCommands()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(cyan + "Pokedex > " + reset)
		reader.Scan()
		input := reader.Text()
		switch input {
		case "help":
			commands["help"].callback()
		case "exit":
			commands["exit"].callback()
		default:
			fmt.Println(red + "Unknown command" + reset)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
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
	
