package calculators

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCaseOne(t *testing.T) {
	shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := caseOne()
	duration := time.Hour * 24

	assert.Equal(t, "x4 strike", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))
	assert.Equal(t, "speed", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, nextCloneUpgradeCost))

	medianIndex, shinyCount, medianProbability := giantCalc.PrintProbabilityMedian(duration, shinyMods)
	assert.Equal(t, 316, medianIndex)
	assert.Equal(t, 316, shinyCount)
	assert.Less(t, medianProbability, .50993)
	assert.Greater(t, medianProbability, .50992)

	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(duration)
	assert.Equal(t, 707752, gennedStones)
	assert.Equal(t, 3465559, minedStones)

	dmgOutput := stoneCalc.PrintDamageChange(duration, shinyMods)
	expectedGenMsg := "ascended generated: 226,394 (6,791,840 shiny score): ascended dmg multiplier gained: " +
		"x37.73245 (+73,578 dmg)\nmythic generated: 311,100 (12,444,008 shiny score): mythic dmg multiplier " +
		"gained: x51.85003 (+103,700 dmg)"
	assert.Equal(t, expectedGenMsg, dmgOutput)
}

func TestCaseTwo(t *testing.T) {
	shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := caseTwo()
	duration := time.Hour * 24

	assert.Equal(t, "giant luck", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))
	assert.Equal(t, "speed", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, nextCloneUpgradeCost))

	medianIndex, shinyCount, medianProbability := giantCalc.PrintProbabilityMedian(duration, shinyMods)
	assert.Equal(t, 39, medianIndex)
	assert.Equal(t, 0, shinyCount)
	assert.Less(t, medianProbability, .5612)
	assert.Greater(t, medianProbability, .5610)

	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(duration)
	assert.Equal(t, 628048, gennedStones)
	assert.Equal(t, 926702, minedStones)

	dmgOutput := stoneCalc.PrintDamageChange(duration, shinyMods)
	expectedOutput := "ascended generated: 233,884 (105,879 shiny score): ascended dmg multiplier gained: " +
		"x0.58822 (+1,147 dmg)\nmythic generated: 302,674 (182,694 shiny score): mythic dmg multiplier gained: " +
		"x0.76123 (+1,522 dmg)"
	assert.Equal(t, expectedOutput, dmgOutput)
}

func TestCaseThree(t *testing.T) {
	shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := caseThree()
	duration := time.Hour * 24

	assert.Equal(t, "speed", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))
	assert.Equal(t, "speed", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, nextCloneUpgradeCost))

	medianIndex, shinyCount, medianProbability := giantCalc.PrintProbabilityMedian(duration, shinyMods)
	assert.Equal(t, 25, medianIndex)
	assert.Equal(t, 0, shinyCount)
	assert.Less(t, medianProbability, 0.53199)
	assert.Greater(t, medianProbability, 0.53197)

	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(duration)
	assert.Equal(t, 680904, gennedStones)
	assert.Equal(t, 237897, minedStones)

	dmgOutput := stoneCalc.PrintDamageChange(duration, shinyMods)
	expectedOutput := "ascended generated: 238,217 (107,840 shiny score): ascended dmg multiplier gained: " +
		"x0.59912 (+1,168 dmg)\nmythic generated: 299,298 (180,656 shiny score): mythic dmg multiplier gained: " +
		"x0.75274 (+1,505 dmg)"
	assert.Equal(t, expectedOutput, dmgOutput)
}

func TestCaseFour(t *testing.T) {
	shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := caseFour()
	duration := time.Hour * 24

	assert.Equal(t, "x5 strike", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))
	assert.Equal(t, "speed", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, nextCloneUpgradeCost))

	medianIndex, shinyCount, medianProbability := giantCalc.PrintProbabilityMedian(duration, shinyMods)
	assert.Equal(t, 320, medianIndex)
	assert.Equal(t, 320, shinyCount)
	assert.Less(t, medianProbability, 0.51237)
	assert.Greater(t, medianProbability, 0.51235)

	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(duration)
	assert.Equal(t, 707752, gennedStones)
	assert.Equal(t, 3467013, minedStones)
}

func TestCaseFive(t *testing.T) {
	shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := caseFive()
	duration := time.Hour * 24

	assert.Equal(t, "giant luck", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))
	assert.Equal(t, "speed", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, nextCloneUpgradeCost))

	medianIndex, shinyCount, medianProbability := giantCalc.PrintProbabilityMedian(duration, shinyMods)
	assert.Equal(t, 328, medianIndex)
	assert.Equal(t, 328, shinyCount)
	assert.Less(t, medianProbability, 0.5171499)
	assert.Greater(t, medianProbability, 0.5171490)

	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(duration)
	assert.Equal(t, 707752, gennedStones)
	assert.Equal(t, 3471829, minedStones)
}

