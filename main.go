package main

import (
	"fmt"
	"bufio"
	"os"	
)

const (
    reset  = "\033[0m"
    red    = "\033[31m"
    cyan  = "\033[32m"
    yellow = "\033[33m"
    blue   = "\033[34m"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}


func main() {
	commands := map[string]cliCommand {
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error {
				fmt.Println("Welcome to the Pokedex!")
				fmt.Println("Usage:")
				return nil	
			},
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func() error {
				os.Exit(0)
				return nil
			},
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := scanner.Text() 
		fmt.Print(cyan + "Pokedex > " + reset)

		switch input {
		case "help": 
			commands["help"].callback()
			for _, cmd := range commands {
				fmt.Printf(cyan+"- %s: "+yellow+"%s\n"+reset, cmd.name, cmd.description)
			}
		case "exit": 
			commands["exit"].callback()
		default:
			fmt.Println(red + "This is not a valid command" + reset)
		}
	}

}