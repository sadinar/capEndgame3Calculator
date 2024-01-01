package main

import (
	"capEndgame3Calculator/calculators"
	"fmt"
	"time"
)

func main() {
	giantCalc := calculators.NewGiantCalculator(
		calculators.NewOverclockConfig(false, true, true, true, true),
		1.07,
		1.2,
		0.65,
		map[int]int{
			2: 73,
			3: 73,
			4: 73,
			5: 73,
		},
		69,
	)
	fmt.Println(
		fmt.Sprintf(
			"current chance per strike: %.10f%%",
			giantCalc.CalculateChancePerSTrike(1.0)*100,
		),
	)

	fmt.Println("next upgrade should be", giantCalc.GetNextUpgrade())

	giantCalc.PrintProbabilityDistribution(time.Hour*24, 0.783)

	//fullPlanCalculator := calculators.NewGiantCalculator(
	//	calculators.NewOverclockConfig(false, false, false, false, false),
	//	1,
	//	1,
	//	0.5,
	//	map[int]int{2: 0, 3: 0, 4: 0, 5: 0},
	//	0,
	//)
	//fullPlanCalculator.CalculateUpgradePath()
}
