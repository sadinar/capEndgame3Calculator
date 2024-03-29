package calculators

import (
	"capEndgame3Calculator/upgrade_data"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"time"
)

const DoubleStrike = 2
const TripleStrike = 3
const QuadrupleStrike = 4
const QuintupleStrike = 5
const GiantLuck = 9000
const NoChange = 0

type upgradeCostList map[int]int
type strikeUpgrades map[int]int
type strikeOdds map[int]float64

type Giant struct {
	strikePrices       upgradeCostList
	giantLuckPrices    upgradeCostList
	miningModifiers    MiningModifiers
	giantLuckModifiers GiantModifiers
	shinyLuckOverclock bool
	printer            *message.Printer
}

func NewGiantCalculator(mm MiningModifiers, lm GiantModifiers, giantShinyLuckOverclocked bool) Giant {
	return Giant{
		strikePrices:       upgrade_data.GetStrikePrices(),
		giantLuckPrices:    upgrade_data.GetGiantLuckPrices(),
		miningModifiers:    mm,
		giantLuckModifiers: lm,
		shinyLuckOverclock: giantShinyLuckOverclocked,
		printer:            message.NewPrinter(language.English),
	}
}

func (gc *Giant) GetNextUpgrade(speedUpgradeCost int) string {
	bestUpgrade := gc.findCartUpgrade()
	upgradeCost := 0
	strikeCosts := upgrade_data.GetStrikePrices()
	giantLuckCosts := upgrade_data.GetGiantLuckPrices()

	switch bestUpgrade {
	case DoubleStrike:
		upgradeCost = strikeCosts[gc.miningModifiers.StrikeUpgrades[DoubleStrike]+1]
	case TripleStrike:
		upgradeCost = strikeCosts[gc.miningModifiers.StrikeUpgrades[TripleStrike]+1]
	case QuadrupleStrike:
		upgradeCost = strikeCosts[gc.miningModifiers.StrikeUpgrades[QuadrupleStrike]+1]
	case QuintupleStrike:
		upgradeCost = strikeCosts[gc.miningModifiers.StrikeUpgrades[QuintupleStrike]+1]
	case GiantLuck:
		upgradeCost = giantLuckCosts[gc.miningModifiers.GiantLuckLevel] + 1
	}

	if speedUpgradeCost != UpgradeComplete && gc.isSpeedBetterUpgrade(bestUpgrade, upgradeCost, speedUpgradeCost) {
		return "speed"
	}

	if bestUpgrade == NoChange {
		return "n/a"
	}

	if bestUpgrade == GiantLuck {
		return "giant luck"
	}

	return fmt.Sprintf("x%d strike", bestUpgrade)
}

func (gc *Giant) CalculateUpgradePath() {
	fmt.Println("------------------------------------------------------------")
	fmt.Println("| x2 | x3 | x4 | x5 | giant |    chance/hit   | stone cost")
	for {
		if gc.findCartUpgrade() == NoChange {
			return
		}

		nextUpgrade := gc.findCartUpgrade()
		if nextUpgrade == GiantLuck {
			gc.miningModifiers.GiantLuckLevel++
		} else {
			gc.miningModifiers.StrikeUpgrades[nextUpgrade]++
		}

		fmt.Println(
			fmt.Sprintf(
				"|%03d |%03d |%03d |%03d |%03d    | %.12f%% | %d",
				gc.miningModifiers.StrikeUpgrades[DoubleStrike],
				gc.miningModifiers.StrikeUpgrades[TripleStrike],
				gc.miningModifiers.StrikeUpgrades[QuadrupleStrike],
				gc.miningModifiers.StrikeUpgrades[QuintupleStrike],
				gc.miningModifiers.GiantLuckLevel,
				gc.calculateBaseGiantChance(0)*100,
				gc.GetUpgradeCost(),
			),
		)
	}
}

func (gc *Giant) CalculateChancePerStrike() float64 {
	chance := gc.calculateBaseGiantChance(0)
	chance *= gc.miningModifiers.FirstStrike

	return chance
}

func (gc *Giant) PrintProbabilityDistribution(duration time.Duration) {
	dailyAttempts := gc.getEggMineAttempts(duration)
	successProbability := gc.CalculateChancePerStrike()
	successCount, consumedProbabilitySpace := FindReasonableSuccessCeiling(dailyAttempts, successProbability)
	probabilityList := gc.getProbabilityList(successCount, dailyAttempts, successProbability)

	fmt.Println(fmt.Sprintf("0: %.12f%%", probabilityList[0]*100))
	lowIndex, lowProbability := gc.findProbabilityBreakpoint(probabilityList, 0.05)
	msgPrefix := "1-"
	if lowIndex > 1 {
		fmt.Println(fmt.Sprintf("%s%d: %.12f%%", msgPrefix, lowIndex, lowProbability*100))
	} else {
		lowIndex = 0
	}

	for i := lowIndex + 1; i < len(probabilityList); i++ {
		fmt.Println(fmt.Sprintf("%d: %.12f%%", i, probabilityList[i]*100))
	}
	fmt.Println(fmt.Sprintf("%d+: %.12f%%", len(probabilityList), (1-consumedProbabilitySpace)*100))
}

