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
	resp, err := http.Get(url)
	if err != nil {
		return p, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&p)
	return p, err
}
