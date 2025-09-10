package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/kobiee88/pokecache"
)

type LocationAreaBatch struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationArea struct {
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
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			EncounterDetails []struct {
				MinLevel        int `json:"min_level"`
				MaxLevel        int `json:"max_level"`
				ConditionValues []struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"condition_values"`
				Chance int `json:"chance"`
				Method struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

var cache *pokecache.Cache

func initCache() {
	cache = pokecache.NewCache(10 * time.Second)
}

func createApiCall(offset int) {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", offset)
	var myBody []byte
	if data, ok := cache.Get(fullUrl); ok {
		//fmt.Println("<<< Cache hit >>>")
		myBody = data
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
		myBody = body
	}
	//fmt.Printf("%s", body)
	locationBatch := LocationAreaBatch{}
	err := json.Unmarshal(myBody, &locationBatch)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}
	for _, loc := range locationBatch.Results {
		fmt.Println(loc.Name)
	}
}

func printPokemonOfLocationArea(name string) {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)
	var myBody []byte
	if data, ok := cache.Get(fullUrl); ok {
		myBody = data
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
		myBody = body
	}
	locationArea := LocationArea{}
	err := json.Unmarshal(myBody, &locationArea)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}
	fmt.Printf("Pokemon in %s:\n", locationArea.Name)
	for _, poke := range locationArea.PokemonEncounters {
		fmt.Println(poke.Pokemon.Name)
	}
}
