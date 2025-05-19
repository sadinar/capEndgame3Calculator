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
	shinyMods, giantCalc, stoneCalc, bonusPetScoreCalc, nextSpeedUpgradeCost, nextCloneUpgradeCost := loadSadinar()
	duration := time.Hour * 24

	fmt.Println("next giant chance upgrade should be", giantCalc.GetNextUpgrade(nextSpeedUpgradeCost))

	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(duration)
	if minedStones > 0 {
		fmt.Println("next stone upgrade should be", stoneCalc.FindNextUpgrade(nextSpeedUpgradeCost, nextCloneUpgradeCost))
		giantCalc.PrintProbabilityMedian(duration, shinyMods)
	} else {
		fmt.Println("mining instead would result in:")
		fmt.Print("    ")
		giantCalc.PrintProbabilityMedian(duration, shinyMods)
	}
	stoneCalc.PrintDamageChange(duration, shinyMods)
	p := message.NewPrinter(message.MatchLanguage("en"))
	fmt.Println(p.Sprintf("%d stones (%d genned and %d mined) gained in %v", gennedStones+minedStones, gennedStones, minedStones, duration))
	_, gennedMythics, _ := stoneCalc.CalculateTotalGeneratedPets(duration)
	petScore := minedStones + int(gennedMythics) + bonusPetScoreCalc.BonusPetScore(gennedMythics)
	fmt.Println(p.Sprintf("Pet score gained: %d", petScore))

	//fromScratchUpgradePath()
}

func loadSadinar() (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, calculators.PetScore, int, int) {
	//return character_config.ConfigureCalculators("./character_config/ascend_3_crank_sadinar.json")
	return character_config.ConfigureCalculators("./character_config/ascend_3_mine_sadinar.json")
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

	al := calculators.NewAscensionModifiers(0, 0)

	gc := calculators.NewGiantCalculator(mm, gl, al)
	fmt.Println(gc.CalculateUpgradePath())
}
