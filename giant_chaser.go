package main

import "capEndgame3Calculator/calculators"

func main() {
	giantCalc := calculators.NewGiantCalculator()
	//giantCalc.CalculateUpgradePath()
	giantCalc.GetNextUpgrade(73, 65)
}
