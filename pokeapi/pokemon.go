package pokeapi
import (
  "encoding/json"
)

func (c *PokemonClient) GetPokemon(pokemonName *string) (Pokemon, error) {
  url := baseURL + "/pokemon/" + *pokemonName 
  data, err := c.get(&url)

  if err!= nil {
    return Pokemon{}, err
  }

  var pokemon Pokemon
  err = json.Unmarshal(data, &pokemon)

  if err!= nil {
    return Pokemon{}, err
  }

  c.cache.Add(url, data)

  return pokemon, nil
}
