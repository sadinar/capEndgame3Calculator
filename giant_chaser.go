package main

import (
	"capEndgame3Calculator/calculators"
	"fmt"
	"time"
)

func main() {
	giantCalc := calculators.NewGiantCalculator(
		calculators.NewOverclockConfig(false, false, true, true, true),
		1.07,
		1.188,
		0.64,
		map[int]int{
			2: 73,
			3: 73,
			4: 73,
			5: 73,
		},
		67,
	)
	fmt.Println(
		fmt.Sprintf(
			"current chance per strike: %.10f%%",
			giantCalc.CalculateChancePerSTrike(1.0)*100,
		),
	)
	fmt.Println("next upgrade should be", giantCalc.GetNextUpgrade())
	giantCalc.PrintProbabilityDistribution(time.Hour*24, 1.0)
}
