package calculators

import (
	"capEndgame3Calculator/upgrade_data"
	"fmt"
	"time"
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
	mineSpeed                    float64
}

func NewGiantCalculator(ocConfig OverclockConfig, achievementModifier, runeModifier, mineSpeed float64, strikeLevels strikeUpgrades, giantLuckLevel int) GiantCalculator {
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
		mineSpeed:                    mineSpeed,
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
	fmt.Println("----------------------------------")
	fmt.Println("| x2 | x3 | x4 | x5 | giant luck |")
	for {
		if gc.findNextUpgrade() == NoChange {
			return
		}

		nextUpgrade := gc.findNextUpgrade()
		if nextUpgrade == GiantLuck {
			//fmt.Println("upgrade giant luck")
			gc.giantLuckUpgrade++
		} else {
			gc.strikeUpgrades[nextUpgrade]++
			//fmt.Println(fmt.Sprintf("upgrade x%d strike", nextUpgrade))
		}

		//fmt.Println(fmt.Sprintf("giant chance after upgrade: %.10f", gc.calculateBaseGiantChance(NoChange)))
		//fmt.Println(gc.strikeUpgrades, gc.giantLuckUpgrade)
		fmt.Println(
			fmt.Sprintf(
				"|%03d |%03d |%03d |%03d |%03d         |",
				gc.strikeUpgrades[DoubleStrike],
				gc.strikeUpgrades[TripleStrike],
				gc.strikeUpgrades[QuadrupleStrike],
				gc.strikeUpgrades[QuintupleStrike],
				gc.giantLuckUpgrade,
			),
		)
	}
}

func (gc *GiantCalculator) CalculateChancePerSTrike(firstStrikeChance float64) float64 {
	chance := gc.calculateBaseGiantChance(0)

	if gc.overclocks[DoubleStrike] {
		chance *= x2Overclock
	}
	if gc.overclocks[TripleStrike] {
		chance *= x3Overclock
	}
	if gc.overclocks[QuadrupleStrike] {
		chance *= x4Overclock
	}
	if gc.overclocks[QuintupleStrike] {
		chance *= x5Overclock
	}
	if gc.overclocks[GiantLuck] {
		chance *= giantLuckOverclock
	}

	chance *= gc.achievementGiantLuckModifier
	chance *= gc.runeGiantLuckModifier
	chance *= firstStrikeChance

	return chance
}

func (gc *GiantCalculator) PrintProbabilityDistribution(duration time.Duration, firstStrikeChance float64) {
	dailyAttempts := gc.getEggsMined(duration)
	successProbability := gc.CalculateChancePerSTrike(firstStrikeChance)
	successCount, consumedProbabilitySpace := FindReasonableSuccessCeiling(dailyAttempts, successProbability)
	reportedProbabilitySpace := 0.0
	medianProbability := 0.0
	medianSuccesses := 0
	subFiveProbabilitySpace := 0.0
	hasPrintedSubFiveTotal := false
	for i := 0; i <= int(successCount); i++ {
		chance := BinomialProbability(dailyAttempts, uint64(i), successProbability)
		if subFiveProbabilitySpace < .05 && chance < .10 {
			subFiveProbabilitySpace += chance
		} else if !hasPrintedSubFiveTotal {
			hasPrintedSubFiveTotal = true
			fmt.Println(fmt.Sprintf("%d-%d: %.12f%%", 0, i, subFiveProbabilitySpace*100))
		} else {
			fmt.Println(fmt.Sprintf("%d: %.12f%%", i, chance*100))
		}

		if reportedProbabilitySpace < .5 {
			reportedProbabilitySpace += chance
		}
		if medianSuccesses == 0 && reportedProbabilitySpace >= .5 {
			medianSuccesses = i
			medianProbability = chance
		}
	}
	fmt.Println(fmt.Sprintf("%d+: %.12f%%", successCount+1, (1-consumedProbabilitySpace)*100))
	fmt.Println(fmt.Sprintf("Median giant successes: %d @ %.12f%%", medianSuccesses, medianProbability*100))
	fmt.Println(fmt.Sprintf("%.4f%% chance %d or fewer giants and %.4f%% chance over %d giants in %v", reportedProbabilitySpace*100, medianSuccesses, (1-reportedProbabilitySpace)*100, medianSuccesses, duration))
}

func (gc *GiantCalculator) getEggsMined(duration time.Duration) uint64 {
	return uint64(duration.Seconds() * gc.mineSpeed)
}

func (gc *GiantCalculator) findNextUpgrade() int {
	if gc.getRequiredFirstUpgrade() != NoChange {
		return gc.getRequiredFirstUpgrade()
	}

	strikeChoices := gc.listPossibleStrikeUpgrades()
	if len(strikeChoices) == 0 && gc.giantLuckUpgrade == len(gc.giantLuckPrices) {
		return NoChange
	}

	currentGiantChance := gc.calculateBaseGiantChance(NoChange)
	bestStrikeUpgrade := NoChange
	bestStrikeGain := float64(0)
	for _, strike := range strikeChoices {
		chanceGain := gc.calculateBaseGiantChance(strike) - currentGiantChance

		if gc.strikePrices[strike][gc.strikeUpgrades[strike]+1] == 0 {
			panic("ran out of strike prices!")
		}

		upgradeCost := gc.strikePrices[strike][gc.strikeUpgrades[strike]+1]
		gain := chanceGain / float64(upgradeCost)
		if gain > bestStrikeGain {
			bestStrikeUpgrade = strike
			bestStrikeGain = gain
		}
	}

	giantLuckGain := gc.calculateBaseGiantChance(GiantLuck)

	if gc.giantLuckPrices[gc.giantLuckUpgrade+1] == 0 {
		panic("ran out of giant luck prices!")
	}

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

func (gc *GiantCalculator) calculateBaseGiantChance(incrementedChance int) float64 {
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

	return doubleChance * tripleChance * quadrupleChance * quintupleChance * giantLuckChance
}
