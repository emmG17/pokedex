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

  var locations Locations
  err = json.Unmarshal(body, &locations)

  if err != nil {
    return Locations{}, err
  }

  return locations, nil
}
