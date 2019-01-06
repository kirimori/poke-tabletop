package rpg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ref: https://www.d20pfsrd.com/basics-ability-scores/ability-scores/
var stattests = []struct {
	givenScore  int
	expectedMod int
}{
	{1, -5},
	{2, -4},
	{3, -4},
	{4, -3},
	{5, -3},
	{6, -2},
	{7, -2},
	{8, -1},
	{9, -1},
	{10, 0},
	{11, 0},
	{12, 1},
	{13, 1},
	{14, 2},
	{15, 2},
	{16, 3},
	{17, 3},
	{18, 4},
	{19, 4},
	{20, 5},
	{21, 5},
	{22, 6},
	{23, 6},
	{24, 7},
	{25, 7},
	{26, 8},
	{27, 8},
	{28, 9},
	{29, 9},
	{30, 10},
	{31, 10},
	{32, 11},
	{33, 11},
	{34, 12},
	{35, 12},
	{36, 13},
	{37, 13},
	{38, 14},
	{39, 14},
	{40, 15},
	{41, 15},
	{42, 16},
	{43, 16},
	{44, 17},
	{45, 17},
}

func TestStatCalculation(t *testing.T) {
	for _, tt := range stattests {
		name := fmt.Sprintf("ability score %d should have modifier %d", tt.givenScore, tt.expectedMod)
		t.Run(name, func(t *testing.T) {
			stat := Stat{Score: tt.givenScore}
			assert.Equal(t, tt.expectedMod, stat.Modifier())
		})
	}
}
