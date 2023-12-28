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

	trialsPerDay := uint64(0.64 * 60 * 60 * 24)
	pSuccess := giantCalc.CalculateChancePerSTrike(0.875)
	successCount, totalProbabilitySpace := calculators.FindReasonableProbability(trialsPerDay, pSuccess)
	for i := 0; i <= int(successCount); i++ {
		chance := calculators.BinomialProbability(trialsPerDay, uint64(i), pSuccess)
		fmt.Println(fmt.Sprintf("%d: %.12f%%", i, chance*100))
	}
	fmt.Println(fmt.Sprintf("%d+: %.12f%%", successCount+1, (1-totalProbabilitySpace)*100))
}
