package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
  name string
  desc string
  cb func(*Config) error
}

func getCommands() map[string]cliCommand {
  return map[string]cliCommand{
    "help": {
      name: "help",
      desc: "Prints this message",
      cb: displayHelp,
    },
    "exit": {
      name: "exit",
      desc: "Exits the program",
      cb: exitCommand, 
    },
    "map": {
      name: "map",
      desc: "Goes 20 locations forward",
      cb: mapCommand, 
    },
    "mapb": {
      name: "mapb",
      desc: "Goes 20 locations backwards",
      cb: mapBCommand, 
    },
    "explore": {
      name: "explore",
      desc: "Explore a location",
      cb: explore, 
    },
    "catch": {
      name: "catch",
      desc: "Catch a pokemon",
      cb: catch, 
    },
  }
}

func repl(config *Config) error {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Pokedex v1")
  fmt.Println("No rights reserved")
  for {
    fmt.Print("pokedex> ")
    scanner.Scan()
    text := scanner.Text()

    input := strings.Split(text, " ")
    
    commands := getCommands()
    command, ok := commands[input[0]]
    if ok {
      if command.name == "explore" {
        config.SelectedArea = &input[1]
      }
      if command.name == "catch" {
        config.SelectedPokemon = &input[1]
      }
      command.cb(config)
    } else {
      fmt.Println("Unrecognized command, try again")
    }
  }
}
