package pokeapi

import (
	"net/http"
	"time"
)

type PokemonClient struct {
  client http.Client
}

const baseURL = "https://pokeapi.co/api/v2/"

func NewPokemonClient() PokemonClient {
  return PokemonClient{
    client: http.Client{
      Timeout: time.Minute,
    },
  } 
}
