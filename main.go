package main

import "github.com/emmG17/pokedex/pokeapi"

type Config struct {
  Previous *string
  Next *string
  Client pokeapi.PokemonClient
}

func main() {
  config := Config{
    Previous: nil,
    Next: nil,
    Client: pokeapi.NewPokemonClient(5),
  }

  repl(&config)  
}
