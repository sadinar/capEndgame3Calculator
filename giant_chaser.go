package main

import (
	"capEndgame3Calculator/calculators"
	"fmt"
	"golang.org/x/text/message"
	"time"
)

const OneMillion = 1000000

func main() {
	ocConfig := calculators.NewOverclockConfig(true, true, true, true, true, true, false)
	miningMods := calculators.NewMiningModifiers(
		1.1,
		1.2,
		0.72,
		1,
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
	shinyMods := calculators.NewShinyModifiers(14.87) // exactly as seen on stats screen
	duration := time.Hour * 24

	giantCalc := calculators.NewGiantCalculator(ocConfig, miningMods)
	sc := calculators.NewStonesCalculator(
		miningMods,
		472.5, // stats screen w/ ingot: 472.5    w/o ingot: 322.5
		61,    // as shown on stats screen
		3,     // as shown on stats screen
		160,   // as shown in stats pane
		calculators.MythicEgg,
		true,
	)

	if giantCalc.SpeedComparison(1800000, duration*5) {
		fmt.Println("next giant chance upgrade should be speed")
	} else {
		fmt.Println("next giant chance upgrade should be", giantCalc.GetNextUpgrade())
	}
	fmt.Println("next stone upgrade should be", sc.FindNextUpgrade(1800000, OneMillion))

	giantCalc.PrintProbabilityMedian(duration, shinyMods)
	gennedStones, minedStones := sc.CalculateCombinedStones(duration)
	sc.PrintDamageChange(duration, shinyMods)
	p := message.NewPrinter(message.MatchLanguage("en"))
	fmt.Println(p.Sprintf("%d stones (%d genned and %d mined) gained in %v", gennedStones+minedStones, gennedStones, minedStones, duration))
	fmt.Println(p.Sprintf("Extra stones per day from +0.1%% genned pets: %v", float64(gennedStones)*0.001))
	fmt.Println(p.Sprintf("Shiny odds w/o mine OC: %.4f%%", 100*shinyMods.CalculateShinyOdds()))
}
