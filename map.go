package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type config struct {
	Next     string
	Previous string
}

type locationArea struct {
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
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func getLocation(c *config) error {
	{
		temp, _ := strings.CutPrefix(c.Next, "https://pokeapi.co/api/v2/location-area/")
		index, _ := strconv.Atoi(temp)
		if index > 20 {
			index -= 20
			temp = strconv.Itoa(index)
			c.Previous = "https://pokeapi.co/api/v2/location-area/" + temp
		} else {
			c.Previous = ""
		}
	}
	for range 20 {
		res, err := http.Get(c.Next)
		if err != nil {
			return err
		}

		if res.StatusCode == 404 {
			temp, _ := strings.CutPrefix(c.Next, "https://pokeapi.co/api/v2/location-area/")
			index, _ := strconv.Atoi(temp)
			index++
			temp = strconv.Itoa(index)
			c.Next = "https://pokeapi.co/api/v2/location-area/" + temp
			continue
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		res.Body.Close()

		var location locationArea
		err = json.Unmarshal(data, &location)
		if err != nil {
			return err
		}

		fmt.Println(location.Name)
		temp, _ := strings.CutPrefix(c.Next, "https://pokeapi.co/api/v2/location-area/")
		index, _ := strconv.Atoi(temp)
		index++
		temp = strconv.Itoa(index)
		c.Next = "https://pokeapi.co/api/v2/location-area/" + temp
	}
	return nil
}

func getLocationb(c *config) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
	} else {
		c.Next = c.Previous
		getLocation(c)
	}
	return nil
}
