package pokeapi

import (
	"net/http"
	"time"
  "github.com/emmG17/pokedex/internal/pokecache"
)

type PokemonClient struct {
  client http.Client
  cache  pokecache.Cache
}

const baseURL = "https://pokeapi.co/api/v2/"

func NewPokemonClient(interval int) PokemonClient {
  return PokemonClient{
    client: http.Client{
      Timeout: time.Minute,
    },
    cache: pokecache.NewCache(interval),
  } 
}
