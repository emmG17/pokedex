package pokeapi

import (
	"net/http"
	"time"
  "errors"
  "io"
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

func (c *PokemonClient) get(url *string) ([]byte, error) {
  var data []byte
  if val, err := c.cache.Get(*url); err == nil  {
     data = val 
  } else {
    req, err := http.NewRequest("GET", *url, nil)

    if err!= nil {
      return nil, err
    }

    res, err := c.client.Do(req)

    if err!= nil {
      return nil, err
    }

    defer res.Body.Close()

    if res.StatusCode > 399 {
      return nil, errors.New("Response failed with status code " + res.Status) 
    }

    body, err := io.ReadAll(res.Body)

    if err!= nil {
      return nil, err
    }

    data = body
  }
  return data, nil
}
