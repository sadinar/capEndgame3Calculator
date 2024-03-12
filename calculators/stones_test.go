package calculators

import (
	"capEndgame3Calculator/upgrade_data"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCalculateStrikeImprovementMargin(t *testing.T) {
	userMods := NewMiningModifiers(
		0.72,
		100,
		.5+70*upgrade_data.PerStepGiantLuckImprovement,
		0,
		map[int]int{
			2: 74,
			3: 74,
			4: 74,
			5: 74,
		},
		70,
		map[int]float64{
			2: 100,
			3: 75,
			4: 50,
			5: 25,
		},
	)
	genMods := NewEggGenerationModifiers(0.43, 0, 100, MythicEgg, false)
	sc := NewStonesCalculator(userMods, genMods)

	margin := sc.calculateStrikeImprovementMargin(5, time.Hour*24*365)
	assert.Equal(t, "0.000472", fmt.Sprintf("%5f", margin))
	margin = sc.calculateStrikeImprovementMargin(2, time.Hour*24)
	assert.NotEqual(t, 0.0, margin)
}

func TestFindNextStoneUpgrade(t *testing.T) {
	userMods := NewMiningModifiers(
		0.72,
		100,
		70*upgrade_data.PerStepGiantLuckImprovement,
		0,
		map[int]int{
			2: 74,
			3: 74,
			4: 74,
			5: 74,
		},
		70,
		map[int]float64{
			2: 74 * upgrade_data.PerStepStrikeImprovement,
			3: 74 * upgrade_data.PerStepStrikeImprovement,
			4: 74 * upgrade_data.PerStepStrikeImprovement,
			5: 74 * upgrade_data.PerStepStrikeImprovement,
		},
	)
	genMods := NewEggGenerationModifiers(0.43, 0, 100, MythicEgg, false)
	sc := NewStonesCalculator(userMods, genMods)

	result := sc.FindNextUpgrade(1800000, 10000000000)
	assert.Equal(t, "speed", result)

	result = sc.FindNextUpgrade(18000000, 10000000000)
	assert.Equal(t, "speed", result)

	result = sc.FindNextUpgrade(180000000, 10000000000)
	assert.Equal(t, "x2 strike", result)

	result = sc.FindNextUpgrade(180000000, 5)
	assert.Equal(t, "clone luck", result)
}
