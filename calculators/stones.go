package calculators

import (
	"capEndgame3Calculator/upgrade_data"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"time"
)

const MaxGenSpeed = 5.0
const MaxFactoryCalcify = 2.0
const RubyPickMiningBonus = 2.0
const PerLevelEggModifier = 0.50
const StoneOverclockModifier = 1.5
const CommonEgg = 1
const UncommonEgg = 2
const RareEgg = 3
const EpicEgg = 4
const LegendaryEgg = 5
const ProdigiousEgg = 6
const AscendedEgg = 7
const MythicEgg = 8
const PerLevelSpeedModifier = 0.1
const PerLevelCloneModifier = 0.001

type Stones struct {
	firstStrike       float64
	x2Strike          float64
	x3Strike          float64
	x4Strike          float64
	x5Strike          float64
	mineSpeed         float64
	miningStoneBonus  float64
	eggLuck           float64
	cloneLuck         float64
	calcifyChance     float64
	eggLevel          int
	stonesOverclocked bool
	recursiveClone    bool
	miningModifiers   MiningModifiers
	printer           *message.Printer
}

func NewStonesCalculator(mm MiningModifiers, miningStoneBonus, eggLuck, cloneLuck, calcifyChance float64, eggLevel int, ocConfig OverclockConfig, recursiveClone bool) Stones {
	sc := Stones{
		firstStrike:       mm.FirstStrike,
		x2Strike:          float64(mm.StrikeUpgrades[DoubleStrike]) * upgrade_data.PerStepStrikeImprovement,
		x3Strike:          float64(mm.StrikeUpgrades[TripleStrike]) * upgrade_data.PerStepStrikeImprovement,
		x4Strike:          float64(mm.StrikeUpgrades[QuadrupleStrike]) * upgrade_data.PerStepStrikeImprovement,
		x5Strike:          float64(mm.StrikeUpgrades[QuintupleStrike]) * upgrade_data.PerStepStrikeImprovement,
		mineSpeed:         mm.MineSpeed,
		miningStoneBonus:  1 + (miningStoneBonus / 100),
		eggLuck:           eggLuck,
		cloneLuck:         cloneLuck,
		calcifyChance:     1 + (calcifyChance / 100),
		eggLevel:          eggLevel,
		stonesOverclocked: ocConfig[StoneOverclockIndex],
		recursiveClone:    recursiveClone,
		miningModifiers:   mm,
		printer:           message.NewPrinter(language.English),
	}

	if ocConfig[QuintupleStrike] {
		sc.x5Strike *= 2
	}

	if ocConfig[QuadrupleStrike] {
		sc.x4Strike *= 1.8
	}

	if ocConfig[TripleStrike] {
		sc.x3Strike *= 1.6
	}

	if ocConfig[DoubleStrike] {
		sc.x2Strike *= 1.4
	}

	return sc
}

func (sc *Stones) CalculateCombinedStones(period time.Duration) (generatedStones, minedStones int) {
	return sc.CalculateGeneratedStones(period), sc.CalculateMinedStones(period)
}

func (sc *Stones) CalculateGeneratedStones(period time.Duration) int {
	if period < time.Second {
		return 0
	}

	_, totalMythics, _ := sc.calculateTotalGeneratedPets(period)

	return int(totalMythics * sc.calcifyChance)
}

func (sc *Stones) PrintDamageChange(period time.Duration, sMods ShinyModifiers) {
	if period < time.Second {
		return
	}

	_, totalMythics, totalAscended := sc.calculateTotalGeneratedPets(period)

	ascDmgMultiplier := totalAscended / 6.0 / 1000.0 * sMods.CalculateShinyOdds()
	mythDmgMultiplier := totalMythics / 6.0 / 1000.0 * sMods.CalculateShinyOdds()

	fmt.Println(
		sc.printer.Sprintf(
			"ascended generated: %d ascended dmg multiplier gained: x%.5f (+%d dmg)",
			int(totalAscended),
			ascDmgMultiplier,
			int(1950*ascDmgMultiplier),
		),
	)
	fmt.Println(
		sc.printer.Sprintf(
			"mythic generated: %d mythic dmg multiplier gained: x%.5f (+%d dmg)",
			int(totalMythics),
			mythDmgMultiplier,
			int(2000*mythDmgMultiplier),
		),
	)
}

func (sc *Stones) CalculateMinedStones(period time.Duration) int {
	if period < time.Second {
		return 0
	}

	stonesPerStrike := 1.0
	for i := 2; i <= sc.eggLevel; i++ {
		stonesPerStrike += PerLevelEggModifier
	}

	regularStrikes := 0.0
	regularStrikes = sc.mineSpeed * period.Seconds()
	regularStrikes *= sc.firstStrike

	x2Strikes := regularStrikes * sc.x2Strike
	regularStrikes -= x2Strikes

	x3Strikes := x2Strikes * sc.x3Strike
	x2Strikes -= x3Strikes

	x4Strikes := x3Strikes * sc.x4Strike
	x3Strikes -= x4Strikes

	x5Strikes := x4Strikes * sc.x5Strike
	x4Strikes -= x5Strikes

	stones := regularStrikes * stonesPerStrike * sc.miningStoneBonus
	stones += x2Strikes * stonesPerStrike * sc.miningStoneBonus * 2
	stones += x3Strikes * stonesPerStrike * sc.miningStoneBonus * 3
	stones += x4Strikes * stonesPerStrike * sc.miningStoneBonus * 4
	stones += x5Strikes * stonesPerStrike * sc.miningStoneBonus * 5

	if sc.stonesOverclocked {
		return int(stones * 1.5)
	}

	return int(stones)
}

