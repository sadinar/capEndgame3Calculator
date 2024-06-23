package main

import (
	"capEndgame3Calculator/calculators"
	"capEndgame3Calculator/character_config"
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

	//fromScratchUpgradePath()
}

func loadSadinar() (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, int, int) {
	return character_config.ConfigureCalculators("./character_config/sadinar.json")
}

func loadSadinalt() (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, int, int) {
	miningMods := calculators.NewMiningModifiers(
		1.19, // exactly as on stats screen
		100,  // exactly as shown on the wooden board behind egg
		.325, // exactly as on stats screen
		300,  // exactly as on stats screen
		map[int]int{
			2: 74,
			3: 76,
			4: 79,
			5: 79,
		},
		85,
		map[int]float64{
			2: 39.9,  // exactly as on stats screen
			3: 12.13, // exactly as on stats screen
			4: 4.312, // exactly as on stats screen
			5: 1.703, // exactly as on stats screen
		},
		true,
		false,
		true,
		true,
	)
	generationMods := calculators.NewEggGenerationModifiers(
		49,    // as shown on stats screen
		4.8,   // as shown on stats screen
		127.5, // as shown in stats pane
		calculators.MythicEgg,
		true,
	)
	shinyMods := calculators.NewShinyModifiers(51.07) // exactly as seen on stats screen
	LabMods := calculators.NewGiantModifiers(1.08, 1.03, 1.1, 1.2, true, true)

	giantCalc := calculators.NewGiantCalculator(miningMods, LabMods)
	stoneCalc := calculators.NewStonesCalculator(miningMods, generationMods)

	nextSpeedUpgradeCost := 1*Million + 6*HundredThousand
	nextCloneUpgradeCost := 4 * HundredThousand

	return shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost
}

func loadAltinar() (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, int, int) {
	miningMods := calculators.NewMiningModifiers(
		1.17, // exactly as on stats screen
		100,  // exactly as shown on the wooden board behind egg
		.227, // exactly as on stats screen
		300,  // exactly as on stats screen
		map[int]int{
			2: 72,
			3: 75,
			4: 75,
			5: 75,
		},
		72,
		map[int]float64{
			2: 39.2,  // exactly as on stats screen
			3: 11.76, // exactly as on stats screen
			4: 3.969, // exactly as on stats screen
			5: 1.488, // exactly as on stats screen
		},
		true,
		true,
		true,
		true,
	)
	generationMods := calculators.NewEggGenerationModifiers(
		48,    // as shown on stats screen
		4.7,   // as shown on stats screen
		107.5, // as shown in stats pane
		calculators.MythicEgg,
		true,
	)
	shinyMods := calculators.NewShinyModifiers(15.49) // exactly as seen on stats screen
	LabMods := calculators.NewGiantModifiers(1.08, 1.03, 1.1, 1.2, true, false)

	giantCalc := calculators.NewGiantCalculator(miningMods, LabMods)
	stoneCalc := calculators.NewStonesCalculator(miningMods, generationMods)

	nextSpeedUpgradeCost := 1*Million + 3*HundredThousand
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
	fmt.Println(gc.CalculateUpgradePath())
}
