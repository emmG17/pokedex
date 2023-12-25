package pokeapi

import (
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

func GetLocations(url string) Locations {
  if url == "" {
    url = "https://pokeapi.co/api/v2/location/"
  }

  var locations Locations
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

  if err != nil {
    log.Fatal(err)
  }
  return locations
}
