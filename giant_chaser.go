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
	miningMods := calculators.NewMiningModifiers(
		.76+.5, // exactly as on stats screen
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
			5: 1.076, // exactly as on stats screen
		},
	)
	generationMods := calculators.NewEggGenerationModifiers(
		51,    // as shown on stats screen
		6.5,   // as shown on stats screen
		127.5, // as shown in stats pane
		calculators.MythicEgg,
		true,
	)

	shinyMods := calculators.NewShinyModifiers(75.12) // exactly as seen on stats screen
	duration := time.Hour * 24

	giantCalc := calculators.NewGiantCalculator(miningMods, true)
	stoneCalc := calculators.NewStonesCalculator(miningMods, generationMods)

	nextSpeedUpgradeCost := 2*Million + 2*HundredThousand
	fmt.Println("next giant chance upgrade should be", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))
	fmt.Println("next stone upgrade should be", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, 1*Million))

	giantCalc.PrintProbabilityMedian(duration, shinyMods)
	gennedStones, minedStones := stoneCalc.CalculateCombinedStones(duration)
	stoneCalc.PrintDamageChange(duration, shinyMods)
	p := message.NewPrinter(message.MatchLanguage("en"))
	fmt.Println(p.Sprintf("%d stones (%d genned and %d mined) gained in %v", gennedStones+minedStones, gennedStones, minedStones, duration))
}
