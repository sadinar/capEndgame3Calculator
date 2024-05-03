package calculators

import (
	"capEndgame3Calculator/upgrade_data"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"time"
)

const MaxGenSpeed = 5.0
const PerLevelEggModifier = 0.5
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
const MaxClones = 9
const UniqueAscendedPets = 6
const UniqueMythicPets = 6
const BaseShinyDivisor = 1000
const TitaniumAscendedBaseDamage = 1950
const TitaniumMythicBaseDamage = 2000

type Stones struct {
	generationModifiers EggGenerationModifiers
	miningModifiers     MiningModifiers
	printer             *message.Printer
}

func NewStonesCalculator(mm MiningModifiers, egm EggGenerationModifiers) Stones {
	sc := Stones{
		generationModifiers: egm,
		miningModifiers:     mm,
		printer:             message.NewPrinter(language.English),
	}

	return sc
}

func (sc *Stones) CalculateStonesProduced(period time.Duration) (generatedStones, minedStones int) {
	return sc.CalculateGeneratedStones(period), sc.CalculateMinedStones(period)
}

func (sc *Stones) CalculateGeneratedStones(period time.Duration) int {
	if period < time.Second {
		return 0
	}

	_, totalMythics, _ := sc.calculateTotalGeneratedPets(period)

	return int(totalMythics * sc.generationModifiers.CalcifyChance)
}

func (sc *Stones) PrintDamageChange(period time.Duration, sMods ShinyModifiers) string {
	if period < time.Second {
		return "no change"
	}

	_, totalMythics, totalAscended := sc.calculateTotalGeneratedPets(period)

	ascDmgMultiplier := totalAscended / UniqueAscendedPets / BaseShinyDivisor * sMods.CalculateShinyOdds()
	mythDmgMultiplier := totalMythics / UniqueMythicPets / BaseShinyDivisor * sMods.CalculateShinyOdds()

	ascendedOutput := sc.printer.Sprintf(
		"ascended generated: %d (%d shiny score): ascended dmg multiplier gained: x%.5f (+%d dmg)",
		int(totalAscended),
		int(sMods.CalculateShinyOdds()*totalAscended*AscendedShinyScore),
		ascDmgMultiplier,
		int(TitaniumAscendedBaseDamage*ascDmgMultiplier),
	)
	mythicOutput := sc.printer.Sprintf(
		"mythic generated: %d (%d shiny score): mythic dmg multiplier gained: x%.5f (+%d dmg)",
		int(totalMythics),
		int(sMods.CalculateShinyOdds()*totalMythics*MythicShinyScore),
		mythDmgMultiplier,
		int(TitaniumMythicBaseDamage*mythDmgMultiplier),
	)

	fmt.Println(ascendedOutput)
	fmt.Println(mythicOutput)

	return ascendedOutput + "\n" + mythicOutput
}

func (sc *Stones) CalculateMinedStones(period time.Duration) int {
	if period < time.Second {
		return 0
	}

	stonesPerStrike := 1.0
	for i := 1; i < sc.generationModifiers.EggLevel; i++ {
		stonesPerStrike += PerLevelEggModifier
	}
	stonesPerStrike *= sc.miningModifiers.MiningStoneMultiplier
	if stonesPerStrike < 1 {
		stonesPerStrike = 1
	}

	regularStrikes := 0.0
	regularStrikes = sc.miningModifiers.MineSpeed * period.Seconds()
	regularStrikes *= sc.miningModifiers.FirstStrike

	x2Strikes := regularStrikes * sc.miningModifiers.StrikeOdds[DoubleStrike]
	x3Strikes := regularStrikes * sc.miningModifiers.StrikeOdds[TripleStrike]
	x4Strikes := regularStrikes * sc.miningModifiers.StrikeOdds[QuadrupleStrike]
	x5Strikes := regularStrikes * sc.miningModifiers.StrikeOdds[QuintupleStrike]

	regularStrikes -= x2Strikes
	x2Strikes -= x3Strikes
	x3Strikes -= x4Strikes
	x4Strikes -= x5Strikes

	stones := regularStrikes * stonesPerStrike
	stones += x2Strikes * stonesPerStrike * 2
	stones += x3Strikes * stonesPerStrike * 3
	stones += x4Strikes * stonesPerStrike * 4
	stones += x5Strikes * stonesPerStrike * 5

	return int(stones)
}

func (sc *Stones) calculateTotalGeneratedPets(period time.Duration) (total, mythics, ascended float64) {
	totalEggs := 0.0
	eggsPerSecond := MaxGenSpeed
	totalEggs = eggsPerSecond * period.Seconds()
	clonedEggs := totalEggs * sc.generationModifiers.CloneLuck
	if sc.generationModifiers.HasRecursiveClone {
		clonedEggs += float64(sc.recursivelyClone(1, int(clonedEggs)))
	}
	totalEggs += clonedEggs

	directMythics := sc.generationModifiers.EggLuck * totalEggs
	totalAscended := totalEggs - directMythics
	fusedMythics := totalAscended / 3
	totalMythics := directMythics + fusedMythics

	return totalEggs, totalMythics, totalAscended
}

