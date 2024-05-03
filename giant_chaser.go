package main

import (
	"capEndgame3Calculator/calculators"
	"fmt"
	"golang.org/x/text/message"
	"time"
)

const HundredThousand = 100000
const Million = 1000000

func main() {
	//shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := loadSadinar()
	shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := loadSadinalt()
	//shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := loadAltinar()
	duration := time.Hour * 24

	fmt.Println("next giant chance upgrade should be", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))
	fmt.Println("next stone upgrade should be", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, nextCloneUpgradeCost))

	giantCalc.PrintProbabilityMedian(duration, shinyMods)
	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(duration)
	stoneCalc.PrintDamageChange(duration, shinyMods)
	p := message.NewPrinter(message.MatchLanguage("en"))
	fmt.Println(p.Sprintf("%d stones (%d genned and %d mined) gained in %v", gennedStones+minedStones, gennedStones, minedStones, duration))

	//fromScratchUpgradePath()
}

func loadSadinar() (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, int, int) {
	miningMods := calculators.NewMiningModifiers(
		1.02+.5, // exactly as on stats screen
		100,     // exactly as shown on the wooden board behind egg
		.516,    // exactly as on stats screen
		333.8,   // exactly as on stats screen
		map[int]int{
			2: 80,
			3: 88,
			4: 89,
			5: 88,
		},
		100,
		map[int]float64{
			2: 42,    // exactly as on stats screen
			3: 14.78, // exactly as on stats screen
			4: 5.921, // exactly as on stats screen
			5: 2.605, // exactly as on stats screen
		},
		true,
		true,
		true,
		true,
	)
	generationMods := calculators.NewEggGenerationModifiers(
		51,    // as shown on stats screen
		6.5,   // as shown on stats screen
		127.5, // as shown in stats pane
		calculators.MythicEgg,
		true,
	)
	shinyMods := calculators.NewShinyModifiers(100) // exactly as seen on stats screen
	giantLuckMods := calculators.NewGiantModifiers(1, 1, 1.1, 1.2, true, false)

	giantCalc := calculators.NewGiantCalculator(miningMods, giantLuckMods)
	stoneCalc := calculators.NewStonesCalculator(miningMods, generationMods)

	nextSpeedUpgradeCost := 2*Million + 5*HundredThousand
	nextCloneUpgradeCost := 1*Million + HundredThousand

	return shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost
}

func loadSadinalt() (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, int, int) {
	miningMods := calculators.NewMiningModifiers(
		1.13, // exactly as on stats screen
		100,  // exactly as shown on the wooden board behind egg
		.01,  // exactly as on stats screen
		390,  // exactly as on stats screen
		map[int]int{
			2: 71,
			3: 72,
			4: 71,
			5: 72,
		},
		64,
		map[int]float64{
			2: 20.25, // exactly as on stats screen
			3: 3.645, // exactly as on stats screen
			4: 0.647, // exactly as on stats screen
			5: 0.116, // exactly as on stats screen
		},
		false,
		false,
		false,
		false,
	)
	generationMods := calculators.NewEggGenerationModifiers(
		49,    // as shown on stats screen
		5.8,   // as shown on stats screen
		107.5, // as shown in stats pane
		calculators.MythicEgg,
		true,
	)
	shinyMods := calculators.NewShinyModifiers(1.509) // exactly as seen on stats screen
	LabMods := calculators.NewGiantModifiers(1.04, 1, 1.06, 1.2, false, false)

	giantCalc := calculators.NewGiantCalculator(miningMods, LabMods)
	stoneCalc := calculators.NewStonesCalculator(miningMods, generationMods)

	nextSpeedUpgradeCost := 8 * HundredThousand
	nextCloneUpgradeCost := 4 * HundredThousand

	return shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost
}

func loadAltinar() (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, int, int) {
	miningMods := calculators.NewMiningModifiers(
		.86,  // exactly as on stats screen
		100,  // exactly as shown on the wooden board behind egg
		.034, // exactly as on stats screen
		165,  // exactly as on stats screen
		map[int]int{
			2: 70,
			3: 70,
			4: 70,
			5: 70,
		},
		57,
		map[int]float64{
			2: 23.5,  // exactly as on stats screen
			3: 4.113, // exactly as on stats screen
			4: 1.295, // exactly as on stats screen
			5: 0.453, // exactly as on stats screen
		},
		false,
		false,
		false,
		true,
	)
	generationMods := calculators.NewEggGenerationModifiers(
		48,    // as shown on stats screen
		5.7,   // as shown on stats screen
		127.5, // as shown in stats pane
		calculators.UncommonEgg,
		true,
	)
	shinyMods := calculators.NewShinyModifiers(1.509) // exactly as seen on stats screen
	LabMods := calculators.NewGiantModifiers(1.048, 1, 1.05, 1.2, false, false)

	giantCalc := calculators.NewGiantCalculator(miningMods, LabMods)
	stoneCalc := calculators.NewStonesCalculator(miningMods, generationMods)

	nextSpeedUpgradeCost := 3 * HundredThousand
	nextCloneUpgradeCost := 3 * HundredThousand

	return shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost
}

func fromScratchUpgradePath() {
	mm := calculators.NewMiningModifiers(
		.5,
		100,
		0,
		100,
		map[int]int{2: 0, 3: 0, 4: 0, 5: 0},
		0,
		map[int]float64{2: 0, 3: 0, 4: 0, 5: 0},
		false,
		false,
		false,
		false,
	)

	gl := calculators.NewGiantModifiers(
		1,
		1,
		1,
		1,
		false,
		false,
	)

	gc := calculators.NewGiantCalculator(mm, gl)
	gc.CalculateUpgradePath()
}
