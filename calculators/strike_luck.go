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

type StrikeLuck struct {
	strikeUpgrades   strikeUpgrades
	strikePrices     map[int]upgradeCostList
	giantLuckUpgrade int
	giantLuckPrices  upgradeCostList
}

func New() StrikeLuck {
	return StrikeLuck{
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

func (sl *StrikeLuck) CalculateUpgradePath() {
	for {
		if sl.findNextUpgrade() == 0 {
			return
		}

		nextUpgrade := sl.findNextUpgrade()
		if nextUpgrade == GiantLuck {
			fmt.Println("upgrade giant luck")
		} else {
			fmt.Println(fmt.Sprintf("upgrade x%d strike", nextUpgrade))
		}

		if nextUpgrade == GiantLuck {
			sl.giantLuckUpgrade++
		} else {
			sl.strikeUpgrades[nextUpgrade]++
		}

		fmt.Println(sl.strikeUpgrades, sl.giantLuckUpgrade)
	}
}

func (sl *StrikeLuck) findNextUpgrade() int {
	if sl.getRequiredFirstUpgrade() != NoChange {
		return sl.getRequiredFirstUpgrade()
	}

	strikeChoices := sl.listPossibleStrikeUpgrades()
	if len(strikeChoices) == 0 && sl.giantLuckUpgrade == len(sl.giantLuckPrices) {
		return NoChange
	}

	currentGiantChance := sl.calculateGiantRollChance(NoChange)
	bestUpgrade := NoChange
	bestGain := float64(0)
	for _, strike := range strikeChoices {
		chanceGain := sl.calculateGiantRollChance(strike) - currentGiantChance
		upgradeCost := sl.strikePrices[strike][sl.strikeUpgrades[strike]+1]
		gain := chanceGain / float64(upgradeCost)
		if gain > bestGain {
			bestUpgrade = strike
			bestGain = gain
		}
	}

	giantLuckGain := sl.calculateGiantRollChance(GiantLuck)
	upgradeCost := sl.giantLuckPrices[sl.giantLuckUpgrade+1]
	gain := giantLuckGain / float64(upgradeCost)
	if gain > bestGain {
		return GiantLuck
	}

	return bestUpgrade
}

func (sl *StrikeLuck) getRequiredFirstUpgrade() int {
	if sl.strikeUpgrades[SingleStrike] == 0 {
		return SingleStrike
	}
	if sl.strikeUpgrades[DoubleStrike] == 0 {
		return DoubleStrike
	}
	if sl.strikeUpgrades[TripleStrike] == 0 {
		return TripleStrike
	}
	if sl.strikeUpgrades[QuadrupleStrike] == 0 {
		return QuadrupleStrike
	}
	if sl.strikeUpgrades[QuintupleStrike] == 0 {
		return QuintupleStrike
	}
	if sl.giantLuckUpgrade == 0 {
		return GiantLuck
	}

	return NoChange
}

func (sl *StrikeLuck) listPossibleStrikeUpgrades() []int {
	strikeChoices := make([]int, 0)
	if sl.strikeUpgrades[SingleStrike] < len(sl.strikePrices[SingleStrike]) {
		strikeChoices = append(strikeChoices, SingleStrike)
	}
	if sl.strikeUpgrades[DoubleStrike] < len(sl.strikePrices[DoubleStrike]) {
		strikeChoices = append(strikeChoices, DoubleStrike)
	}
	if sl.strikeUpgrades[TripleStrike] < len(sl.strikePrices[TripleStrike]) {
		strikeChoices = append(strikeChoices, TripleStrike)
	}
	if sl.strikeUpgrades[QuadrupleStrike] < len(sl.strikePrices[QuadrupleStrike]) {
		strikeChoices = append(strikeChoices, QuadrupleStrike)
	}
	if sl.strikeUpgrades[QuintupleStrike] < len(sl.strikePrices[QuintupleStrike]) {
		strikeChoices = append(strikeChoices, QuintupleStrike)
	}

	return strikeChoices
}

func (sl *StrikeLuck) calculateGiantRollChance(incrementedChance int) float64 {
	singleChance := float64(sl.strikeUpgrades[SingleStrike]) * upgrade_data.PerStepStrikeImprovement
	doubleChance := float64(sl.strikeUpgrades[DoubleStrike]) * upgrade_data.PerStepStrikeImprovement
	tripleChance := float64(sl.strikeUpgrades[TripleStrike]) * upgrade_data.PerStepStrikeImprovement
	quadrupleChance := float64(sl.strikeUpgrades[QuadrupleStrike]) * upgrade_data.PerStepStrikeImprovement
	quintupleChance := float64(sl.strikeUpgrades[QuintupleStrike]) * upgrade_data.PerStepStrikeImprovement
	giantLuckChance := float64(sl.giantLuckUpgrade) * upgrade_data.PerStepGiantLuckImprovement

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
