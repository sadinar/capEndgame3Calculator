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
	printer            *message.Printer
}

func NewGiantCalculator(mm MiningModifiers, lm GiantModifiers) Giant {
	return Giant{
		strikePrices:       upgrade_data.GetStrikePrices(),
		giantLuckPrices:    upgrade_data.GetGiantLuckPrices(),
		miningModifiers:    mm,
		giantLuckModifiers: lm,
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

func (gc *Giant) CalculateUpgradePath() string {
	output := "------------------------------------------------------------\n"
	output += "| x2 | x3 | x4 | x5 | giant |    chance/hit   | stone cost\n"
	originalStrikeUpgrades := strikeUpgrades{
		DoubleStrike:    gc.miningModifiers.StrikeUpgrades[DoubleStrike],
		TripleStrike:    gc.miningModifiers.StrikeUpgrades[TripleStrike],
		QuadrupleStrike: gc.miningModifiers.StrikeUpgrades[QuadrupleStrike],
		QuintupleStrike: gc.miningModifiers.StrikeUpgrades[QuintupleStrike],
	}
	originalStrikeOdds := strikeOdds{
		DoubleStrike:    gc.miningModifiers.StrikeOdds[DoubleStrike],
		TripleStrike:    gc.miningModifiers.StrikeOdds[TripleStrike],
		QuadrupleStrike: gc.miningModifiers.StrikeOdds[QuadrupleStrike],
		QuintupleStrike: gc.miningModifiers.StrikeOdds[QuintupleStrike],
	}
	originalGiantLuckLevel := gc.miningModifiers.GiantLuckLevel
	originalGiantOdds := gc.miningModifiers.GiantOdds

	for {
		if gc.findCartUpgrade() == NoChange {
			gc.miningModifiers.StrikeUpgrades = originalStrikeUpgrades
			gc.miningModifiers.StrikeOdds = originalStrikeOdds
			gc.miningModifiers.GiantLuckLevel = originalGiantLuckLevel
			gc.miningModifiers.GiantOdds = originalGiantOdds

			return output
		}

		nextUpgrade := gc.findCartUpgrade()
		if nextUpgrade == GiantLuck {
			gc.miningModifiers.GiantLuckLevel++
		} else {
			gc.miningModifiers.StrikeUpgrades[nextUpgrade]++
		}

		modifiedStrikeOdds := strikeOdds{DoubleStrike: 0, TripleStrike: 0, QuadrupleStrike: 0, QuintupleStrike: 0}
		switch nextUpgrade {
		case DoubleStrike:
			increasedOdds := gc.calculateIncreasedDoubleStrikeOdds()
			modifiedStrikeOdds[DoubleStrike] = increasedOdds

			if gc.miningModifiers.StrikeUpgrades[TripleStrike] == 0 {
				gc.miningModifiers.StrikeOdds = modifiedStrikeOdds
				continue
			}

			increasedOdds = gc.miningModifiers.StrikeOdds[TripleStrike] / gc.miningModifiers.StrikeOdds[DoubleStrike] * increasedOdds
			modifiedStrikeOdds[TripleStrike] = increasedOdds

			increasedOdds = gc.miningModifiers.StrikeOdds[QuadrupleStrike] / gc.miningModifiers.StrikeOdds[TripleStrike] * increasedOdds
			modifiedStrikeOdds[QuadrupleStrike] = increasedOdds

			increasedOdds = gc.miningModifiers.StrikeOdds[QuintupleStrike] / gc.miningModifiers.StrikeOdds[QuadrupleStrike] * increasedOdds
			modifiedStrikeOdds[QuintupleStrike] = increasedOdds

			gc.miningModifiers.StrikeOdds = modifiedStrikeOdds
			gc.miningModifiers.GiantOdds = gc.getOriginalGiantOdds() * increasedOdds
		case TripleStrike:
			modifiedStrikeOdds[DoubleStrike] = gc.miningModifiers.StrikeOdds[DoubleStrike]

			increasedOdds := gc.calculateIncreasedTripleStrikeOdds() * gc.miningModifiers.StrikeOdds[DoubleStrike]
			modifiedStrikeOdds[TripleStrike] = increasedOdds

			if gc.miningModifiers.StrikeUpgrades[QuadrupleStrike] == 0 {
				gc.miningModifiers.StrikeOdds = modifiedStrikeOdds
				continue
			}

			increasedOdds = gc.miningModifiers.StrikeOdds[QuadrupleStrike] / gc.miningModifiers.StrikeOdds[TripleStrike] * increasedOdds
			modifiedStrikeOdds[QuadrupleStrike] = increasedOdds

			increasedOdds = gc.miningModifiers.StrikeOdds[QuintupleStrike] / gc.miningModifiers.StrikeOdds[QuadrupleStrike] * increasedOdds
			modifiedStrikeOdds[QuintupleStrike] = increasedOdds

			gc.miningModifiers.StrikeOdds = modifiedStrikeOdds
			gc.miningModifiers.GiantOdds = gc.getOriginalGiantOdds() * increasedOdds
		case QuadrupleStrike:
			modifiedStrikeOdds[DoubleStrike] = gc.miningModifiers.StrikeOdds[DoubleStrike]
			modifiedStrikeOdds[TripleStrike] = gc.miningModifiers.StrikeOdds[TripleStrike]

			increasedOdds := gc.calculateIncreasedQuadrupleStrikeOdds() * gc.miningModifiers.StrikeOdds[TripleStrike]
			modifiedStrikeOdds[QuadrupleStrike] = increasedOdds

			if gc.miningModifiers.StrikeUpgrades[QuintupleStrike] == 0 {
				gc.miningModifiers.StrikeOdds = modifiedStrikeOdds
				continue
			}

			increasedOdds = gc.miningModifiers.StrikeOdds[QuintupleStrike] / gc.miningModifiers.StrikeOdds[QuadrupleStrike] * increasedOdds
			modifiedStrikeOdds[QuintupleStrike] = increasedOdds

			gc.miningModifiers.StrikeOdds = modifiedStrikeOdds
			gc.miningModifiers.GiantOdds = gc.getOriginalGiantOdds() * increasedOdds
		case QuintupleStrike:
			increasedOdds := gc.calculateIncreasedQuintupleStrikeOdds() * gc.miningModifiers.StrikeOdds[QuadrupleStrike]
			gc.miningModifiers.StrikeOdds[QuintupleStrike] = increasedOdds

			if gc.miningModifiers.GiantLuckLevel == 0 {
				continue
			}

			gc.miningModifiers.GiantOdds = gc.getOriginalGiantOdds() * increasedOdds
		case GiantLuck:
			gc.miningModifiers.GiantOdds = gc.calculateIncreasedGiantOdds() * gc.miningModifiers.StrikeOdds[QuintupleStrike]
		}

		output += fmt.Sprintf(
			"|%03d |%03d |%03d |%03d |%03d    | %.12f%% | %d\n",
			gc.miningModifiers.StrikeUpgrades[DoubleStrike],
			gc.miningModifiers.StrikeUpgrades[TripleStrike],
			gc.miningModifiers.StrikeUpgrades[QuadrupleStrike],
			gc.miningModifiers.StrikeUpgrades[QuintupleStrike],
			gc.miningModifiers.GiantLuckLevel,
			gc.calculateGiantChance(0)*100,
			gc.GetUpgradeCost(),
		)
	}
}

func (gc *Giant) CalculateChancePerStrike() float64 {
	chance := gc.calculateGiantChance(0)
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

func (gc *Giant) PrintProbabilityMedian(duration time.Duration, sMods ShinyModifiers) (int, int, float64) {
	dailyAttempts := gc.getEggMineAttempts(duration)
	successProbability := gc.CalculateChancePerStrike()
	successCount, _ := FindReasonableSuccessCeiling(dailyAttempts, successProbability)
	probabilityList := gc.getProbabilityList(successCount, dailyAttempts, successProbability)

	medianIndex, medianProbability := gc.findProbabilityBreakpoint(probabilityList, 0.5)
	shinyOdds := sMods.CalculateShinyOdds()
	if gc.giantLuckModifiers.shinyOverclocked {
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

	return medianIndex, shinyCount, medianProbability
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
	nonSpeedGiantCount := float64(nonSpeedMineStrikes) * gc.calculateGiantChance(bestNonSpeedUpgrade)
	nonSpeedEfficiency := nonSpeedGiantCount / float64(nonSpeedUpgradeCost)

	gc.miningModifiers.MineSpeed += upgrade_data.PerStepSpeedImprovement
	speedMineStrikes := gc.getEggMineAttempts(time.Hour * 24 * 5)
	speedGiantCount := float64(speedMineStrikes) * gc.calculateGiantChance(NoChange)
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

	currentGiantChance := gc.calculateGiantChance(NoChange)
	bestStrikeUpgrade := NoChange
	bestStrikeGain := float64(0)
	for _, strike := range strikeChoices {
		chanceGain := gc.calculateGiantChance(strike) - currentGiantChance

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

	giantLuckGain := gc.calculateGiantChance(GiantLuck) - currentGiantChance
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

func (gc *Giant) calculateGiantChance(incrementedChance int) float64 {
	if incrementedChance == NoChange {
		return gc.miningModifiers.GiantOdds
	}

	switch incrementedChance {
	case DoubleStrike:
		return gc.propagateModifiedChance(DoubleStrike, gc.calculateIncreasedDoubleStrikeOdds())
	case TripleStrike:
		return gc.miningModifiers.StrikeOdds[DoubleStrike] * gc.propagateModifiedChance(TripleStrike, gc.calculateIncreasedTripleStrikeOdds())
	case QuadrupleStrike:
		return gc.miningModifiers.StrikeOdds[TripleStrike] * gc.propagateModifiedChance(QuadrupleStrike, gc.calculateIncreasedQuadrupleStrikeOdds())
	case QuintupleStrike:
		return gc.miningModifiers.StrikeOdds[QuadrupleStrike] * gc.propagateModifiedChance(QuintupleStrike, gc.calculateIncreasedQuintupleStrikeOdds())
	case GiantLuck:
		return gc.miningModifiers.GiantOdds / gc.getOriginalGiantOdds() * gc.calculateIncreasedGiantOdds()
	default:
		panic("unknown calculate giant chance option")
	}
}

func (gc *Giant) calculateIncreasedDoubleStrikeOdds() float64 {
	increment := upgrade_data.PerStepStrikeImprovement
	if gc.miningModifiers.x2Overclock {
		increment *= X2OverclockMultiplier
	}
	increasedDoubleOdds := gc.miningModifiers.StrikeOdds[DoubleStrike] + increment

	return increasedDoubleOdds
}

func (gc *Giant) calculateIncreasedTripleStrikeOdds() float64 {
	increment := upgrade_data.PerStepStrikeImprovement
	if gc.miningModifiers.x3Overclock {
		increment *= X3OverclockMultiplier
	}
	originalTripleOdds := gc.miningModifiers.StrikeOdds[TripleStrike] / gc.miningModifiers.StrikeOdds[DoubleStrike]
	increasedTripleOdds := originalTripleOdds + increment

	return increasedTripleOdds
}

func (gc *Giant) calculateIncreasedQuadrupleStrikeOdds() float64 {
	increment := upgrade_data.PerStepStrikeImprovement
	if gc.miningModifiers.x4Overclock {
		increment *= X4OverclockMultiplier
	}
	originalQuadOdds := gc.miningModifiers.StrikeOdds[QuadrupleStrike] / gc.miningModifiers.StrikeOdds[TripleStrike]
	increasedQuadOdds := originalQuadOdds + increment

	return increasedQuadOdds
}

func (gc *Giant) calculateIncreasedQuintupleStrikeOdds() float64 {
	increment := upgrade_data.PerStepStrikeImprovement
	if gc.miningModifiers.x5Overclock {
		increment *= X5OverclockMultiplier
	}
	originalQuintOdds := gc.miningModifiers.StrikeOdds[QuintupleStrike] / gc.miningModifiers.StrikeOdds[QuadrupleStrike]
	increasedQuintOdds := originalQuintOdds + increment

	return increasedQuintOdds
}

func (gc *Giant) calculateIncreasedGiantOdds() float64 {
	modifiers := gc.giantLuckModifiers.t7GiantLuck * gc.giantLuckModifiers.t8GiantLuck
	modifiers *= gc.giantLuckModifiers.rune * gc.giantLuckModifiers.achievement
	if gc.giantLuckModifiers.luckOverclocked {
		modifiers *= GiantLuckOverclockMultiplier
	}
	increasedGiantOdds := gc.getOriginalGiantOdds() + upgrade_data.PerStepGiantLuckImprovement*modifiers

	return increasedGiantOdds
}

func (gc *Giant) propagateModifiedChance(modifiedStrike int, modifiedStrikeOdds float64) float64 {
	if modifiedStrike == QuintupleStrike {
		return gc.getOriginalGiantOdds() * modifiedStrikeOdds
	}

	nextStrikeBaseOdds := gc.miningModifiers.StrikeOdds[modifiedStrike+1] / gc.miningModifiers.StrikeOdds[modifiedStrike]
	nextStrikeIncreasedOdds := nextStrikeBaseOdds * modifiedStrikeOdds

	return gc.propagateModifiedChance(modifiedStrike+1, nextStrikeIncreasedOdds)
}

func (gc *Giant) getOriginalGiantOdds() float64 {
	return gc.miningModifiers.GiantOdds / gc.miningModifiers.StrikeOdds[QuintupleStrike]
}

func (gc *Giant) getProbabilityList(successCount, trials uint64, successProbability float64) map[int]float64 {
	probabilityList := make(map[int]float64, 0)
	for i := 0; i <= int(successCount); i++ {
		chance := BinomialProbability(trials, uint64(i), successProbability)
		probabilityList[i] = chance
	}

	return probabilityList
}