func (sc *Stones) calculateTotalGeneratedPets(period time.Duration) (total, mythics, ascended float64) {
	totalEggs := 0.0
	eggsPerSecond := MaxGenSpeed
	totalEggs = eggsPerSecond * period.Seconds()
	clonedEggs := totalEggs * sc.cloneLuck
	if sc.recursiveClone {
		clonedEggs += clonedEggs * sc.cloneLuck
	}
	totalEggs += clonedEggs

	directMythics := sc.eggLuck * totalEggs
	totalAscended := totalEggs - directMythics
	fusedMythics := totalAscended / 3
	totalMythics := directMythics + fusedMythics

	return totalEggs, totalMythics, totalAscended
}

func (sc *Stones) FindNextUpgrade(speedCost, cloneCost int) string {
	bestUpgrade := ""
	bestCostMargin := 0.0
	testDuration := time.Hour * 24

	for i := DoubleStrike; i <= QuintupleStrike; i++ {
		margin := sc.calculateStrikeImprovementMargin(i, testDuration)
		if margin > bestCostMargin {
			bestUpgrade = fmt.Sprintf("x%d strike", i)
			bestCostMargin = margin
		}
	}

	margin := sc.calculateSpeedImprovementMargin(speedCost, testDuration)
	if margin > bestCostMargin {
		bestUpgrade = "speed"
		bestCostMargin = margin
	}

	margin = sc.calculateCloneImprovementMargin(cloneCost, testDuration)
	if margin > bestCostMargin {
		bestUpgrade = "clone luck"
	}

	return bestUpgrade
}

func (sc *Stones) calculateStrikeImprovementMargin(strikeType int, period time.Duration) float64 {
	strikeLevel := sc.miningModifiers.StrikeUpgrades[strikeType] + 1
	strikeCosts := upgrade_data.GetStrikePrices()

	upgradeCalculator := sc.getBaselineComparator()
	baselineStones := upgradeCalculator.CalculateMinedStones(period)

	switch strikeType {
	case DoubleStrike:
		upgradeCalculator.x2Strike += upgrade_data.PerStepStrikeImprovement
	case TripleStrike:
		upgradeCalculator.x3Strike += upgrade_data.PerStepStrikeImprovement
	case QuadrupleStrike:
		upgradeCalculator.x4Strike += upgrade_data.PerStepStrikeImprovement
	case QuintupleStrike:
		upgradeCalculator.x5Strike += upgrade_data.PerStepStrikeImprovement
	}

	postUpgradeStones := upgradeCalculator.CalculateMinedStones(period)
	upgradeCost := strikeCosts[strikeLevel]
	return float64(postUpgradeStones-baselineStones) / float64(upgradeCost)
}

func (sc *Stones) calculateSpeedImprovementMargin(upgradeCost int, period time.Duration) float64 {
	upgradeCalculator := sc.getBaselineComparator()

	baselineStones := upgradeCalculator.CalculateMinedStones(period)
	upgradeCalculator.mineSpeed += PerLevelSpeedModifier
	postUpgradeStones := upgradeCalculator.CalculateMinedStones(period)

	return float64(postUpgradeStones-baselineStones) / float64(upgradeCost)
}

func (sc *Stones) calculateCloneImprovementMargin(upgradeCost int, period time.Duration) float64 {
	upgradeCalculator := sc.getBaselineComparator()

	baselineStones := upgradeCalculator.CalculateGeneratedStones(period)
	upgradeCalculator.cloneLuck += PerLevelCloneModifier
	postUpgradeStones := upgradeCalculator.CalculateMinedStones(period)

	return float64(postUpgradeStones-baselineStones) / float64(upgradeCost)
}

func (sc *Stones) getBaselineComparator() Stones {
	return Stones{
		x2Strike:         float64(sc.miningModifiers.StrikeUpgrades[DoubleStrike]) * upgrade_data.PerStepStrikeImprovement,
		x3Strike:         float64(sc.miningModifiers.StrikeUpgrades[TripleStrike]) * upgrade_data.PerStepStrikeImprovement,
		x4Strike:         float64(sc.miningModifiers.StrikeUpgrades[QuadrupleStrike]) * upgrade_data.PerStepStrikeImprovement,
		x5Strike:         float64(sc.miningModifiers.StrikeUpgrades[QuintupleStrike]) * upgrade_data.PerStepStrikeImprovement,
		eggLuck:          sc.eggLuck,
		cloneLuck:        sc.cloneLuck,
		eggLevel:         MythicEgg,
		mineSpeed:        sc.miningModifiers.MineSpeed,
		firstStrike:      1,
		miningStoneBonus: RubyPickMiningBonus,
	}
}
