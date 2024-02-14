package main

import (
	"capEndgame3Calculator/calculators"
	"fmt"
	"golang.org/x/text/message"
	"time"
)

const OneMillion = 1000000

func main() {
	ocConfig := calculators.NewOverclockConfig(true, true, true, true, true, false)
	userMods := calculators.NewUserModifiers(
		1.1,
		1.2,
		0.72,
		1,
		map[int]int{
			2: 74,
			3: 74,
			4: 74,
			5: 74,
		},
		70,
	)
	shinyMods := calculators.NewShinyModifiers(
		1.1,
		1.09,
		1.2,
		10,
		1,
		65*OneMillion,
	)
	duration := time.Hour * 24

	giantCalc := calculators.NewGiantCalculator(ocConfig, userMods)
	sc := calculators.NewStonesCalculator(
		userMods,
		calculators.RubyPick,
		0.43,
		calculators.MythicEgg,
		ocConfig,
	)

	fmt.Println("next giant chance upgrade should be", giantCalc.GetNextUpgrade())
	fmt.Println("next stone upgrade should be", sc.FindNextUpgrade(1800000))

	giantCalc.PrintProbabilityMedian(duration, shinyMods)
	gennedStones, minedStones := sc.CalculateCombinedStones(duration)
	sc.PrintDamageChange(duration, shinyMods)
	p := message.NewPrinter(message.MatchLanguage("en"))
	fmt.Println(p.Sprintf("%d stones (%d genned and %d mined) gained in %v", gennedStones+minedStones, gennedStones, minedStones, duration))
}
