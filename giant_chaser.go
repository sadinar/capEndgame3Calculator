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
	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(duration)
	stoneCalc.PrintDamageChange(duration, shinyMods)
	p := message.NewPrinter(message.MatchLanguage("en"))
	fmt.Println(p.Sprintf("%d stones (%d genned and %d mined) gained in %v", gennedStones+minedStones, gennedStones, minedStones, duration))
}

func loadSadinar() (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, int, int) {
	miningMods := calculators.NewMiningModifiers(
		1.00+.5, // exactly as on stats screen
		100,     // exactly as shown on the wooden board behind egg
		.247,    // exactly as on stats screen
		408.8,   // exactly as on stats screen
		map[int]int{
			2: 78,
			3: 79,
			4: 80,
			5: 79,
		},
		90,
		map[int]float64{
			2: 30.8,  // exactly as on stats screen
			3: 9.733, // exactly as on stats screen
			4: 3.504, // exactly as on stats screen
			5: 1.384, // exactly as on stats screen
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
		1.12, // exactly as on stats screen
		100,  // exactly as shown on the wooden board behind egg
		.04,  // exactly as on stats screen
		210,  // exactly as on stats screen
		map[int]int{
			2: 71,
			3: 71,
			4: 71,
			5: 71,
		},
		61,
		map[int]float64{
			2: 24.15, // exactly as on stats screen
			3: 4.287, // exactly as on stats screen
			4: 1.37,  // exactly as on stats screen
			5: 0.486, // exactly as on stats screen
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
		calculators.ProdigiousEgg,
		true,
	)
	shinyMods := calculators.NewShinyModifiers(1.509) // exactly as seen on stats screen
	LabMods := calculators.NewGiantModifiers(1.04, 1, 1.06, 1.2, false, false)

	giantCalc := calculators.NewGiantCalculator(miningMods, LabMods)
	stoneCalc := calculators.NewStonesCalculator(miningMods, generationMods)

	nextSpeedUpgradeCost := 6 * HundredThousand
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
