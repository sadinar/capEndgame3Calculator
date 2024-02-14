package calculators

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCalculateStrikeImprovementMargin(t *testing.T) {
	ocConfig := NewOverclockConfig(true, true, false, true, true, true)
	userMods := NewUserModifiers(
		1.1,
		1.2,
		0.72,
		1,
		map[int]int{
			2: 74,
			3: 74,
			4: 74,
			5: 74,
		},
		70,
	)
	sc := NewStonesCalculator(
		userMods,
		RubyPick,
		0.43,
		MythicEgg,
		ocConfig,
	)

	assert.Equal(t, 0.185, sc.x3Strike)
	margin := sc.calculateStrikeImprovementMargin(5, time.Hour)
	assert.Equal(t, 0.0, margin)
	margin = sc.calculateStrikeImprovementMargin(2, time.Hour)
	assert.NotEqual(t, 0.0, margin)
	assert.Equal(t, 0.185, sc.x3Strike)
}

func TestFindNextStoneUpgrade(t *testing.T) {
	ocConfig := NewOverclockConfig(true, true, false, true, true, true)
	userMods := NewUserModifiers(
		1.1,
		1.2,
		0.72,
		1,
		map[int]int{
			2: 74,
			3: 74,
			4: 74,
			5: 74,
		},
		70,
	)
	sc := NewStonesCalculator(
		userMods,
		RubyPick,
		0.43,
		MythicEgg,
		ocConfig,
	)

	result := sc.FindNextUpgrade(1800000)
	assert.Equal(t, "next stone upgrade should be speed", result)

	result = sc.FindNextUpgrade(18000000)
	assert.Equal(t, "next stone upgrade should be speed", result)

	result = sc.FindNextUpgrade(180000000)
	assert.Equal(t, "next stone upgrade should be x2 strike", result)
}
