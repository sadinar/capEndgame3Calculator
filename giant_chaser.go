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
	shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := loadSadinar()
	//shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := loadSadinalt()
	//shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := loadAltinar()
	duration := time.Hour * 24

	fmt.Println("next giant chance upgrade should be", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))
	fmt.Println("next stone upgrade should be", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, nextCloneUpgradeCost))

	giantCalc.PrintProbabilityMedian(duration, shinyMods)
	gennedStones, minedStones := stoneCalc.CalculateCombinedStones(duration)
	stoneCalc.PrintDamageChange(duration, shinyMods)
	p := message.NewPrinter(message.MatchLanguage("en"))
	fmt.Println(p.Sprintf("%d stones (%d genned and %d mined) gained in %v", gennedStones+minedStones, gennedStones, minedStones, duration))
}

func loadSadinar() (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, int, int) {
	miningMods := calculators.NewMiningModifiers(
		.96+.5, // exactly as on stats screen
		100,    // exactly as shown on the wooden board behind egg
		.149,   // exactly as on stats screen
		408.8,  // exactly as on stats screen
		map[int]int{
			2: 74,
			3: 74,
			4: 74,
			5: 74,
		},
		70,
		map[int]float64{
			2: 29.4,  // exactly as on stats screen
			3: 8.702, // exactly as on stats screen
			4: 2.898, // exactly as on stats screen
			5: 1.072, // exactly as on stats screen
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
	LabMods := calculators.NewGiantModifiers(1, 1, 1.1, 1.2, true)

	giantCalc := calculators.NewGiantCalculator(miningMods, LabMods, true)
	stoneCalc := calculators.NewStonesCalculator(miningMods, generationMods)

	nextSpeedUpgradeCost := 2*Million + 5*HundredThousand
	nextCloneUpgradeCost := 1*Million + HundredThousand

	return shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost
}

func loadSadinalt() (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, int, int) {
	miningMods := calculators.NewMiningModifiers(
		1.04, // exactly as on stats screen
		100,  // exactly as shown on the wooden board behind egg
		.035, // exactly as on stats screen
		165,  // exactly as on stats screen
		map[int]int{
			2: 69,
			3: 69,
			4: 69,
			5: 69,
		},
		59,
		map[int]float64{
			2: 23.5,  // exactly as on stats screen
			3: 4.113, // exactly as on stats screen
			4: 1.295, // exactly as on stats screen
			5: 0.453, // exactly as on stats screen
		},
		false,
		false,
		true,
		true,
	)
	generationMods := calculators.NewEggGenerationModifiers(
		49,    // as shown on stats screen
		5.8,   // as shown on stats screen
		107.5, // as shown in stats pane
		calculators.UncommonEgg,
		true,
	)
	shinyMods := calculators.NewShinyModifiers(1.509) // exactly as seen on stats screen
	LabMods := calculators.NewGiantModifiers(1.04, 1, 1.06, 1.2, false)

	giantCalc := calculators.NewGiantCalculator(miningMods, LabMods, false)
	stoneCalc := calculators.NewStonesCalculator(miningMods, generationMods)

	nextSpeedUpgradeCost := 6 * HundredThousand
	nextCloneUpgradeCost := 4 * HundredThousand

	return shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost
}

func loadAltinar() (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, int, int) {
	miningMods := calculators.NewMiningModifiers(
		.6,   // exactly as on stats screen
		100,  // exactly as shown on the wooden board behind egg
		.026, // exactly as on stats screen
		165,  // exactly as on stats screen
		map[int]int{
			2: 68,
			3: 67,
			4: 67,
			5: 67,
		},
		54,
		map[int]float64{
			2: 23,    // exactly as on stats screen
			3: 3.853, // exactly as on stats screen
			4: 1.162, // exactly as on stats screen
			5: 0.389, // exactly as on stats screen
		},
		false,
		false,
		false,
		true,
	)
	generationMods := calculators.NewEggGenerationModifiers(
		47,    // as shown on stats screen
		5.7,   // as shown on stats screen
		107.5, // as shown in stats pane
		calculators.CommonEgg,
		true,
	)
	shinyMods := calculators.NewShinyModifiers(1.253) // exactly as seen on stats screen
	LabMods := calculators.NewGiantModifiers(1.072, 1, 1.03, 1.138, false)

	giantCalc := calculators.NewGiantCalculator(miningMods, LabMods, false)
	stoneCalc := calculators.NewStonesCalculator(miningMods, generationMods)

	nextSpeedUpgradeCost := 3 * HundredThousand
	nextCloneUpgradeCost := 3 * HundredThousand

	return shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost
}
