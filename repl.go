package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
  name string
  desc string
  cb func() error
}

var commands map[string]cliCommand = map[string]cliCommand{
"help": {
    name: "help",
    desc: "Prints this message",
    cb: func() error {
      fmt.Println("Commands:\nexit - Exits the program\nhelp - Prints this message")
      return nil
    },
  },
  "exit": {
    name: "exit",
    desc: "Exits the program",
    cb: func() error {
      os.Exit(0) 
      return nil
    },
  },
}

func repl() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Pokedex v1")
  fmt.Println("No rights reserved")
  for ;; {
    fmt.Print("> ")
    scanner.Scan()
    text := scanner.Text()
    
    command, ok := commands[text]
    if ok {
      command.cb()
    } else {
      fmt.Println(text)
    }
  }
}