func TestCaseSix(t *testing.T) {
	shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := caseSix()
	duration := time.Hour * 12

	assert.Equal(t, "x5 strike", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))
	assert.Equal(t, "speed", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, nextCloneUpgradeCost))

	medianIndex, shinyCount, medianProbability := giantCalc.PrintProbabilityMedian(duration, shinyMods)
	assert.Equal(t, 258, medianIndex)
	assert.Equal(t, 258, shinyCount)
	assert.Less(t, medianProbability, 0.5150691417427202)
	assert.Greater(t, medianProbability, 0.5150691417427200)

	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(duration)
	assert.Equal(t, 353874, gennedStones)
	assert.Equal(t, 1603333, minedStones)

	dmgOutput := stoneCalc.PrintDamageChange(duration, shinyMods)
	expectedMessage := "ascended generated: 113,196 (3,395,905 shiny score): ascended dmg multiplier gained: " +
		"x18.86614 (+36,788 dmg)\nmythic generated: 155,549 (6,221,977 shiny score): mythic dmg multiplier " +
		"gained: x25.92490 (+51,849 dmg)"
	assert.Equal(t, expectedMessage, dmgOutput)
}

func TestCaseSeven(t *testing.T) {
	shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := caseSeven()
	duration := time.Hour * 12

	assert.Equal(t, "x3 strike", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))
	assert.Equal(t, "speed", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, nextCloneUpgradeCost))

	medianIndex, shinyCount, medianProbability := giantCalc.PrintProbabilityMedian(duration, shinyMods)
	assert.Equal(t, 335, medianIndex)
	assert.Equal(t, 335, shinyCount)
	assert.Less(t, medianProbability, 0.51700732893414999)
	assert.Greater(t, medianProbability, 0.5170073289341498)

	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(duration)
	assert.Equal(t, 353874, gennedStones)
	assert.Equal(t, 1627952, minedStones)
}

func TestNoRemainingUpgrades(t *testing.T) {
	_, giantCalc, stoneCalc, _, _ := caseOne()
	giantCalc.miningModifiers.StrikeUpgrades[2] = 100
	giantCalc.miningModifiers.StrikeUpgrades[3] = 100
	giantCalc.miningModifiers.StrikeUpgrades[4] = 100
	giantCalc.miningModifiers.StrikeUpgrades[5] = 100
	giantCalc.miningModifiers.GiantLuckLevel = 100

	assert.Equal(t, "n/a", giantCalc.GetNextUpgrade(UpgradeComplete))
	assert.Equal(t, "n/a", stoneCalc.FindNextUpgrade(UpgradeComplete, UpgradeComplete))
}

func TestSpeedEvaluated(t *testing.T) {
	_, giantCalc, _, _, _ := caseSix()

	assert.Equal(t, "speed", giantCalc.GetNextUpgrade(1900000))
}

func TestCloneEvaluated(t *testing.T) {
	_, _, stoneCalc, nextSpeedUpgradeCost, _ := caseSix()

	assert.Equal(t, "clone luck", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, 1500))
}

func caseOne() (ShinyModifiers, Giant, Stones, int, int) {
	miningMods := NewMiningModifiers(
		1.00+.5,
		100,
		.244,
		408.8,
		map[int]int{
			2: 78,
			3: 79,
			4: 79,
			5: 79,
		},
		90,
		map[int]float64{
			2: 30.8,
			3: 9.733,
			4: 3.46,
			5: 1.367,
		},
		true,
		true,
		true,
		true,
	)
	generationMods := NewEggGenerationModifiers(
		51,
		6.5,
		127.5,
		MythicEgg,
		true,
	)
	shinyMods := NewShinyModifiers(100)
	giantLuckMods := NewGiantModifiers(1, 1, 1.1, 1.2, true, false)

	giantCalc := NewGiantCalculator(miningMods, giantLuckMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods)

	return shinyMods, giantCalc, stoneCalc, 2500000, 1100000
}

func caseTwo() (ShinyModifiers, Giant, Stones, int, int) {
	miningMods := NewMiningModifiers(
		1.12,
		100,
		.04,
		210,
		map[int]int{
			2: 71,
			3: 71,
			4: 71,
			5: 71,
		},
		61,
		map[int]float64{
			2: 24.15,
			3: 4.287,
			4: 1.37,
			5: 0.486,
		},
		false,
		false,
		true,
		true,
	)
	generationMods := NewEggGenerationModifiers(
		49,
		5.8,
		107.5,
		ProdigiousEgg,
		true,
	)
	shinyMods := NewShinyModifiers(1.509)
	LabMods := NewGiantModifiers(1.04, 1, 1.06, 1.2, false, false)

	giantCalc := NewGiantCalculator(miningMods, LabMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods)

	return shinyMods, giantCalc, stoneCalc, 600000, 400000
}

