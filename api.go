package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
	"time"
	"github.com/kobiee88/pokedex/internal/pokecache"
)

type LocationAreaBatch struct {
    Count    int `json:"count"`
    Next     string `json:"next"`
    Previous string `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

var cache *pokecache.Cache

func initCache() {
	cache = pokecache.NewCache(10 * time.Second)
}

func createApiCall(offset int) {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", offset)
	body := []byte{}
	if cache.Get(fullUrl) {
		fmt.Println("Cache hit!")
		data, _ := cache.Get(fullUrl)
		body = data
	} else {
		res, err := http.Get(fullUrl)
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
		cache.Add(fullUrl, body)
	}
	//fmt.Printf("%s", body)
	locationBatch := LocationAreaBatch{}
    err = json.Unmarshal(body, &locationBatch)
    if err != nil {
        fmt.Println(fmt.Sprintf("ERROR: %s", err))
    }
    for _, loc := range locationBatch.Results {
        fmt.Println(loc.Name)
    }
}