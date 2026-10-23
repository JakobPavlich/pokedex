package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Area struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetAreas(pageURL string) (string, *string) {
	requestURL := "https://pokeapi.co/api/v2/location-area/"
	if pageURL != "" {
		requestURL = pageURL
	}

	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	area := Area{}
	err = json.Unmarshal(body, &area)
	if err != nil {
		fmt.Println(err)
	}
	results := area.Results
	for _, result := range results {
		fmt.Println(result.Name)
	}
	return area.Next, area.Previous
}
