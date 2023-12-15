package pokeapi

import (
	"fmt"
  "log"
	"io"
  "net/http"
  "encoding/json"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations() Locations {
  var locations Locations
  url := fmt.Sprintf("https://pokeapi.co/api/v2/location/")
  res, err := http.Get(url)

  if err != nil {
    log.Fatal(err)
  }

  body, err := io.ReadAll(res.Body)
  res.Body.Close()
  
  if res.StatusCode > 299 {
    log.Fatalf("Response failed with status code %d", res.StatusCode)
  }
  if err!= nil {
    log.Fatal(err)
  }

  err = json.Unmarshal(body, &locations)

  if err!= nil {
    log.Fatal(err)
  }
  return locations
}
