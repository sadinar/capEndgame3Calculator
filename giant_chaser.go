package main

import (
	"capEndgame3Calculator/calculators"
	"fmt"
)

func main() {
	giantCalc := calculators.NewGiantCalculator(
		calculators.NewOverclockConfig(false, false, true, true, true),
		1.07,
		1.184,
		0.64,
		map[int]int{
			2: 73,
			3: 73,
			4: 73,
			5: 73,
		},
		67,
	)
	//giantCalc.CalculateUpgradePath()
	fmt.Println(
		fmt.Sprintf(
			"current chance per strike: %.10f%%",
			giantCalc.CalculateChancePerSTrike(0.875)*100,
		),
	)
	fmt.Println("next upgrade should be", giantCalc.GetNextUpgrade())
	giantCalc.PrintProbabilityDistribution()
}
