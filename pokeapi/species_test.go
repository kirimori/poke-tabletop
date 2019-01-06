package pokeapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSpeciesLinkGet(t *testing.T) {
	sl := SpeciesLink{
		URL: "https://pokeapi.co/api/v2/pokemon-species/25/",
	}
	species, err := sl.Get()
	require.NoError(t, err)
	assert.Equal(t, 25, species.ID)
}
