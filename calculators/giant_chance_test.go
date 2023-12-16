package calculators

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateGiantRollChance(t *testing.T) {
	sl := getCalculatorWithDummyPrices()
	assert.Equal(t, float64(0), sl.calculateGiantRollChance(0))

	sl.strikeUpgrades[1] = 1
	assert.Equal(t, float64(0), sl.calculateGiantRollChance(0))

	sl.strikeUpgrades[2] = 1
	sl.strikeUpgrades[3] = 1
	sl.strikeUpgrades[5] = 1
	assert.Equal(t, float64(0), sl.calculateGiantRollChance(0))

	sl.strikeUpgrades[4] = 1
	assert.Equal(t, float64(0), sl.calculateGiantRollChance(0))

	sl.giantLuckUpgrade = 1
	expected := fmt.Sprintf("%.14f", .01*.01*.01*.01*.001)
	calcResult := fmt.Sprintf("%.14f", sl.calculateGiantRollChance(0))
	assert.Equal(t, expected, calcResult)

	expected = fmt.Sprintf("%.14f", .02*.01*.01*.01*.001)
	calcResult = fmt.Sprintf("%.14f", sl.calculateGiantRollChance(2))
	assert.Equal(t, expected, calcResult)
	calcResult = fmt.Sprintf("%.14f", sl.calculateGiantRollChance(3))
	assert.Equal(t, expected, calcResult)
	calcResult = fmt.Sprintf("%.14f", sl.calculateGiantRollChance(4))
	assert.Equal(t, expected, calcResult)
	calcResult = fmt.Sprintf("%.14f", sl.calculateGiantRollChance(5))
	assert.Equal(t, expected, calcResult)

	sl.strikeUpgrades[2] = 2
	sl.strikeUpgrades[3] = 9
	sl.strikeUpgrades[4] = 1
	sl.strikeUpgrades[5] = 7
	expected = fmt.Sprintf("%.14f", .02*.09*.01*.07*.001)
	calcResult = fmt.Sprintf("%.14f", sl.calculateGiantRollChance(0))
	assert.Equal(t, expected, calcResult)
	expected = fmt.Sprintf("%.14f", .03*.09*.01*.07*.001)
	calcResult = fmt.Sprintf("%.14f", sl.calculateGiantRollChance(2))
	assert.Equal(t, expected, calcResult)
	expected = fmt.Sprintf("%.14f", .02*.10*.01*.07*.001)
	calcResult = fmt.Sprintf("%.14f", sl.calculateGiantRollChance(3))
	assert.Equal(t, expected, calcResult)
	expected = fmt.Sprintf("%.14f", .02*.09*.02*.07*.001)
	calcResult = fmt.Sprintf("%.14f", sl.calculateGiantRollChance(4))
	assert.Equal(t, expected, calcResult)
	expected = fmt.Sprintf("%.14f", .02*.09*.01*.08*.001)
	calcResult = fmt.Sprintf("%.14f", sl.calculateGiantRollChance(5))
	assert.Equal(t, expected, calcResult)
	sl.giantLuckUpgrade = 5
	expected = fmt.Sprintf("%.14f", .02*.09*.01*.07*.005)
	calcResult = fmt.Sprintf("%.14f", sl.calculateGiantRollChance(0))
	assert.Equal(t, expected, calcResult)
	expected = fmt.Sprintf("%.14f", .02*.09*.01*.07*.006)
	calcResult = fmt.Sprintf("%.14f", sl.calculateGiantRollChance(9000))
	assert.Equal(t, expected, calcResult)
}

func TestFindNextUpgrade(t *testing.T) {
	sl := getCalculatorWithDummyPrices()
	assert.Equal(t, 2, sl.findNextUpgrade())

	sl.strikeUpgrades[2] = 1
	assert.Equal(t, 3, sl.findNextUpgrade())

	sl.strikeUpgrades[3] = 1
	assert.Equal(t, 4, sl.findNextUpgrade())

	sl.strikeUpgrades[4] = 1
	assert.Equal(t, 5, sl.findNextUpgrade())

	sl.strikeUpgrades[5] = 1
	assert.Equal(t, 9000, sl.findNextUpgrade())

	sl.giantLuckUpgrade = 1
	assert.Equal(t, 2, sl.findNextUpgrade())

	sl.strikeUpgrades[3] = 3
	sl.strikeUpgrades[4] = 4
	assert.Equal(t, 2, sl.findNextUpgrade())

	sl.strikeUpgrades[2] = 2
	sl.strikeUpgrades[3] = 2
	sl.strikeUpgrades[4] = 2
	assert.Equal(t, 5, sl.findNextUpgrade())
}

func getCalculatorWithDummyPrices() GiantCalculator {
	return GiantCalculator{
		strikeUpgrades: map[int]int{},
		strikePrices: map[int]upgradeCostList{
			DoubleStrike: {
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
			TripleStrike: {
				1:  30000,
				2:  60000,
				3:  90000,
				4:  120000,
				5:  150000,
				6:  180000,
				7:  210000,
				8:  240000,
				9:  270000,
				10: 300000,
			},
			QuadrupleStrike: {
				1:  40000,
				2:  80000,
				3:  120000,
				4:  160000,
				5:  200000,
				6:  240000,
				7:  280000,
				8:  320000,
				9:  360000,
				10: 400000,
			},
			QuintupleStrike: {
				1:  50000,
				2:  100000,
				3:  150000,
				4:  200000,
				5:  250000,
				6:  300000,
				7:  350000,
				8:  400000,
				9:  450000,
				10: 500000,
			},
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
