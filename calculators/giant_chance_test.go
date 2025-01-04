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

func TestCaseEight(t *testing.T) {
	shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := caseEight()
	duration := time.Hour * 12

	assert.Equal(t, "x2 strike", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))
	assert.Equal(t, "speed", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, nextCloneUpgradeCost))

	medianIndex, shinyCount, medianProbability := giantCalc.PrintProbabilityMedian(duration, shinyMods)
	assert.Equal(t, 29, medianIndex)
	assert.Equal(t, 7, shinyCount)
	assert.Less(t, medianProbability, 0.544371071059593)
	assert.Greater(t, medianProbability, 0.5443710710595928)

	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(duration)
	assert.Equal(t, 314023, gennedStones)
	assert.Equal(t, 1234330, minedStones)
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

func TestFullUpgradePath(t *testing.T) {
	miningMods := NewMiningModifiers(
		1.02+.5,
		100,
		.564,
		348.8,
		map[int]int{
			2: 80,
			3: 91,
			4: 91,
			5: 91,
		},
		100,
		map[int]float64{
			2: 42,
			3: 15.29,
			4: 6.26,
			5: 2.848,
		},
		true,
		true,
		true,
		true,
	)
	giantLuckMods := NewGiantModifiers(1, 1, 1.1, 1.2, true, false)
	ascendMods := NewAscensionModifiers(0, 1)
	giantCalc := NewGiantCalculator(miningMods, giantLuckMods, ascendMods)

	fullPath := giantCalc.CalculateUpgradePath()
	expectedOutput := "" +
		"------------------------------------------------------------\n" +
		"| x2 | x3 | x4 | x5 | giant |    chance/hit   | stone cost\n" +
		"|080 |091 |092 |091 |100    | 0.564000000000% | 189950000\n" +
		"|080 |091 |092 |092 |100    | 0.564000000000% | 192450000\n" +
		"|080 |092 |092 |092 |100    | 0.564000000000% | 194950000\n" +
		"|080 |092 |093 |092 |100    | 0.564000000000% | 197450000\n" +
		"|080 |092 |093 |093 |100    | 0.564000000000% | 199950000\n" +
		"|080 |093 |093 |093 |100    | 0.564000000000% | 202450000\n" +
		"|080 |093 |094 |093 |100    | 0.564000000000% | 204950000\n" +
		"|080 |093 |094 |094 |100    | 0.564000000000% | 207450000\n" +
		"|080 |094 |094 |094 |100    | 0.564000000000% | 209950000\n" +
		"|080 |094 |095 |094 |100    | 0.564000000000% | 212450000\n" +
		"|080 |094 |095 |095 |100    | 0.564000000000% | 214950000\n" +
		"|080 |095 |095 |095 |100    | 0.564000000000% | 217450000\n" +
		"|080 |095 |096 |095 |100    | 0.564000000000% | 219950000\n" +
		"|080 |095 |096 |096 |100    | 0.564000000000% | 222450000\n" +
		"|080 |096 |096 |096 |100    | 0.564000000000% | 224950000\n" +
		"|080 |096 |097 |096 |100    | 0.564000000000% | 227450000\n" +
		"|080 |096 |097 |097 |100    | 0.564000000000% | 229950000\n" +
		"|081 |096 |097 |097 |100    | 0.564000000000% | 231950000\n" +
		"|081 |097 |097 |097 |100    | 0.564000000000% | 234450000\n" +
		"|082 |097 |097 |097 |100    | 0.564000000000% | 236450000\n" +
		"|082 |097 |098 |097 |100    | 0.564000000000% | 238950000\n" +
		"|082 |097 |098 |098 |100    | 0.564000000000% | 241450000\n" +
		"|082 |098 |098 |098 |100    | 0.564000000000% | 243950000\n" +
		"|083 |098 |098 |098 |100    | 0.564000000000% | 245950000\n" +
		"|083 |098 |099 |098 |100    | 0.564000000000% | 248450000\n" +
		"|083 |098 |099 |099 |100    | 0.564000000000% | 250950000\n" +
		"|083 |099 |099 |099 |100    | 0.564000000000% | 253450000\n" +
		"|084 |099 |099 |099 |100    | 0.564000000000% | 255450000\n" +
		"|084 |099 |100 |099 |100    | 0.564000000000% | 257950000\n" +
		"|084 |099 |100 |100 |100    | 0.564000000000% | 260450000\n" +
		"|084 |100 |100 |100 |100    | 0.564000000000% | 262950000\n" +
		"|085 |100 |100 |100 |100    | 0.564000000000% | 264950000\n" +
		"|086 |100 |100 |100 |100    | 0.564000000000% | 266950000\n" +
		"|087 |100 |100 |100 |100    | 0.564000000000% | 268950000\n" +
		"|088 |100 |100 |100 |100    | 0.564000000000% | 270950000\n" +
		"|089 |100 |100 |100 |100    | 0.564000000000% | 272950000\n" +
		"|090 |100 |100 |100 |100    | 0.564000000000% | 274950000\n" +
		"|091 |100 |100 |100 |100    | 0.564000000000% | 277450000\n" +
		"|092 |100 |100 |100 |100    | 0.564000000000% | 279950000\n" +
		"|093 |100 |100 |100 |100    | 0.564000000000% | 282450000\n" +
		"|094 |100 |100 |100 |100    | 0.564000000000% | 284950000\n" +
		"|095 |100 |100 |100 |100    | 0.564000000000% | 287450000\n" +
		"|096 |100 |100 |100 |100    | 0.564000000000% | 289950000\n" +
		"|097 |100 |100 |100 |100    | 0.564000000000% | 292450000\n" +
		"|098 |100 |100 |100 |100    | 0.564000000000% | 294950000\n" +
		"|099 |100 |100 |100 |100    | 0.564000000000% | 297450000\n" +
		"|100 |100 |100 |100 |100    | 0.564000000000% | 299950000\n"

	assert.Equal(t, expectedOutput, fullPath)
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
		false,
	)
	shinyMods := NewShinyModifiers(100)
	giantLuckMods := NewGiantModifiers(1, 1, 1.1, 1.2, true, false)
	ascendMods := NewAscensionModifiers(0, 1)

	giantCalc := NewGiantCalculator(miningMods, giantLuckMods, ascendMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods, ascendMods)

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
		false,
	)
	shinyMods := NewShinyModifiers(1.509)
	LabMods := NewGiantModifiers(1.04, 1, 1.06, 1.2, false, false)
	ascendMods := NewAscensionModifiers(0, 1)

	giantCalc := NewGiantCalculator(miningMods, LabMods, ascendMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods, ascendMods)

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
		false,
	)
	shinyMods := NewShinyModifiers(1.509)
	LabMods := NewGiantModifiers(1.048, 1, 1.05, 1.2, false, false)
	ascendMods := NewAscensionModifiers(0, 1)

	giantCalc := NewGiantCalculator(miningMods, LabMods, ascendMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods, ascendMods)

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
		false,
	)
	shinyMods := NewShinyModifiers(100)
	giantLuckMods := NewGiantModifiers(1, 1, 1.1, 1.2, true, false)
	ascendMods := NewAscensionModifiers(0, 0)

	giantCalc := NewGiantCalculator(miningMods, giantLuckMods, ascendMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods, ascendMods)

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
		false,
	)
	shinyMods := NewShinyModifiers(100)
	giantLuckMods := NewGiantModifiers(1, 1, 1.1, 1.2, true, false)
	ascendMods := NewAscensionModifiers(0, 1)

	giantCalc := NewGiantCalculator(miningMods, giantLuckMods, ascendMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods, ascendMods)

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
		false,
	)
	shinyMods := NewShinyModifiers(100)
	giantLuckMods := NewGiantModifiers(1, 1, 1.1, 1.2, true, false)
	ascendMods := NewAscensionModifiers(0, .99)

	giantCalc := NewGiantCalculator(miningMods, giantLuckMods, ascendMods)
	stoneCalc := NewStonesCalculator(miningMods, generationMods, ascendMods)

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
		false,
	)
	shinyMods := NewShinyModifiers(100)
	giantLuckMods := NewGiantModifiers(1, 1, 1.1, 1.2, true, false)

	giantCalc := NewGiantCalculator(miningMods, giantLuckMods, AscensionModifiers{})
	stoneCalc := NewStonesCalculator(miningMods, generationMods, AscensionModifiers{})

	return shinyMods, giantCalc, stoneCalc, 2500000, 1100000
}

