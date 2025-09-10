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

type Pokemon struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Base_Experience int    `json:"base_experience"`
	Height          int    `json:"height"`
	Is_Default      bool   `json:"is_default"`
	Order           int    `json:"order"`
	Weight          int    `json:"weight"`
	Abilities       []struct {
		Is_Hidden bool `json:"is_hidden"`
		Slot      int  `json:"slot"`
		Ability   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	Game_Indices []struct {
		Game_Index int `json:"game_index"`
		Version    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Held_Items               []interface{} `json:"held_items"`
	Location_Area_Encounters string        `json:"location_area_encounters"`
	Moves                    []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		Version_Group_Details []struct {
			Level_Learned_At int `json:"level_learned_at"`
			Version_Group    struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			Move_Learn_Method struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Past_Types     []interface{} `json:"past_types"`
	Past_Abilities []interface{} `json:"past_abilities"`
	Species        struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		Back_Default       string `json:"back_default"`
		Back_Shiny         string `json:"back_shiny"`
		Front_Default      string `json:"front_default"`
		Front_Shiny        string `json:"front_shiny"`
		Front_Female       string `json:"front_female"`
		Front_Shiny_Female string `json:"front_shiny_female"`
		Back_Female        string `json:"back_female"`
		Back_Shiny_Female  string `json:"back_shiny_female"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		Base_Stat int `json:"base_stat"`
		Effort    int `json:"effort"`
		Stat      struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
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

func getPokemon(name string) (Pokemon, bool) {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	pokemon := Pokemon{}
	res, err := http.Get(fullUrl)
	if err != nil {
		fmt.Printf("%s is not a valid Pokemon name\n", name)
		return pokemon, false
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("%s is not a valid Pokemon name\n", name)
		return pokemon, false
	}
	if err != nil {
		fmt.Printf("%s is not a valid Pokemon name\n", name)
		return pokemon, false
	}
	cache.Add(fullUrl, body)
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return pokemon, false
	}
	return pokemon, true
}
