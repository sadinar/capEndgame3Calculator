package calculators

import (
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
const PerLevelEggModifier = 0.5
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

func NewStonesCalculator(firstStrike, x2Strike, x3Strike, x4Strike, x5Strike, mineSpeed, pickModifier, eggLuck float64, eggLevel int, stonesOverclocked, x2Overclocked, x3Overclocked, x4Overclocked, x5Overclocked bool) Stones {
	sc := Stones{
		firstStrike:       firstStrike,
		x2Strike:          x2Strike,
		x3Strike:          x3Strike,
		x4Strike:          x4Strike,
		x5Strike:          x5Strike,
		mineSpeed:         mineSpeed,
		pickModifier:      pickModifier,
		eggLuck:           eggLuck,
		eggLevel:          eggLevel,
		stonesOverclocked: stonesOverclocked,
	}

	if x5Overclocked {
		sc.x5Strike *= 2
	}

	if x4Overclocked {
		sc.x4Strike *= 1.8
	}

	if x3Overclocked {
		sc.x3Strike *= 1.6
	}

	if x2Overclocked {
		sc.x2Strike *= 1.4
	}

	return sc
}

func (sc *Stones) CalculateStones(period time.Duration) int {
	generatedStones := sc.calculateGeneratedStones(period)
	minedStones := sc.calculateMinedStones(period)

	return generatedStones + minedStones
}

func (sc *Stones) calculateGeneratedStones(period time.Duration) int {
	if period < time.Second {
		return 0
	}

	totalEggs := 0.0
	eggsPerSecond := MaxGenSpeed
	if period < time.Minute {
		totalEggs = eggsPerSecond * period.Seconds()
	} else if period < time.Hour {
		totalEggs = eggsPerSecond * 60 * period.Minutes()
	} else {
		totalEggs = eggsPerSecond * 60 * 60 * period.Hours()
	}

	directMythics := sc.eggLuck * totalEggs
	fusedMythics := (totalEggs - directMythics) / 3

	return int(directMythics) + int(fusedMythics)*MaxCalcify
}

func (sc *Stones) calculateMinedStones(period time.Duration) int {
	if period < time.Second {
		return 0
	}

	stonesPerStrike := 1.0
	for i := 2; i <= sc.eggLevel; i++ {
		stonesPerStrike += PerLevelEggModifier
	}

	stoneModifier := sc.pickModifier
	if sc.stonesOverclocked {
		stoneModifier *= 1.5
	}

	regularStrikes := 0.0
	if period < time.Minute {
		regularStrikes = sc.mineSpeed * period.Seconds()
	} else if period < time.Hour {
		regularStrikes = sc.mineSpeed * 60 * period.Minutes()
	} else {
		regularStrikes = sc.mineSpeed * 60 * 60 * period.Hours()
	}
	regularStrikes *= sc.firstStrike

	x2Strikes := regularStrikes * sc.x2Strike
	regularStrikes -= x2Strikes

	x3Strikes := x2Strikes * sc.x3Strike
	x2Strikes -= x3Strikes

	x4Strikes := x3Strikes * sc.x4Strike
	x3Strikes -= x4Strikes

	x5Strikes := x4Strikes * sc.x5Strike
	x4Strikes -= x5Strikes

	stones := regularStrikes * stonesPerStrike * stoneModifier
	stones += x2Strikes * stonesPerStrike * stoneModifier * 2
	stones += x3Strikes * stonesPerStrike * stoneModifier * 3
	stones += x4Strikes * stonesPerStrike * stoneModifier * 4
	stones += x5Strikes * stonesPerStrike * stoneModifier * 5

	return int(stones)
}
