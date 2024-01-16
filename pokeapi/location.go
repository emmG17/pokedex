package pokeapi
import (
	"encoding/json"
)

func (c *PokemonClient) GetLocations(url *string) (Locations, error) {
  if url == nil {
    defaultUrl := baseURL + "/location-area"
    url = &defaultUrl
  }

  data, err := c.get(url)

  if err != nil {
    return Locations{}, err
  }

  var locations Locations
  err = json.Unmarshal(data, &locations)

  if err != nil {
    return Locations{}, err
  }

  c.cache.Add(*url, data)

  return locations, nil
}


func (c *PokemonClient) GetLocation(locationName *string) (Location, error) {
  
  url := baseURL + "/location-area/" + *locationName

  data, err := c.get(&url)

  if err!= nil {
    return Location{}, err
  }

  var location Location
  err = json.Unmarshal(data, &location)

  if err!= nil {
    return Location{}, err
  }

  c.cache.Add(url, data)

  return location, nil
}
