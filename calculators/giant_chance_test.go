package calculators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNextUpgrade(t *testing.T) {
	miningMods := NewMiningModifiers(
		1.00+.5, // exactly as on stats screen
		100,     // exactly as shown on the wooden board behind egg
		.151,    // exactly as on stats screen
		408.8,   // exactly as on stats screen
		map[int]int{
			2: 74,
			3: 74,
			4: 74,
			5: 75,
		},
		70,
		map[int]float64{
			2: 29.4,  // exactly as on stats screen
			3: 8.702, // exactly as on stats screen
			4: 2.898, // exactly as on stats screen
			5: 1.087, // exactly as on stats screen
		},
		true,
		true,
		true,
		true,
	)
	LabMods := NewGiantModifiers(1, 1, 1.1, 1.2, true, true)

	gc := NewGiantCalculator(miningMods, LabMods)

	assert.Equal(t, "x3 strike", gc.GetNextUpgrade(2500000))

	gc.miningModifiers.StrikeUpgrades[TripleStrike] = 75
	assert.Equal(t, "x4 strike", gc.GetNextUpgrade(2500000))
	assert.Equal(t, "speed", gc.GetNextUpgrade(250000))

	gc.miningModifiers.StrikeUpgrades[DoubleStrike] = 72
	assert.Equal(t, "x2 strike", gc.GetNextUpgrade(2500000))
}
