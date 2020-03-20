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

func TestGetSpecies(t *testing.T) {
	number := 25
	species, err := GetSpecies(number)
	require.NoError(t, err)
	assert.Equal(t, number, species.ID)

	number = 0
	species, err = GetSpecies(number)
	require.Error(t, err)
}

func TestPokemonLinkIntegration(t *testing.T) {
	name := "pikachu"
	pokemon, err := GetPokemon(name)
	require.NoError(t, err)

	species, err := pokemon.Species.Get()
	require.NoError(t, err)
	assert.Equal(t, name, species.Name)

	evChain, err := species.EvolutionChain.Get()
	require.NoError(t, err)

	pichuSpecies, err := species.EvolvesFromSpecies.Get()
	require.NoError(t, err)
	otherEVChain, err := pichuSpecies.EvolutionChain.Get()
	require.NoError(t, err)

	assert.Equal(t, evChain, otherEVChain)
}
