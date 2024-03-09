package main

import (
	"capEndgame3Calculator/calculators"
	"fmt"
	"golang.org/x/text/message"
	"time"
)

const OneMillion = 1000000

func main() {
	miningMods := calculators.NewMiningModifiers(
		1.22, // exactly as on stats screen
		100,  // exactly as shown on the wooden board behind egg
		.131, // exactly as on stats screen
		map[int]int{
			2: 74,
			3: 74,
			4: 74,
			5: 74,
		},
		70,
		map[int]float64{
			2: 25.9,  // exactly as on stats screen
			3: 7.666, // exactly as on stats screen
			4: 2.553, // exactly as on stats screen
			5: 0.945, // exactly as on stats screen
		},
	)
	shinyMods := calculators.NewShinyModifiers(62.97) // exactly as seen on stats screen
	duration := time.Hour * 24

	giantCalc := calculators.NewGiantCalculator(miningMods, true)
	sc := calculators.NewStonesCalculator(
		miningMods,
		405,   // stats screen
		62,    // as shown on stats screen
		5.9,   // as shown on stats screen
		127.5, // as shown in stats pane
		calculators.MythicEgg,
		true,
	)

	fmt.Println("next giant chance upgrade should be", giantCalc.GetNextUpgrade(1800000))
	fmt.Println("next stone upgrade should be", sc.FindNextUpgrade(1800000, 500000))

	giantCalc.PrintProbabilityMedian(duration, shinyMods)
	gennedStones, minedStones := sc.CalculateCombinedStones(duration)
	sc.PrintDamageChange(duration, shinyMods)
	p := message.NewPrinter(message.MatchLanguage("en"))
	fmt.Println(p.Sprintf("%d stones (%d genned and %d mined) gained in %v", gennedStones+minedStones, gennedStones, minedStones, duration))
}