func caseEight() (ShinyModifiers, Giant, Stones, int, int) {
	miningMods := NewMiningModifiers(
		1.16,  // exactly as on stats screen
		100,   // exactly as shown on the wooden board behind egg
		.058,  // exactly as on stats screen
		401.2, // exactly as on stats screen
		map[int]int{
			2: 72,
			3: 74,
			4: 74,
			5: 74,
		},
		67,
		map[int]float64{
			2: 28.7,  // exactly as on stats screen
			3: 5.31,  // exactly as on stats screen
			4: 1.768, // exactly as on stats screen
			5: 0.654, // exactly as on stats screen
		},
		true,
		false,
		true,
		true,
	)
	generationMods := NewEggGenerationModifiers(
		49,    // as shown on stats screen
		5.8,   // as shown on stats screen
		107.5, // as shown in stats pane
		MythicEgg,
		true,
		false,
	)
	shinyMods := NewShinyModifiers(24.26) // exactly as seen on stats screen
	LabMods := NewGiantModifiers(1, 1, 1.1, 1.2, false, false)

	giantCalc := NewGiantCalculator(miningMods, LabMods, AscensionModifiers{})
	stoneCalc := NewStonesCalculator(miningMods, generationMods, AscensionModifiers{})

	return shinyMods, giantCalc, stoneCalc, 1200000, 400000
}
