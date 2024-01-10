package calculators

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateGiantRollChance(t *testing.T) {
	sl := getCalculatorWithDummyPrices()
	assert.Equal(t, float64(0), sl.calculateBaseGiantChance(0))

	sl.strikeUpgrades[1] = 1
	assert.Equal(t, float64(0), sl.calculateBaseGiantChance(0))

	sl.strikeUpgrades[2] = 1
	sl.strikeUpgrades[3] = 1
	sl.strikeUpgrades[5] = 1
	assert.Equal(t, float64(0), sl.calculateBaseGiantChance(0))

	sl.strikeUpgrades[4] = 1
	assert.Equal(t, float64(0), sl.calculateBaseGiantChance(0))

	sl.giantLuckUpgrade = 1
	expected := fmt.Sprintf("%.16f", 0.0025*0.0025*0.0025*0.0025*0.001)
	calcResult := fmt.Sprintf("%.16f", sl.calculateBaseGiantChance(0))
	assert.Equal(t, expected, calcResult)

	expected = fmt.Sprintf("%.14f", .0025*2*0.0025*0.0025*0.0025*.001)
	calcResult = fmt.Sprintf("%.14f", sl.calculateBaseGiantChance(2))
	assert.Equal(t, expected, calcResult)
	calcResult = fmt.Sprintf("%.14f", sl.calculateBaseGiantChance(3))
	assert.Equal(t, expected, calcResult)
	calcResult = fmt.Sprintf("%.14f", sl.calculateBaseGiantChance(4))
	assert.Equal(t, expected, calcResult)
	calcResult = fmt.Sprintf("%.14f", sl.calculateBaseGiantChance(5))
	assert.Equal(t, expected, calcResult)

	sl.strikeUpgrades[2] = 2
	sl.strikeUpgrades[3] = 9
	sl.strikeUpgrades[4] = 1
	sl.strikeUpgrades[5] = 7
	expected = fmt.Sprintf("%.14f", .0025*2*.0025*9*.0025*1*.0025*7*.001)
	calcResult = fmt.Sprintf("%.14f", sl.calculateBaseGiantChance(0))
	assert.Equal(t, expected, calcResult)
	expected = fmt.Sprintf("%.14f", .0025*3*.0025*9*.0025*1*.0025*7*.001)
	calcResult = fmt.Sprintf("%.14f", sl.calculateBaseGiantChance(2))
	assert.Equal(t, expected, calcResult)
	expected = fmt.Sprintf("%.14f", .0025*2*.0025*10*.0025*1*.0025*7*.001)
	calcResult = fmt.Sprintf("%.14f", sl.calculateBaseGiantChance(3))
	assert.Equal(t, expected, calcResult)
	expected = fmt.Sprintf("%.14f", .0025*2*.0025*9*.0025*2*.0025*7*.001)
	calcResult = fmt.Sprintf("%.14f", sl.calculateBaseGiantChance(4))
	assert.Equal(t, expected, calcResult)
	expected = fmt.Sprintf("%.14f", .0025*2*.0025*9*.0025*1*.0025*8*.001)
	calcResult = fmt.Sprintf("%.14f", sl.calculateBaseGiantChance(5))
	assert.Equal(t, expected, calcResult)
	sl.giantLuckUpgrade = 5
	expected = fmt.Sprintf("%.14f", .0025*2*.0025*9*.0025*1*.0025*7*.001*5)
	calcResult = fmt.Sprintf("%.14f", sl.calculateBaseGiantChance(0))
	assert.Equal(t, expected, calcResult)
	expected = fmt.Sprintf("%.14f", .0025*2*.0025*9*.0025*1*.0025*7*.001*6)
	calcResult = fmt.Sprintf("%.14f", sl.calculateBaseGiantChance(9000))
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

func TestGetNextUpgrade(t *testing.T) {
	gc := NewGiantCalculator(
		NewOverclockConfig(false, false, false, false, false),
		1.0,
		1.0,
		0.5,
		map[int]int{
			2: 0,
			3: 0,
			4: 0,
			5: 0,
		},
		0,
	)
	nu := gc.GetNextUpgrade()
	assert.Equal(t, "upgrade x2 strike", nu)

	gc.strikeUpgrades[DoubleStrike] = 1
	nu = gc.GetNextUpgrade()
	assert.Equal(t, "upgrade x3 strike", nu)

	gc.strikeUpgrades[TripleStrike] = 1
	nu = gc.GetNextUpgrade()
	assert.Equal(t, "upgrade x4 strike", nu)

	gc.strikeUpgrades[QuadrupleStrike] = 1
	nu = gc.GetNextUpgrade()
	assert.Equal(t, "upgrade x5 strike", nu)

	gc.strikeUpgrades[QuintupleStrike] = 1
	nu = gc.GetNextUpgrade()
	assert.Equal(t, "giant luck", nu)

	gc.strikeUpgrades[DoubleStrike] = 70
	gc.strikeUpgrades[TripleStrike] = 70
	gc.strikeUpgrades[QuadrupleStrike] = 70
	gc.strikeUpgrades[QuintupleStrike] = 70
	nu = gc.GetNextUpgrade()
	assert.Equal(t, "giant luck", nu)

	gc.giantLuckUpgrade = 59
	nu = gc.GetNextUpgrade()
	assert.Equal(t, "giant luck", nu)

	gc.giantLuckUpgrade = 60
	nu = gc.GetNextUpgrade()
	assert.Equal(t, "upgrade x2 strike", nu)
}

func getCalculatorWithDummyPrices() GiantCalculator {
	return GiantCalculator{
		strikeUpgrades: map[int]int{},
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
