package main

import (
	"bufio"
	"fmt"
	"os"
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
      desc: "Gets 20 locations",
      cb: mapCommand, 
    },
  }
}

func displayHelp(config *Config) error {  
  commands := getCommands()
  for _, command := range commands {
    fmt.Printf("%v: %v\n", command.name, command.desc)
  }
  return nil
}

func exitCommand(config *Config) error {
  os.Exit(0)
  return nil
}

func mapCommand(config *Config) error {
  locations, err := config.Client.GetLocations(config.Next)
  if err != nil {
    return err
  }
  config.Next = locations.Next
  config.Previous = locations.Previous
  for _, location := range locations.Results {
    fmt.Println(location.Name)
  }
  return nil
}

func repl(config *Config) error {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Pokedex v1")
  fmt.Println("No rights reserved")
  for {
    fmt.Print("pokedex> ")
    scanner.Scan()
    text := scanner.Text()
    
    commands := getCommands()
    command, ok := commands[text]
    if ok {
      command.cb(config)
    } else {
      fmt.Println("Unrecognized command, try again")
    }
  }
}
