package calculators

type MiningModifiers struct {
	MineSpeed      float64
	FirstStrike    float64
	StrikeUpgrades strikeUpgrades
	StrikeOdds     strikeOdds
	GiantLuckLevel int
	GiantOdds      float64
}

func NewMiningModifiers(mineSpeed, firstStrike, giantOdds float64, strikeUpgrades strikeUpgrades, giantLuckLevel int, strikeOdds strikeOdds) MiningModifiers {
	strikeOdds[DoubleStrike] = strikeOdds[DoubleStrike] / 100
	strikeOdds[TripleStrike] = strikeOdds[TripleStrike] / 100
	strikeOdds[QuadrupleStrike] = strikeOdds[QuadrupleStrike] / 100
	strikeOdds[QuintupleStrike] = strikeOdds[QuintupleStrike] / 100

	return MiningModifiers{
		MineSpeed:      mineSpeed,
		FirstStrike:    firstStrike,
		StrikeUpgrades: strikeUpgrades,
		GiantLuckLevel: giantLuckLevel,
		GiantOdds:      giantOdds / 100,
		StrikeOdds:     strikeOdds,
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
