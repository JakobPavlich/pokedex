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

type AreaWithEncounters struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetPokemonsInArea(areaName string, cache *internal.Cache) {
	requestURL := "https://pokeapi.co/api/v2/location-area/" + areaName
	body, contain := cache.Get(requestURL)
	if contain {
		area := AreaWithEncounters{}
		err := json.Unmarshal(body, &area)
		if err != nil {
			fmt.Println("Failed to unmarshal AreaWithEncounters")
		}
		encounters := area.PokemonEncounters
		fmt.Printf("Exploring %v...\n", area.Name)
		fmt.Println("Found Pokemon:")
		for _, encounter := range encounters {
			fmt.Printf(" - %v\n", encounter.Pokemon.Name)
		}

	}
	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatalf("Po http.Get(): %v", err)
	}

	body, err = io.ReadAll(res.Body)
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
	area := AreaWithEncounters{}
	err = json.Unmarshal(body, &area)
	if err != nil {
		fmt.Println("Failed to unmarshal AreaWithEncounters")
	}
	encounters := area.PokemonEncounters
	fmt.Printf("Exploring %v...\n", area.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}
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
