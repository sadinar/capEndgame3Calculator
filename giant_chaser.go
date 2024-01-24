package main

import (
	"capEndgame3Calculator/calculators"
	"fmt"
	"time"
)

func main() {
	ocConfig := calculators.NewOverclockConfig(true, false, false, true, true, false)
	userMods := calculators.NewUserModifiers(
		1.09,
		1.2,
		0.65,
		1,
		map[int]int{
			2: 74,
			3: 73,
			4: 73,
			5: 73,
		},
		70,
	)

	giantCalc := calculators.NewGiantCalculator(ocConfig, userMods)
	fmt.Println("next giant chance upgrade should be", giantCalc.GetNextUpgrade())
	giantCalc.PrintProbabilityMedian(time.Hour * 24)

	sc := calculators.NewStonesCalculator(
		userMods,
		calculators.RubyPick,
		0.43,
		calculators.MythicEgg,
		ocConfig,
	)
	duration := time.Hour * 24
	fmt.Println(fmt.Sprintf("%d stones gained in %v", sc.CalculateCombinedStones(duration), duration))
}
