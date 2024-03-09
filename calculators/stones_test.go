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
		1,
		.5+70*upgrade_data.PerStepGiantLuckImprovement,
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
	sc := NewStonesCalculator(
		userMods,
		219,
		0.43,
		0,
		100,
		MythicEgg,
		false,
	)

	margin := sc.calculateStrikeImprovementMargin(5, time.Hour*24*365)
	assert.Equal(t, "0.005432", fmt.Sprintf("%5f", margin))
	margin = sc.calculateStrikeImprovementMargin(2, time.Hour)
	assert.NotEqual(t, 0.0, margin)
}

func TestFindNextStoneUpgrade(t *testing.T) {
	userMods := NewMiningModifiers(
		0.72,
		1,
		70*upgrade_data.PerStepGiantLuckImprovement,
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
	sc := NewStonesCalculator(
		userMods,
		100,
		0.43,
		0,
		100,
		MythicEgg,
		false,
	)

	result := sc.FindNextUpgrade(1800000, 10000000000)
	assert.Equal(t, "speed", result)

	result = sc.FindNextUpgrade(18000000, 10000000000)
	assert.Equal(t, "speed", result)

	result = sc.FindNextUpgrade(180000000, 10000000000)
	assert.Equal(t, "x2 strike", result)

	result = sc.FindNextUpgrade(180000000, 5)
	assert.Equal(t, "clone luck", result)
}
