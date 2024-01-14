package main

import (
	"capEndgame3Calculator/calculators"
	"fmt"
	"time"
)

func main() {
	giantCalc := calculators.NewGiantCalculator(
		calculators.NewOverclockConfig(false, true, true, true, true),
		1.09,
		1.2,
		0.65,
		map[int]int{
			2: 74,
			3: 73,
			4: 73,
			5: 73,
		},
		70,
	)
	//fmt.Println(
	//	fmt.Sprintf(
	//		"current chance per strike: %.10f%%",
	//		giantCalc.CalculateChancePerSTrike(1.0)*100,
	//	),
	//)
	//
	//fmt.Println("next upgrade should be", giantCalc.GetNextUpgrade())
	//
	giantCalc.PrintProbabilityDistribution(time.Hour*24, 1)
	//
	//giantCalc.CalculateUpgradePath()

	//fullPlanCalculator := calculators.NewGiantCalculator(
	//	calculators.NewOverclockConfig(false, false, false, false, false),
	//	1,
	//	1,
	//	0.5,
	//	map[int]int{2: 0, 3: 0, 4: 0, 5: 0},
	//	0,
	//)
	//fullPlanCalculator.CalculateUpgradePath()

	sc := calculators.NewStonesCalculator(1, .1825, .1825, .1825, .1825, .65, calculators.QuartzPick, 0.43, calculators.LegendaryEgg, false, false, true, true, true)
	//sc := calculators.NewStonesCalculator(1, calculators.RubyPick, 0.40, calculators.MythicEgg, true)
	//sc := calculators.NewStonesCalculator(.6, calculators.AmethystPick, 0.45, calculators.UncommonEgg, true)
	fmt.Println(sc.CalculateStones(time.Hour * 24))
}
