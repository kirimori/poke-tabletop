package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

// GetPokemon retrieves a pokemon by name
func GetPokemon(name string) (p Pokemon, err error) {
	if len(name) == 0 {
		return p, errors.New("No name given")
	}
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, name)
	err = getFromURL(url, &p)
	return p, err
}

// GetSpecies retrieves a pokemon species
func GetSpecies(number int) (s Species, err error) {
	url := fmt.Sprintf("%s/pokemon-species/%d", baseURL, number)
	err = getFromURL(url, &s)
	return s, err
}

func getFromURL(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected response code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
