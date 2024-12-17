package main

import (
	"capEndgame3Calculator/calculators"
	"capEndgame3Calculator/character_config"
	"fmt"
	"golang.org/x/text/message"
	"time"
)

const HundredThousand = 100000
const Million = 1000000

func main() {
	shinyMods, giantCalc, stoneCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := loadSadinar()
	duration := time.Hour * 24

	fmt.Println("next giant chance upgrade should be", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))
	fmt.Println("next stone upgrade should be", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, nextCloneUpgradeCost))

	giantCalc.PrintProbabilityMedian(duration, shinyMods)
	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(duration)
	stoneCalc.PrintDamageChange(duration, shinyMods)
	p := message.NewPrinter(message.MatchLanguage("en"))
	fmt.Println(p.Sprintf("%d stones (%d genned and %d mined) gained in %v", gennedStones+minedStones, gennedStones, minedStones, duration))

	//fromScratchUpgradePath()
}

func loadSadinar() (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, int, int) {
	return character_config.ConfigureCalculators("./character_config/ascend_2_sadinar.json")
}

func fromScratchUpgradePath() {
	mm := calculators.NewMiningModifiers(
		.5,
		100,
		0,
		100,
		map[int]int{2: 0, 3: 0, 4: 0, 5: 0},
		0,
		map[int]float64{2: 0, 3: 0, 4: 0, 5: 0},
		false,
		false,
		false,
		false,
	)

	gl := calculators.NewGiantModifiers(
		1,
		1,
		1,
		1,
		false,
		false,
	)

	gc := calculators.NewGiantCalculator(mm, gl)
	fmt.Println(gc.CalculateUpgradePath())
}
