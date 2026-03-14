package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func main() {
	var commands map[string]cliCommand

	commands = map[string]cliCommand{
		"help": {
			name: "help",
			description: "Show available commands",
			callback: func() error {
				return commandHelp(commands)
			},
		},
		"exit": {
			name: "exit",
			description: "Exit the program",
			callback: commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		words := cleanInput(text)

		if len(words) == 0 {
			continue
		}

		cmd := words[0]
		// args := words[1:]

		if command, ok := commands[cmd]; ok {
			err := command.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command:", cmd)
		}

	}
	return
}

func commandHelp(commands map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!\nUsage: \n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
