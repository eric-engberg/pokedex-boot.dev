package main

import (
	"bufio"
	"fmt"
	"os"
	"net/http"
	"encoding/json"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

type Location struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	Next *string
	Previous *string
}

func main() {

	baseURL := "https://pokeapi.co/api/v2/"
	var commands map[string]cliCommand
	var config config

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
		"map": {
			name: "map",
			description: "Get a list of locations",
			callback: func() error {
				return commandMap(&config, baseURL)
			},
		},
		"mapb": {
			name: "mapb",
			description: "Get the previous page of locations",
			callback: func() error {
				return commandMapb(&config)
			},
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

func commandMap(config *config, baseURL string) error {
	url := baseURL + "/location-area"

	if config.Next != nil {
		url = *config.Next
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var locations Location
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)

	if err != nil {
		return err
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	fmt.Println("Locations:")
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(config *config) error {

	if config.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *config.Previous

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var locations Location
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)

	if err != nil {
		return err
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	fmt.Println("Locations:")
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
