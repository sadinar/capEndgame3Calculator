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
		0.72, // exactly as on stats screen
		1, // from the wooden board behind egg
		.137, // exactly as on stats screen
		map[int]int{
			2: 74,
			3: 74,
			4: 74,
			5: 74,
		},
		70,
		map[int]float64{
			2: 27.02, // exactly as on stats screen
			3: 7.998, // exactly as on stats screen
			4: 2.663, // exactly as on stats screen
			5: .985,  // exactly as on stats screen
		},
	)
	shinyMods := calculators.NewShinyModifiers(47.4) // exactly as seen on stats screen
	duration := time.Hour * 24

	giantCalc := calculators.NewGiantCalculator(miningMods, false)
	sc := calculators.NewStonesCalculator(
		miningMods,
		352.5, // stats screen w/ ingot: 472.5    w/o ingot: 322.5
		73,           // as shown on stats screen
		5.7,         // as shown on stats screen
		160,      // as shown in stats pane
		calculators.MythicEgg,
		true,
	)

	fmt.Println("next giant chance upgrade should be", giantCalc.GetNextUpgrade(1800000))
	fmt.Println("next stone upgrade should be", sc.FindNextUpgrade(1800000, OneMillion))

	giantCalc.PrintProbabilityMedian(duration, shinyMods)
	gennedStones, minedStones := sc.CalculateCombinedStones(duration)
	sc.PrintDamageChange(duration, shinyMods)
	p := message.NewPrinter(message.MatchLanguage("en"))
	fmt.Println(p.Sprintf("%d stones (%d genned and %d mined) gained in %v", gennedStones+minedStones, gennedStones, minedStones, duration))
}
