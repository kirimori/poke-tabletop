package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

// GetPokemon retrieves a pokemon by name
func GetPokemon(name string) (p Pokemon, err error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, name)
	resp, err := http.Get(url)
	if err != nil {
		return p, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&p)
	return p, err
}
