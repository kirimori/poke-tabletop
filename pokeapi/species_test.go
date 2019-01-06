package pokeapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSpeciesLinkGet(t *testing.T) {
	tt := []struct {
		name string

		expectedID int
		expectErr  bool
		url        string
	}{
		{
			url:        "https://pokeapi.co/api/v2/pokemon-species/25/",
			expectedID: 25,
			name:       "as seen in pokemon return value",
		}, {
			url:        "https://pokeapi.co/api/v2/pokemon-species/25",
			expectedID: 25,
			name:       "without trailing slash",
		}, {
			url:       "",
			expectErr: true,
			name:      "empty url",
		}, {
			url:       "https://pokeapi.co/api/v2/pokemon-species/twenty/",
			expectErr: true,
			name:      "failed string conversion",
		}, {
			url:       "https://pokeapi.co/api/pokemon-species/25/",
			expectErr: true,
			name:      "bad base url",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			species, err := SpeciesLink{
				URL: test.url,
			}.Get()

			if test.expectErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expectedID, species.ID)
			}
		})
	}
}
