package pokeapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvolutionChainLinkGet(t *testing.T) {
	tt := []struct {
		name string

		expectErr bool
		url       string
	}{
		{
			url:  "https://pokeapi.co/api/v2/evolution-chain/10/",
			name: "as seen in pokemon return value",
		}, {
			url:  "https://pokeapi.co/api/v2/evolution-chain/10",
			name: "without trailing slash",
		}, {
			url:       "",
			expectErr: true,
			name:      "empty url",
		}, {
			url:       "https://pokeapi.co/api/v2/evolution-chain/ten/",
			expectErr: true,
			name:      "failed string conversion",
		}, {
			url:       "https://pokeapi.co/api/evolution-chain/10/",
			expectErr: true,
			name:      "bad base url",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			ecl := EvolutionChainLink{
				Link: Link{
					URL: test.url,
				},
			}
			ec, err := ecl.Get()

			if test.expectErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, 10, ec.ID)

				assert.Equal(t, "pichu", ec.Chain.Species.Name)

				require.Equal(t, 1, len(ec.Chain.EvolvesTo))
				pikachuChain := ec.Chain.EvolvesTo[0]
				assert.Equal(t, "pikachu", pikachuChain.Species.Name)

				require.Equal(t, 1, len(pikachuChain.EvolvesTo))
				assert.Equal(t, "raichu", pikachuChain.EvolvesTo[0].Species.Name)
			}
		})
	}
}
