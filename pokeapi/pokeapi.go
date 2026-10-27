package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/JakobPavlich/pokedex/internal"
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

func GetAreas(pageURL string, cache *internal.Cache) (string, *string) {
	requestURL := "https://pokeapi.co/api/v2/location-area/"
	if pageURL != "" {
		requestURL = pageURL
	}
	cacheBody, contain := cache.Get(requestURL)
	if contain {
		area := Area{}
		err := json.Unmarshal(cacheBody, &area)
		if err != nil {
			fmt.Println(err)
		}
		results := area.Results
		for _, result := range results {
			fmt.Println(result.Name)
		}
		return area.Next, area.Previous
	}
	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatalf("Po http.Get(): %v", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Couldnt read the body content: %v:", err)
	}
	cache.Add(requestURL, body)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	err = res.Body.Close()
	if err != nil {
		fmt.Printf("Couldnt close the response body: %v", err)
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
