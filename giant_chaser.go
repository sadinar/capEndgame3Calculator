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

	trialsPerDay := giantCalc.GetEggsMinedPerDay()
	probabilitySuccess := giantCalc.CalculateChancePerSTrike(1)
	successCount, consumedProbabilitySpace := calculators.FindReasonableProbability(trialsPerDay, probabilitySuccess)
	reportedProbabilitySpace := 0.0
	medianProbability := 0.0
	medianSuccesses := 0
	for i := 0; i <= int(successCount); i++ {
		chance := calculators.BinomialProbability(trialsPerDay, uint64(i), probabilitySuccess)
		fmt.Println(fmt.Sprintf("%d: %.12f%%", i, chance*100))
		reportedProbabilitySpace += chance
		if medianSuccesses == 0 && reportedProbabilitySpace >= .5 {
			medianSuccesses = i
			medianProbability = chance
		}
	}
	fmt.Println(fmt.Sprintf("%d+: %.12f%%", successCount+1, (1-consumedProbabilitySpace)*100))
	fmt.Println(fmt.Sprintf("Median giant successes: %d @ %.12f%%", medianSuccesses, medianProbability*100))
}
