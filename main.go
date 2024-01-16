package main

import "github.com/emmG17/pokedex/pokeapi"

type Config struct {
  Previous *string
  Next *string
  SelectedArea *string
  SelectedPokemon *string
  Client pokeapi.PokemonClient
}

var Pokedex = make(map[string]pokeapi.Pokemon )

func main() {
  config := Config{
    Previous: nil,
    Next: nil,
    Client: pokeapi.NewPokemonClient(5),
  }

  repl(&config)  
}
