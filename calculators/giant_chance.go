package calculators

import (
	"capEndgame3Calculator/upgrade_data"
	"fmt"
)

const DoubleStrike = 2
const TripleStrike = 3
const QuadrupleStrike = 4
const QuintupleStrike = 5
const GiantLuck = 9000
const NoChange = 0
const x2Overclock = 1.4
const x3Overclock = 1.6
const x4Overclock = 1.8
const x5Overclock = 2.0
const giantLuckOverclock = 1.5

type upgradeCostList map[int]int
type strikeUpgrades map[int]int
type OverclockConfig map[int]bool

func NewOverclockConfig(x2, x3, x4, x5, giant bool) OverclockConfig {
	return OverclockConfig{
		DoubleStrike:    x2,
		TripleStrike:    x3,
		QuadrupleStrike: x4,
		QuintupleStrike: x5,
		GiantLuck:       giant,
	}
}

type GiantCalculator struct {
	strikeUpgrades               strikeUpgrades
	strikePrices                 map[int]upgradeCostList
	giantLuckUpgrade             int
	giantLuckPrices              upgradeCostList
	overclocks                   OverclockConfig
	achievementGiantLuckModifier float64
	runeGiantLuckModifier        float64
}

func NewGiantCalculator(ocConfig OverclockConfig, achievementModifier, runeModifier float64, strikeLevels strikeUpgrades, giantLuckLevel int) GiantCalculator {
	if achievementModifier < 1 {
		achievementModifier = 1
	}
	if runeModifier < 1 {
		runeModifier = 1
	}

	return GiantCalculator{
		strikeUpgrades: strikeLevels,
		strikePrices: map[int]upgradeCostList{
			DoubleStrike:    upgrade_data.GetStrikePrices(),
			TripleStrike:    upgrade_data.GetStrikePrices(),
			QuadrupleStrike: upgrade_data.GetStrikePrices(),
			QuintupleStrike: upgrade_data.GetStrikePrices(),
		},
		giantLuckUpgrade:             giantLuckLevel,
		giantLuckPrices:              upgrade_data.GetGiantLuckPrices(),
		overclocks:                   ocConfig,
		achievementGiantLuckModifier: achievementModifier,
		runeGiantLuckModifier:        runeModifier,
	}
}

func (gc *GiantCalculator) GetNextUpgrade() string {
	nextUpgrade := gc.findNextUpgrade()
	if nextUpgrade == GiantLuck {
		return "giant luck"
	} else {
		return fmt.Sprintf("upgrade x%d strike", nextUpgrade)
	}
}

func (gc *GiantCalculator) CalculateUpgradePath() {
	for {
		if gc.findNextUpgrade() == NoChange {
			return
		}

		nextUpgrade := gc.findNextUpgrade()
		if nextUpgrade == GiantLuck {
			fmt.Println("upgrade giant luck")
			gc.giantLuckUpgrade++
		} else {
			gc.strikeUpgrades[nextUpgrade]++
			fmt.Println(fmt.Sprintf("upgrade x%d strike", nextUpgrade))
		}

		fmt.Println(fmt.Sprintf("giant chance after upgrade: %.10f", gc.calculateGiantRollChance(NoChange)))
		fmt.Println(gc.strikeUpgrades, gc.giantLuckUpgrade)
	}
}

func (gc *GiantCalculator) CalculateChancePerSTrike(firstStrikeChance float64) float64 {
	chance := gc.calculateGiantRollChance(0)
	chance *= gc.achievementGiantLuckModifier
	chance *= gc.runeGiantLuckModifier
	chance *= firstStrikeChance

	return chance
}

