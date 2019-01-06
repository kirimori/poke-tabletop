package pokeapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPokemon(t *testing.T) {
	name := "pikachu"
	pokemon, err := GetPokemon(name)
	require.NoError(t, err)
	assert.Equal(t, name, pokemon.Name)

	name = "missingno"
	pokemon, err = GetPokemon(name)
	assert.Error(t, err)

	name = ""
	pokemon, err = GetPokemon(name)
	assert.Error(t, err)
}