func caseThree() (ShinyModifiers, Giant, Stones, int, int) {
	miningMods := NewMiningModifiers(
		.86,
		100,
		.034,
		165,
		map[int]int{
			2: 70,
			3: 70,
			4: 70,
			5: 70,
		},
		57,
		map[int]float64{
			2: 23.5,
			3: 4.113,
			4: 1.295,
			5: 0.453,
		},
		false,
		false,
		false,
		true,
	)
	generationMods := NewEggGenerationModifiers(
		48,
		5.7,
		127.5,
		UncommonEgg,
		true,
	)
	shinyMods := NewShinyModifiers(1.509)
	LabMods := NewGiantModifiers(1.048, 1, 1.05, 1.2, false, false)

	giantCalc := NewGiantCalculator(miningMods, LabMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods)

	return shinyMods, giantCalc, stoneCalc, 300000, 300000
}

func caseFour() (ShinyModifiers, Giant, Stones, int, int) {
	miningMods := NewMiningModifiers(
		1.00+.5,
		100,
		.247,
		408.8,
		map[int]int{
			2: 78,
			3: 79,
			4: 80,
			5: 79,
		},
		90,
		map[int]float64{
			2: 30.8,
			3: 9.733,
			4: 3.504,
			5: 1.384,
		},
		true,
		true,
		true,
		true,
	)
	generationMods := NewEggGenerationModifiers(
		51,
		6.5,
		127.5,
		MythicEgg,
		true,
	)
	shinyMods := NewShinyModifiers(100)
	giantLuckMods := NewGiantModifiers(1, 1, 1.1, 1.2, true, false)

	giantCalc := NewGiantCalculator(miningMods, giantLuckMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods)

	return shinyMods, giantCalc, stoneCalc, 2500000, 1100000
}

func caseFive() (ShinyModifiers, Giant, Stones, int, int) {
	miningMods := NewMiningModifiers(
		1.00+.5,
		100,
		.253,
		408.8,
		map[int]int{
			2: 78,
			3: 80,
			4: 80,
			5: 80,
		},
		90,
		map[int]float64{
			2: 30.8,
			3: 9.856,
			4: 3.548,
			5: 1.419,
		},
		true,
		true,
		true,
		true,
	)
	generationMods := NewEggGenerationModifiers(
		51,
		6.5,
		127.5,
		MythicEgg,
		true,
	)
	shinyMods := NewShinyModifiers(100)
	giantLuckMods := NewGiantModifiers(1, 1, 1.1, 1.2, true, false)

	giantCalc := NewGiantCalculator(miningMods, giantLuckMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods)

	return shinyMods, giantCalc, stoneCalc, 2500000, 1100000
}

func caseSix() (ShinyModifiers, Giant, Stones, int, int) {
	miningMods := NewMiningModifiers(
		1.02+.5,
		100,
		.393,
		333.8,
		map[int]int{
			2: 80,
			3: 81,
			4: 81,
			5: 80,
		},
		100,
		map[int]float64{
			2: 42,
			3: 13.61,
			4: 4.96,
			5: 1.984,
		},
		true,
		true,
		true,
		true,
	)
	generationMods := NewEggGenerationModifiers(
		51,
		6.5,
		127.5,
		MythicEgg,
		true,
	)
	shinyMods := NewShinyModifiers(100)
	giantLuckMods := NewGiantModifiers(1, 1, 1.1, 1.2, true, false)

	giantCalc := NewGiantCalculator(miningMods, giantLuckMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods)

	return shinyMods, giantCalc, stoneCalc, 2500000, 1100000
}

func caseSeven() (ShinyModifiers, Giant, Stones, int, int) {
	miningMods := NewMiningModifiers(
		1.02+.5,
		100,
		.51,
		333.8,
		map[int]int{
			2: 80,
			3: 87,
			4: 89,
			5: 88,
		},
		100,
		map[int]float64{
			2: 42,
			3: 14.62,
			4: 5.854,
			5: 2.576,
		},
		true,
		true,
		true,
		true,
	)
	generationMods := NewEggGenerationModifiers(
		51,
		6.5,
		127.5,
		MythicEgg,
		true,
	)
	shinyMods := NewShinyModifiers(100)
	giantLuckMods := NewGiantModifiers(1, 1, 1.1, 1.2, true, false)

	giantCalc := NewGiantCalculator(miningMods, giantLuckMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods)

	return shinyMods, giantCalc, stoneCalc, 2500000, 1100000
}
