package main

import (
	"capEndgame3Calculator/calculators"
	"fmt"
)

func main() {
	giantCalc := calculators.NewGiantCalculator(
		calculators.OverclockConfig{QuadrupleEnabled: true, QuintupleEnabled: true, GiantLuckEnabled: true},
		1.06,
		1.146,
		map[int]int{
			2: 73,
			3: 73,
			4: 73,
			5: 73,
		},
		66,
	)
	//giantCalc.CalculateUpgradePath()
	fmt.Println(
		fmt.Sprintf(
			"current chance per strike: %.10f%%",
			giantCalc.CalculateChancePerSTrike()*100,
		),
	)
	fmt.Println("next upgrade should be", giantCalc.GetNextUpgrade())
}
