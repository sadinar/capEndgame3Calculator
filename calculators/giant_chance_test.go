package calculators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNextUpgrade(t *testing.T) {
	sl := getCalculatorWithDummyPrices()
	assert.Equal(t, 2, sl.findCartUpgrade())

	sl.miningModifiers.StrikeUpgrades[2] = 1
	sl.miningModifiers.StrikeOdds[2] = 0.0025
	assert.Equal(t, 3, sl.findCartUpgrade())

	sl.miningModifiers.StrikeUpgrades[3] = 1
	sl.miningModifiers.StrikeOdds[3] = sl.miningModifiers.StrikeOdds[2] * 0.0025
	assert.Equal(t, 4, sl.findCartUpgrade())

	sl.miningModifiers.StrikeUpgrades[4] = 1
	sl.miningModifiers.StrikeOdds[4] = sl.miningModifiers.StrikeOdds[3] * 0.0025
	assert.Equal(t, 5, sl.findCartUpgrade())

	sl.miningModifiers.StrikeUpgrades[5] = 1
	sl.miningModifiers.StrikeOdds[5] = sl.miningModifiers.StrikeOdds[4] * 0.0025
	assert.Equal(t, 9000, sl.findCartUpgrade())

	sl.miningModifiers.GiantLuckLevel = 1
	sl.miningModifiers.GiantOdds = .001
	assert.Equal(t, 5, sl.findCartUpgrade())

	sl.miningModifiers.StrikeUpgrades[3] = 3
	sl.miningModifiers.StrikeOdds[3] = sl.miningModifiers.StrikeOdds[2] * 0.0025 * 3
	sl.miningModifiers.StrikeUpgrades[4] = 4
	sl.miningModifiers.StrikeOdds[4] = sl.miningModifiers.StrikeOdds[3] * 0.0025 * 4
	assert.Equal(t, 5, sl.findCartUpgrade())
}

func TestGetNextUpgrade(t *testing.T) {
	gc := NewGiantCalculator(
		NewMiningModifiers(
			0.5,
			1.0,
			0,
			0,
			map[int]int{
				2: 0,
				3: 0,
				4: 0,
				5: 0,
			},
			0,
			map[int]float64{
				2: 0.0,
				3: 0.0,
				4: 0.0,
				5: 0.0,
			},
		),
		false,
	)
	nu := gc.GetNextUpgrade(1000000)
	assert.Equal(t, "x2 strike", nu)

	gc.miningModifiers.StrikeUpgrades[DoubleStrike] = 1
	nu = gc.GetNextUpgrade(1000000)
	assert.Equal(t, "x3 strike", nu)

	gc.miningModifiers.StrikeUpgrades[TripleStrike] = 1
	nu = gc.GetNextUpgrade(1000000)
	assert.Equal(t, "x4 strike", nu)

	gc.miningModifiers.StrikeUpgrades[QuadrupleStrike] = 1
	nu = gc.GetNextUpgrade(1000000)
	assert.Equal(t, "x5 strike", nu)

	gc.miningModifiers.StrikeUpgrades[QuintupleStrike] = 1
	nu = gc.GetNextUpgrade(1000000)
	assert.Equal(t, "giant luck", nu)

	gc.miningModifiers.StrikeUpgrades[DoubleStrike] = 70
	gc.miningModifiers.StrikeUpgrades[TripleStrike] = 70
	gc.miningModifiers.StrikeUpgrades[QuadrupleStrike] = 70
	gc.miningModifiers.StrikeUpgrades[QuintupleStrike] = 70
	nu = gc.GetNextUpgrade(1000000)
	assert.Equal(t, "giant luck", nu)
}

func getCalculatorWithDummyPrices() Giant {
	return Giant{
		miningModifiers: NewMiningModifiers(
			0.5,
			1,
			0,
			0,
			strikeUpgrades{DoubleStrike: 0, TripleStrike: 0, QuadrupleStrike: 0, QuintupleStrike: 0},
			0,
			strikeOdds{DoubleStrike: 0, TripleStrike: 0, QuadrupleStrike: 0, QuintupleStrike: 0},
		),
		strikePrices: upgradeCostList{
			1:  20000,
			2:  40000,
			3:  60000,
			4:  80000,
			5:  100000,
			6:  120000,
			7:  140000,
			8:  160000,
			9:  180000,
			10: 200000,
		},
		giantLuckPrices: map[int]int{
			1:  100000,
			2:  1000000,
			3:  10000000,
			4:  100000000,
			5:  1000000000,
			6:  10000000000,
			7:  100000000000,
			8:  1000000000000,
			9:  10000000000000,
			10: 100000000000000,
		},
	}
}
