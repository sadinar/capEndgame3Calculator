package calculators

import (
	"capEndgame3Calculator/upgrade_data"
	"fmt"
	"time"
)

const MaxGenSpeed = 5.0
const MaxCalcify = 2
const StonePick = 1.0
const EmeraldPick = 1.1
const SapphirePick = 1.2
const AmethystPick = 1.3
const TopazPick = 1.4
const QuartzPick = 1.5
const DiamondPick = 1.75
const RubyPick = 2.0
const PerLevelEggModifier = 0.40
const StoneOverclockModifier = 1.5
const CommonEgg = 1
const UncommonEgg = 2
const RareEgg = 3
const EpicEgg = 4
const LegendaryEgg = 5
const ProdigiousEgg = 6
const AscendedEgg = 7
const MythicEgg = 9

type Stones struct {
	firstStrike       float64
	x2Strike          float64
	x3Strike          float64
	x4Strike          float64
	x5Strike          float64
	mineSpeed         float64
	pickModifier      float64
	eggLuck           float64
	eggLevel          int
	stonesOverclocked bool
}

func NewStonesCalculator(um UserModifiers, pickModifier, eggLuck float64, eggLevel int, ocConfig OverclockConfig) Stones {
	sc := Stones{
		firstStrike:       um.FirstStrike,
		x2Strike:          float64(um.StrikeUpgrades[DoubleStrike]) * upgrade_data.PerStepStrikeImprovement,
		x3Strike:          float64(um.StrikeUpgrades[TripleStrike]) * upgrade_data.PerStepStrikeImprovement,
		x4Strike:          float64(um.StrikeUpgrades[QuadrupleStrike]) * upgrade_data.PerStepStrikeImprovement,
		x5Strike:          float64(um.StrikeUpgrades[QuintupleStrike]) * upgrade_data.PerStepStrikeImprovement,
		mineSpeed:         um.MineSpeed,
		pickModifier:      pickModifier,
		eggLuck:           eggLuck,
		eggLevel:          eggLevel,
		stonesOverclocked: ocConfig[StoneOverclockIndex],
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

func (sc *Stones) CalculateCombinedStones(period time.Duration) int {
	generatedStones := sc.CalculateGeneratedStones(period)
	minedStones := sc.CalculateMinedStones(period)

	return generatedStones + minedStones
}

func (sc *Stones) CalculateGeneratedStones(period time.Duration) int {
	if period < time.Second {
		return 0
	}

	totalEggs := 0.0
	eggsPerSecond := MaxGenSpeed
	totalEggs = eggsPerSecond * period.Seconds()

	shinyLuck := 1.0 / 1000 * 1.1 * 1.1 * 1.2 * 4.2 * 5.2
	directMythics := sc.eggLuck * totalEggs
	totalAscended := totalEggs - directMythics
	ascDmgMult := totalAscended / 6.0 / 1000.0 * shinyLuck
	fusedMythics := totalAscended / 3
	totalMythics := directMythics + fusedMythics
	mythDmgMult := totalMythics / 6.0 / 1000.0 * shinyLuck

	fmt.Println(fmt.Sprintf("ascended generated: %d ascended dmg multiplier gained: x%.5f (+%d dmg)", int(totalAscended), ascDmgMult, int(1950*ascDmgMult)))
	fmt.Println(fmt.Sprintf("mythic generated: %d mythic dmg multiplier gained: x%.5f (+%d dmg)", int(totalMythics), mythDmgMult, int(2000*mythDmgMult)))

	return int(totalMythics) * MaxCalcify
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

	stones := regularStrikes * stonesPerStrike * sc.pickModifier
	stones += x2Strikes * stonesPerStrike * sc.pickModifier * 2
	stones += x3Strikes * stonesPerStrike * sc.pickModifier * 3
	stones += x4Strikes * stonesPerStrike * sc.pickModifier * 4
	stones += x5Strikes * stonesPerStrike * sc.pickModifier * 5

	if sc.stonesOverclocked {
		return int(stones * 1.5)
	}

	return int(stones)
}
