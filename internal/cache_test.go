package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(int(interval))
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestExpectedString(t *testing.T) {
	type Area struct {
		Count    int     `json:"count"`
		Next     string  `json:"next"`
		Previous *string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}
	const interval = 5 * time.Second
	cache := NewCache(int(interval))
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatalf("Po http.Get(): %v", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Couldnt read the body content: %v:", err)
	}
	cache.Add("https://pokeapi.co/api/v2/location-area/", body)
	cachedBody, _ := cache.Get("https://pokeapi.co/api/v2/location-area/")
	firstArea := Area{}
	err = json.Unmarshal(cachedBody, &firstArea)
	if err != nil {
		t.Errorf("First unmarshaling failed")
	}
	comparisonFound := false
	results := firstArea.Results
	for _, result := range results {
		if result.Name == "canalave-city-area" {
			comparisonFound = true
		}
	}
	if comparisonFound == false {
		t.Errorf("Could't find the expected string on the first page")
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(int(baseTime))
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
