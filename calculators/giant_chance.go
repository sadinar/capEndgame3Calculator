package calculators

import (
	"capEndgame3Calculator/upgrade_data"
	"fmt"
)

const SingleStrike = 1
const DoubleStrike = 2
const TripleStrike = 3
const QuadrupleStrike = 4
const QuintupleStrike = 5
const GiantLuck = 9000
const NoChange = 0

type upgradeCostList map[int]int
type strikeUpgrades map[int]int

type GiantCalculator struct {
	strikeUpgrades   strikeUpgrades
	strikePrices     map[int]upgradeCostList
	giantLuckUpgrade int
	giantLuckPrices  upgradeCostList
}

func NewGiantCalculator() GiantCalculator {
	return GiantCalculator{
		strikeUpgrades: strikeUpgrades{},
		strikePrices: map[int]upgradeCostList{
			SingleStrike:    upgrade_data.GetSingleStrikePrices(),
			DoubleStrike:    upgrade_data.GetDoubleStrikePrices(),
			TripleStrike:    upgrade_data.GetTripleStrikePrices(),
			QuadrupleStrike: upgrade_data.GetQuadrupleStrikePrices(),
			QuintupleStrike: upgrade_data.GetQuintupleStrikePrices(),
		},
		giantLuckPrices: upgrade_data.GetGiantLuckPrices(),
	}
}

func (gc *GiantCalculator) CalculateUpgradePath() {
	for {
		if gc.findNextUpgrade() == 0 {
			return
		}

		nextUpgrade := gc.findNextUpgrade()
		if nextUpgrade == GiantLuck {
			fmt.Println("upgrade giant luck")
		} else {
			fmt.Println(fmt.Sprintf("upgrade x%d strike", nextUpgrade))
		}

		if nextUpgrade == GiantLuck {
			gc.giantLuckUpgrade++
		} else {
			gc.strikeUpgrades[nextUpgrade]++
		}

		fmt.Println(fmt.Sprintf("giant chance after upgrade: %.10f", gc.calculateGiantRollChance(NoChange)))
		fmt.Println(gc.strikeUpgrades, gc.giantLuckUpgrade)
	}
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
	if gc.strikeUpgrades[SingleStrike] == 0 {
		return SingleStrike
	}
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
	if gc.strikeUpgrades[SingleStrike] < len(gc.strikePrices[SingleStrike]) {
		strikeChoices = append(strikeChoices, SingleStrike)
	}
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
	singleChance := float64(gc.strikeUpgrades[SingleStrike]) * upgrade_data.PerStepStrikeImprovement
	doubleChance := float64(gc.strikeUpgrades[DoubleStrike]) * upgrade_data.PerStepStrikeImprovement
	tripleChance := float64(gc.strikeUpgrades[TripleStrike]) * upgrade_data.PerStepStrikeImprovement
	quadrupleChance := float64(gc.strikeUpgrades[QuadrupleStrike]) * upgrade_data.PerStepStrikeImprovement
	quintupleChance := float64(gc.strikeUpgrades[QuintupleStrike]) * upgrade_data.PerStepStrikeImprovement
	giantLuckChance := float64(gc.giantLuckUpgrade) * upgrade_data.PerStepGiantLuckImprovement

	switch incrementedChance {
	case SingleStrike:
		singleChance += upgrade_data.PerStepStrikeImprovement
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

	return singleChance * doubleChance * tripleChance * quadrupleChance * quintupleChance * giantLuckChance
}
