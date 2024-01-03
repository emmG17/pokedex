package pokeapi
import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *PokemonClient) GetLocations(url *string) (Locations, error) {
  if url == nil {
    defaultUrl := baseURL + "/location"
    url = &defaultUrl
  }

  var data []byte

  // Check if the URL is in the pokecache
  if val, err := c.cache.Get(*url); err == nil  {
     data = val 
  } else {
    req, err := http.NewRequest("GET", *url, nil)

    if err!= nil {
      return Locations{}, err
    }

    res, err := c.client.Do(req)

    if err != nil {
      return Locations{}, err
    }

    defer res.Body.Close()
  
    if res.StatusCode > 399 {
      return Locations{}, errors.New("Response failed with status code " + res.Status) 
    }

    body, err := io.ReadAll(res.Body)

    if err!= nil {
      return Locations{}, err
    }

    data = body
  }

  var locations Locations
  err := json.Unmarshal(data, &locations)

  if err != nil {
    return Locations{}, err
  }

  c.cache.Add(*url, data)

  return locations, nil
}