func (gc *GiantCalculator) findNextUpgrade() int {
	if gc.getRequiredFirstUpgrade() != NoChange {
		return gc.getRequiredFirstUpgrade()
	}

	strikeChoices := gc.listPossibleStrikeUpgrades()
	if len(strikeChoices) == 0 && gc.giantLuckUpgrade == len(gc.giantLuckPrices) {
		return NoChange
	}

	currentGiantChance := gc.calculateGiantRollChance(NoChange)
	bestStrikeUpgrade := NoChange
	bestStrikeGain := float64(0)
	for _, strike := range strikeChoices {
		chanceGain := gc.calculateGiantRollChance(strike) - currentGiantChance
		upgradeCost := gc.strikePrices[strike][gc.strikeUpgrades[strike]+1]
		gain := chanceGain / float64(upgradeCost)
		if gain > bestStrikeGain {
			bestStrikeUpgrade = strike
			bestStrikeGain = gain
		}
	}

	giantLuckGain := gc.calculateGiantRollChance(GiantLuck)
	upgradeCost := gc.giantLuckPrices[gc.giantLuckUpgrade+1]
	gain := giantLuckGain / float64(upgradeCost)
	if gain > bestStrikeGain {
		return GiantLuck
	}

	return bestStrikeUpgrade
}

func (gc *GiantCalculator) getRequiredFirstUpgrade() int {
	if gc.strikeUpgrades[DoubleStrike] == 0 {
		return DoubleStrike
	}
	if gc.strikeUpgrades[TripleStrike] == 0 {
		return TripleStrike
	}
	if gc.strikeUpgrades[QuadrupleStrike] == 0 {
		return QuadrupleStrike
	}
	if gc.strikeUpgrades[QuintupleStrike] == 0 {
		return QuintupleStrike
	}
	if gc.giantLuckUpgrade == 0 {
		return GiantLuck
	}

	return NoChange
}

func (gc *GiantCalculator) listPossibleStrikeUpgrades() []int {
	strikeChoices := make([]int, 0)
	if gc.strikeUpgrades[DoubleStrike] < len(gc.strikePrices[DoubleStrike]) {
		strikeChoices = append(strikeChoices, DoubleStrike)
	}
	if gc.strikeUpgrades[TripleStrike] < len(gc.strikePrices[TripleStrike]) {
		strikeChoices = append(strikeChoices, TripleStrike)
	}
	if gc.strikeUpgrades[QuadrupleStrike] < len(gc.strikePrices[QuadrupleStrike]) {
		strikeChoices = append(strikeChoices, QuadrupleStrike)
	}
	if gc.strikeUpgrades[QuintupleStrike] < len(gc.strikePrices[QuintupleStrike]) {
		strikeChoices = append(strikeChoices, QuintupleStrike)
	}

	return strikeChoices
}

func (gc *GiantCalculator) calculateGiantRollChance(incrementedChance int) float64 {
	doubleChance := float64(gc.strikeUpgrades[DoubleStrike]) * upgrade_data.PerStepStrikeImprovement
	tripleChance := float64(gc.strikeUpgrades[TripleStrike]) * upgrade_data.PerStepStrikeImprovement
	quadrupleChance := float64(gc.strikeUpgrades[QuadrupleStrike]) * upgrade_data.PerStepStrikeImprovement
	quintupleChance := float64(gc.strikeUpgrades[QuintupleStrike]) * upgrade_data.PerStepStrikeImprovement
	giantLuckChance := float64(gc.giantLuckUpgrade) * upgrade_data.PerStepGiantLuckImprovement

	switch incrementedChance {
	case DoubleStrike:
		doubleChance += upgrade_data.PerStepStrikeImprovement
	case TripleStrike:
		tripleChance += upgrade_data.PerStepStrikeImprovement
	case QuadrupleStrike:
		quadrupleChance += upgrade_data.PerStepStrikeImprovement
	case QuintupleStrike:
		quintupleChance += upgrade_data.PerStepStrikeImprovement
	case GiantLuck:
		giantLuckChance += upgrade_data.PerStepGiantLuckImprovement
	}

	if gc.overclocks[DoubleStrike] {
		doubleChance *= x2Overclock
	}
	if gc.overclocks[TripleStrike] {
		tripleChance *= x3Overclock
	}
	if gc.overclocks[QuadrupleStrike] {
		quadrupleChance *= x4Overclock
	}
	if gc.overclocks[QuintupleStrike] {
		quintupleChance *= x5Overclock
	}
	if gc.overclocks[GiantLuck] {
		giantLuckChance *= giantLuckOverclock
	}

	return doubleChance * tripleChance * quadrupleChance * quintupleChance * giantLuckChance
}
