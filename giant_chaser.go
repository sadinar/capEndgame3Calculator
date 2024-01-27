package main

import (
	"capEndgame3Calculator/calculators"
	"fmt"
	"time"
)

func main() {
	ocConfig := calculators.NewOverclockConfig(true, false, true, true, true, true)
	userMods := calculators.NewUserModifiers(
		1.1,
		1.2,
		0.66,
		1,
		map[int]int{
			2: 74,
			3: 73,
			4: 73,
			5: 73,
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
	fmt.Println(fmt.Sprintf("%d stones gained in %v", sc.CalculateCombinedStones(duration), duration))
}