func (gc *Giant) PrintProbabilityMedian(duration time.Duration, sMods ShinyModifiers) {
	dailyAttempts := gc.getEggMineAttempts(duration)
	successProbability := gc.CalculateChancePerStrike()
	successCount, _ := FindReasonableSuccessCeiling(dailyAttempts, successProbability)
	probabilityList := gc.getProbabilityList(successCount, dailyAttempts, successProbability)

	medianIndex, medianProbability := gc.findProbabilityBreakpoint(probabilityList, 0.5)
	shinyOdds := sMods.CalculateShinyOdds()
	if gc.shinyLuckOverclock {
		shinyOdds *= 1.5
	}

	shinyCount := int(float64(medianIndex) * shinyOdds)
	if shinyCount > medianIndex {
		shinyCount = medianIndex
	}

	fmt.Println(
		gc.printer.Sprintf("median of %d (%d shiny) giants: %.12f%% chance of %d or fewer giants in %v",
			medianIndex,
			shinyCount,
			medianProbability*100,
			medianIndex,
			duration,
		),
	)
}

func (gc *Giant) GetUpgradeCost() int {
	totalCost := 0
	for _, level := range gc.miningModifiers.StrikeUpgrades {
		for i := 1; i <= level; i++ {
			totalCost += gc.strikePrices[i]
		}
	}

	for i := 1; i <= gc.miningModifiers.GiantLuckLevel; i++ {
		totalCost += gc.giantLuckPrices[i]
	}

	return totalCost
}

func (gc *Giant) isSpeedBetterUpgrade(bestNonSpeedUpgrade, nonSpeedUpgradeCost, speedUpgradeCost int) bool {
	if bestNonSpeedUpgrade == NoChange && speedUpgradeCost != UpgradeComplete {
		return true
	}

	nonSpeedMineStrikes := gc.getEggMineAttempts(time.Hour * 24 * 5)
	nonSpeedGiantCount := float64(nonSpeedMineStrikes) * gc.calculateBaseGiantChance(bestNonSpeedUpgrade)
	nonSpeedEfficiency := nonSpeedGiantCount / float64(nonSpeedUpgradeCost)

	gc.miningModifiers.MineSpeed += upgrade_data.PerStepSpeedImprovement
	speedMineStrikes := gc.getEggMineAttempts(time.Hour * 24 * 5)
	speedGiantCount := float64(speedMineStrikes) * gc.calculateBaseGiantChance(NoChange)
	speedEfficiency := speedGiantCount / float64(speedUpgradeCost)
	gc.miningModifiers.MineSpeed -= upgrade_data.PerStepSpeedImprovement

	return speedEfficiency > nonSpeedEfficiency
}

func (gc *Giant) findProbabilityBreakpoint(probabilityList map[int]float64, breakPoint float64) (int, float64) {
	if probabilityList[0] >= 0.5 {
		return 0, probabilityList[0]
	}

	totalProbability := 0.0
	maxIncludedIndex := 0

	for i := 1; i < len(probabilityList); i++ {
		totalProbability += probabilityList[i]
		maxIncludedIndex = i

		if totalProbability >= breakPoint {
			return maxIncludedIndex, totalProbability
		}
	}

	return maxIncludedIndex, totalProbability
}

func (gc *Giant) getEggMineAttempts(duration time.Duration) uint64 {
	return uint64(duration.Seconds() * gc.miningModifiers.MineSpeed)
}

func (gc *Giant) findCartUpgrade() int {
	if gc.getRequiredFirstUpgrade() != NoChange {
		return gc.getRequiredFirstUpgrade()
	}

	strikeChoices := gc.listPossibleStrikeUpgrades()
	if len(strikeChoices) == 0 && gc.miningModifiers.GiantLuckLevel == len(gc.giantLuckPrices) {
		return NoChange
	}

	currentGiantChance := gc.calculateBaseGiantChance(NoChange)
	bestStrikeUpgrade := NoChange
	bestStrikeGain := float64(0)
	for _, strike := range strikeChoices {
		chanceGain := gc.calculateBaseGiantChance(strike) - currentGiantChance

		upgradeCost := gc.strikePrices[gc.miningModifiers.StrikeUpgrades[strike]+1]
		gain := chanceGain / float64(upgradeCost)
		if gain > bestStrikeGain {
			bestStrikeUpgrade = strike
			bestStrikeGain = gain
		}
	}

	if gc.miningModifiers.GiantLuckLevel == len(gc.giantLuckPrices) {
		return bestStrikeUpgrade
	}

	giantLuckGain := gc.calculateBaseGiantChance(GiantLuck) - currentGiantChance
	upgradeCost := gc.giantLuckPrices[gc.miningModifiers.GiantLuckLevel+1]
	gain := giantLuckGain / float64(upgradeCost)
	if gain > bestStrikeGain {
		return GiantLuck
	}

	return bestStrikeUpgrade
}

