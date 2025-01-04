package calculators

const AscendedShinyScore = 30
const MythicShinyScore = 40
const UpgradeComplete = 0
const X2OverclockMultiplier = 1.4
const X3OverclockMultiplier = 1.6
const X4OverclockMultiplier = 1.8
const X5OverclockMultiplier = 2.0
const GiantLuckOverclockMultiplier = 1.5

type MiningModifiers struct {
	MineSpeed             float64
	FirstStrike           float64
	StrikeUpgrades        strikeUpgrades
	StrikeOdds            strikeOdds
	GiantLuckLevel        int
	GiantOdds             float64
	MiningStoneMultiplier float64
	x2Overclock           bool
	x3Overclock           bool
	x4Overclock           bool
	x5Overclock           bool
}

func NewMiningModifiers(mineSpeed, firstStrike, giantOdds, stonesFromMining float64, strikeUpgrades strikeUpgrades, giantLuckLevel int, strikeOdds strikeOdds, x2Overclock, x3Overclock, x4Overclock, x5Overclock bool) MiningModifiers {
	strikeOdds[DoubleStrike] = strikeOdds[DoubleStrike] / 100
	strikeOdds[TripleStrike] = strikeOdds[TripleStrike] / 100
	strikeOdds[QuadrupleStrike] = strikeOdds[QuadrupleStrike] / 100
	strikeOdds[QuintupleStrike] = strikeOdds[QuintupleStrike] / 100

	return MiningModifiers{
		MineSpeed:             mineSpeed,
		FirstStrike:           firstStrike / 100,
		StrikeUpgrades:        strikeUpgrades,
		GiantLuckLevel:        giantLuckLevel,
		GiantOdds:             giantOdds / 100,
		StrikeOdds:            strikeOdds,
		MiningStoneMultiplier: stonesFromMining / 100,
		x2Overclock:           x2Overclock,
		x3Overclock:           x3Overclock,
		x4Overclock:           x4Overclock,
		x5Overclock:           x5Overclock,
	}
}

type ShinyModifiers struct {
	shinyLuck float64
}

func NewShinyModifiers(shinyLuck float64) ShinyModifiers {
	if shinyLuck > 100 {
		shinyLuck = 100
	}
	return ShinyModifiers{shinyLuck: shinyLuck / 100}
}

func (sm ShinyModifiers) CalculateShinyOdds() float64 {
	return sm.shinyLuck
}

type EggGenerationModifiers struct {
	EggLuck           float64
	CloneLuck         float64
	CalcifyChance     float64
	EggLevel          int
	HasRecursiveClone bool
	IsUsingCrank      bool
}

func NewEggGenerationModifiers(eggLuck, cloneLuck, calcifyChance float64, eggLevel int, hasRecursiveClone, isUsingCrank bool) EggGenerationModifiers {
	return EggGenerationModifiers{
		EggLuck:           eggLuck / 100,
		CloneLuck:         cloneLuck / 100,
		CalcifyChance:     1 + (calcifyChance / 100),
		EggLevel:          eggLevel,
		HasRecursiveClone: hasRecursiveClone,
		IsUsingCrank:      isUsingCrank,
	}
}

type GiantModifiers struct {
	t7GiantLuck      float64
	t8GiantLuck      float64
	achievement      float64
	rune             float64
	luckOverclocked  bool
	shinyOverclocked bool
}

func NewGiantModifiers(t7GiantLuck, t8GiantLuck, achievement, rune float64, luckOverclocked, shinyOverclocked bool) GiantModifiers {
	return GiantModifiers{
		t7GiantLuck:      t7GiantLuck,
		t8GiantLuck:      t8GiantLuck,
		achievement:      achievement,
		rune:             rune,
		luckOverclocked:  luckOverclocked,
		shinyOverclocked: shinyOverclocked,
	}
}

type AscensionModifiers struct {
	genSpeedBonus       float64
	giantLuckMultiplier float64
}

func NewAscensionModifiers(wingboltLevel int, giantLuckModifier float64) AscensionModifiers {
	mods := AscensionModifiers{}
	switch wingboltLevel {
	case 1:
		mods.genSpeedBonus = 0.5
	case 2:
		mods.genSpeedBonus = 1
	case 3:
		mods.genSpeedBonus = 1.5
	}

	if giantLuckModifier < 1 {
		giantLuckModifier = 1
	}
	mods.giantLuckMultiplier = giantLuckModifier

	return mods
}
