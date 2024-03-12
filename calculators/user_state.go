package calculators

type MiningModifiers struct {
	MineSpeed             float64
	FirstStrike           float64
	StrikeUpgrades        strikeUpgrades
	StrikeOdds            strikeOdds
	GiantLuckLevel        int
	GiantOdds             float64
	MiningStoneMultiplier float64
}

func NewMiningModifiers(mineSpeed, firstStrike, giantOdds, stonesFromMining float64, strikeUpgrades strikeUpgrades, giantLuckLevel int, strikeOdds strikeOdds) MiningModifiers {
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
	}
}

type ShinyModifiers struct {
	shinyLuck float64
}

func NewShinyModifiers(shinyLuck float64) ShinyModifiers {
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
}

func NewEggGenerationModifiers(eggLuck, cloneLuck, calcifyChance float64, eggLevel int, hasRecursiveClone bool) EggGenerationModifiers {
	return EggGenerationModifiers{
		EggLuck:           eggLuck / 100,
		CloneLuck:         cloneLuck / 100,
		CalcifyChance:     1 + (calcifyChance / 100),
		EggLevel:          eggLevel,
		HasRecursiveClone: hasRecursiveClone,
	}
}