func (sc *Stones) recursivelyClone(cloneDepth int, startingClones int) int {
	if cloneDepth > MaxClones {
		return 0
	}

	newClones := int(float64(startingClones) * sc.generationModifiers.CloneLuck)
	if newClones == 0 {
		return 0
	}
	return newClones + sc.recursivelyClone(cloneDepth+1, newClones)
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

	if len(bestUpgrade) == 0 {
		return "n/a"
	}

	return bestUpgrade
}

func (sc *Stones) calculateStrikeImprovementMargin(strikeType int, period time.Duration) float64 {
	strikeLevel := sc.miningModifiers.StrikeUpgrades[strikeType] + 1
	strikeCosts := upgrade_data.GetStrikePrices()

	if strikeLevel > len(strikeCosts) {
		return 0
	}

	upgradeCalculator := sc.copyComparator()
	baselineStones := upgradeCalculator.CalculateMinedStones(period)

	switch strikeType {
	case DoubleStrike:
		upgradeCalculator.miningModifiers.StrikeOdds[DoubleStrike] += upgrade_data.PerStepStrikeImprovement * 1.4
	case TripleStrike:
		originalTripleOdds := upgradeCalculator.miningModifiers.StrikeOdds[TripleStrike] / upgradeCalculator.miningModifiers.StrikeOdds[DoubleStrike]
		increasedTripleOdds := originalTripleOdds + upgrade_data.PerStepStrikeImprovement*1.6
		upgradeCalculator.miningModifiers.StrikeOdds[TripleStrike] = increasedTripleOdds * upgradeCalculator.miningModifiers.StrikeOdds[DoubleStrike]
	case QuadrupleStrike:
		originalQuadOdds := upgradeCalculator.miningModifiers.StrikeOdds[QuadrupleStrike] / upgradeCalculator.miningModifiers.StrikeOdds[TripleStrike]
		increasedQuadOdds := originalQuadOdds + upgrade_data.PerStepStrikeImprovement*1.8
		upgradeCalculator.miningModifiers.StrikeOdds[QuadrupleStrike] = increasedQuadOdds * upgradeCalculator.miningModifiers.StrikeOdds[TripleStrike]
	case QuintupleStrike:
		originalPentaOdds := upgradeCalculator.miningModifiers.StrikeOdds[QuintupleStrike] / upgradeCalculator.miningModifiers.StrikeOdds[QuadrupleStrike]
		increasedPentaOdds := originalPentaOdds + upgrade_data.PerStepStrikeImprovement*2
		upgradeCalculator.miningModifiers.StrikeOdds[QuintupleStrike] = increasedPentaOdds * upgradeCalculator.miningModifiers.StrikeOdds[QuadrupleStrike]
	}

	postUpgradeStones := upgradeCalculator.CalculateMinedStones(period)
	upgradeCost := strikeCosts[strikeLevel]
	return float64(postUpgradeStones-baselineStones) / float64(upgradeCost)
}

func (sc *Stones) calculateSpeedImprovementMargin(upgradeCost int, period time.Duration) float64 {
	if upgradeCost == UpgradeComplete {
		return 0
	}

	upgradeCalculator := sc.copyComparator()

	baselineStones := upgradeCalculator.CalculateMinedStones(period)
	upgradeCalculator.miningModifiers.MineSpeed += PerLevelSpeedModifier
	postUpgradeStones := upgradeCalculator.CalculateMinedStones(period)

	return float64(postUpgradeStones-baselineStones) / float64(upgradeCost)
}

func (sc *Stones) calculateCloneImprovementMargin(upgradeCost int, period time.Duration) float64 {
	if upgradeCost == UpgradeComplete {
		return 0
	}

	upgradeCalculator := sc.copyComparator()

	baselineStones := upgradeCalculator.CalculateGeneratedStones(period)
	upgradeCalculator.generationModifiers.CloneLuck += PerLevelCloneModifier
	postUpgradeStones := upgradeCalculator.CalculateGeneratedStones(period)

	return float64(postUpgradeStones-baselineStones) / float64(upgradeCost)
}

func (sc *Stones) copyComparator() Stones {
	return Stones{
		miningModifiers: NewMiningModifiers(
			sc.miningModifiers.MineSpeed,
			sc.miningModifiers.FirstStrike*100,
			0,
			sc.miningModifiers.MiningStoneMultiplier*100,
			nil,
			0,
			strikeOdds{
				DoubleStrike:    sc.miningModifiers.StrikeOdds[DoubleStrike] * 100,
				TripleStrike:    sc.miningModifiers.StrikeOdds[TripleStrike] * 100,
				QuadrupleStrike: sc.miningModifiers.StrikeOdds[QuadrupleStrike] * 100,
				QuintupleStrike: sc.miningModifiers.StrikeOdds[QuintupleStrike] * 100,
			},
			sc.miningModifiers.x2Overclock,
			sc.miningModifiers.x3Overclock,
			sc.miningModifiers.x4Overclock,
			sc.miningModifiers.x5Overclock,
		),
		generationModifiers: NewEggGenerationModifiers(
			sc.generationModifiers.EggLuck,
			sc.generationModifiers.CloneLuck,
			sc.generationModifiers.CalcifyChance,
			MythicEgg,
			sc.generationModifiers.HasRecursiveClone,
		),
	}
}
