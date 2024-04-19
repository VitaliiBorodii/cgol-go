package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLufeRun(t *testing.T) {
	gliderSeed := patterns["glider"]
	t.Parallel()

	testCases := []struct {
		name     string
		seed     [][]uint8
		xSize    int
		ySize    int
		cycles   int
		state    []string
		overflow bool
	}{
		{
			name:   "glider_25x25_10_cylces",
			seed:   gliderSeed,
			xSize:  25,
			ySize:  25,
			cycles: 10,
			state:  []string{"3x3", "3x5", "4x4", "4x5", "5x4"},
		},
		{
			name:   "glider_25x25_90_cycles",
			seed:   gliderSeed,
			xSize:  25,
			ySize:  25,
			cycles: 90,
			state:  []string{"23x23", "23x24", "24x23", "24x24"},
		},
		{
			name:   "glider_25x25_100_cycles",
			seed:   gliderSeed,
			xSize:  25,
			ySize:  25,
			cycles: 100,
			state:  []string{"23x23", "23x24", "24x23", "24x24"},
		},
		{
			name:     "glider_25x25_90_cycles_overflow",
			seed:     gliderSeed,
			xSize:    25,
			ySize:    25,
			cycles:   90,
			overflow: true,
			state:    []string{"23x23", "23x0", "0x24", "24x0", "24x24"},
		},
		{
			name:     "glider_25x25_100_cycles_overflow",
			seed:     gliderSeed,
			xSize:    25,
			ySize:    25,
			cycles:   100,
			overflow: true,
			state:    []string{"0x1", "1x2", "1x3", "2x1", "2x2"},
		},
		{
			name:   "glider_100x100_100_cycles",
			seed:   gliderSeed,
			xSize:  100,
			ySize:  100,
			cycles: 100,
			state:  []string{"25x26", "26x27", "26x28", "27x26", "27x27"},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			l := NewLife(tc.seed, tc.xSize, tc.ySize)
			l.WithOverFlow(tc.overflow)
			l.disablePrintingState()
			l.Run(tc.cycles)
			assert.ElementsMatch(t, tc.state, l.getAliveCells())
		})
	}
}
