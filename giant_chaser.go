package main

import (
	"capEndgame3Calculator/calculators"
	"fmt"
	"time"
)

func main() {
	ocConfig := calculators.NewOverclockConfig(true, true, false, true, true, true)
	userMods := calculators.NewUserModifiers(
		1.1,
		1.2,
		0.71,
		1,
		map[int]int{
			2: 74,
			3: 74,
			4: 74,
			5: 74,
		},
		70,
	)
	duration := time.Hour * 24

	giantCalc := calculators.NewGiantCalculator(ocConfig, userMods)
	fmt.Println("next giant chance upgrade should be", giantCalc.GetNextUpgrade())
	giantCalc.PrintProbabilityMedian(duration)

	sc := calculators.NewStonesCalculator(
		userMods,
		calculators.RubyPick,
		0.43,
		calculators.MythicEgg,
		ocConfig,
	)
	gennedStones, minedStones := sc.CalculateCombinedStones(duration)
	fmt.Println(fmt.Sprintf("%d stones (%d genned and %d mined) gained in %v", gennedStones+minedStones, gennedStones, minedStones, duration))
}
