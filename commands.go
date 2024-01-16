package main

import (
  "fmt"
  "os"
)

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

func mapBCommand(config *Config) error {
  locations, err := config.Client.GetLocations(config.Previous)
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

func explore(config *Config) error {
  if config.SelectedArea == nil {
    fmt.Println("No area selected")
    return nil
  }

  locationDetails, err := config.Client.GetLocation(config.SelectedArea)

  if err!= nil {
    return err
  }
  
  for _, pokemon := range locationDetails.PokemonEncounters {
    fmt.Println(pokemon.Pokemon.Name)
  }
  return nil
}

func catch(config *Config) error {
  if config.SelectedPokemon == nil {
    fmt.Println("No pokemon selected")
    return nil
  }

  pokemon, err := config.Client.GetPokemon(config.SelectedPokemon)
   
  if err!= nil {
    return err
  }
  fmt.Println("Throwing pokeball...")
  if catched := CatchPokemon(pokemon.BaseExperience); catched {
    Pokedex[pokemon.Name] = pokemon
    fmt.Printf("%s caught!\n", pokemon.Name)
  } else {
    fmt.Printf("%s has escaped!\n", pokemon.Name)
  }

  return nil
}

  