func (gc *Giant) getRequiredFirstUpgrade() int {
	if gc.miningModifiers.StrikeUpgrades[DoubleStrike] == 0 {
		return DoubleStrike
	}
	if gc.miningModifiers.StrikeUpgrades[TripleStrike] == 0 {
		return TripleStrike
	}
	if gc.miningModifiers.StrikeUpgrades[QuadrupleStrike] == 0 {
		return QuadrupleStrike
	}
	if gc.miningModifiers.StrikeUpgrades[QuintupleStrike] == 0 {
		return QuintupleStrike
	}
	if gc.miningModifiers.GiantLuckLevel == 0 {
		return GiantLuck
	}

	return NoChange
}

func (gc *Giant) listPossibleStrikeUpgrades() []int {
	strikeChoices := make([]int, 0)
	if gc.miningModifiers.StrikeUpgrades[DoubleStrike] < len(gc.strikePrices) {
		strikeChoices = append(strikeChoices, DoubleStrike)
	}
	if gc.miningModifiers.StrikeUpgrades[TripleStrike] < len(gc.strikePrices) {
		strikeChoices = append(strikeChoices, TripleStrike)
	}
	if gc.miningModifiers.StrikeUpgrades[QuadrupleStrike] < len(gc.strikePrices) {
		strikeChoices = append(strikeChoices, QuadrupleStrike)
	}
	if gc.miningModifiers.StrikeUpgrades[QuintupleStrike] < len(gc.strikePrices) {
		strikeChoices = append(strikeChoices, QuintupleStrike)
	}

	return strikeChoices
}

func (gc *Giant) calculateBaseGiantChance(incrementedChance int) float64 {
	if incrementedChance == NoChange {
		return gc.miningModifiers.GiantOdds
	}

	switch incrementedChance {
	case DoubleStrike:
		increment := upgrade_data.PerStepStrikeImprovement
		if gc.miningModifiers.x2Overclock {
			increment *= X2OverclockMultiplier
		}
		increasedDoubleOdds := gc.miningModifiers.StrikeOdds[DoubleStrike] + increment
		return gc.miningModifiers.GiantOdds / gc.miningModifiers.StrikeOdds[DoubleStrike] * increasedDoubleOdds
	case TripleStrike:
		increment := upgrade_data.PerStepStrikeImprovement
		if gc.miningModifiers.x3Overclock {
			increment *= X3OverclockMultiplier
		}
		originalTripleOdds := gc.miningModifiers.StrikeOdds[TripleStrike] / gc.miningModifiers.StrikeOdds[DoubleStrike]
		increasedTripleOdds := originalTripleOdds + increment
		return gc.miningModifiers.GiantOdds / (originalTripleOdds) * increasedTripleOdds
	case QuadrupleStrike:
		increment := upgrade_data.PerStepStrikeImprovement
		if gc.miningModifiers.x4Overclock {
			increment *= X4OverclockMultiplier
		}
		originalQuadOdds := gc.miningModifiers.StrikeOdds[QuadrupleStrike] / gc.miningModifiers.StrikeOdds[TripleStrike]
		increasedQuadOdds := originalQuadOdds + increment
		return gc.miningModifiers.GiantOdds / originalQuadOdds * increasedQuadOdds
	case QuintupleStrike:
		increment := upgrade_data.PerStepStrikeImprovement
		if gc.miningModifiers.x5Overclock {
			increment *= X5OverclockMultiplier
		}
		originalPentaOdds := gc.miningModifiers.StrikeOdds[QuintupleStrike] / gc.miningModifiers.StrikeOdds[QuadrupleStrike]
		increasedPentaOdds := originalPentaOdds + increment
		return gc.miningModifiers.GiantOdds / originalPentaOdds * increasedPentaOdds
	case GiantLuck:
		originalGiantOdds := gc.miningModifiers.GiantOdds / gc.miningModifiers.StrikeOdds[QuintupleStrike]
		modifiers := gc.giantLuckModifiers.t7GiantLuck * gc.giantLuckModifiers.t8GiantLuck * gc.giantLuckModifiers.rune * gc.giantLuckModifiers.achievement
		if gc.giantLuckModifiers.isOverclocked {
			modifiers *= GiantLuckOverclockMultiplier
		}
		increasedGiantOdds := originalGiantOdds + upgrade_data.PerStepGiantLuckImprovement*modifiers
		return gc.miningModifiers.GiantOdds / originalGiantOdds * increasedGiantOdds
	}

	panic("unknown calculate giant chance option")
}

func (gc *Giant) getProbabilityList(successCount, trials uint64, successProbability float64) map[int]float64 {
	probabilityList := make(map[int]float64, 0)
	for i := 0; i <= int(successCount); i++ {
		chance := BinomialProbability(trials, uint64(i), successProbability)
		probabilityList[i] = chance
	}

	return probabilityList
}
