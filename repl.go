package main

import (
	"bufio"
	"fmt"
	"os"
  "github.com/emmG17/pokedex/pokeapi"
)

type cliCommand struct {
  name string
  desc string
  cb func() error
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
      cb: func() error {
        os.Exit(0) 
        return nil
      },
    },
    "map": {
      name: "map",
      desc: "Gets 20 locations",
      cb: func() error {
        locations := pokeapi.GetLocations("")
        for _, location := range locations.Results {
          fmt.Println(location.Name)
        }
        return nil
      },
    },
  }
}

func displayHelp() error {  
  commands := getCommands()
  for _, command := range commands {
    fmt.Printf("%v: %v\n", command.name, command.desc)
  }
  return nil
}

func repl() {
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
      command.cb()
    } else {
      fmt.Println("Unrecognized command, try again")
    }
  }
}
