package calculators

const StoneOverclockIndex = 1337
const ShinyOverclockIndex = 101001

type OverclockConfig map[int]bool

func NewOverclockConfig(stones, x2, x3, x4, x5, giant, shiny bool) OverclockConfig {
	return OverclockConfig{
		StoneOverclockIndex: stones,
		DoubleStrike:        x2,
		TripleStrike:        x3,
		QuadrupleStrike:     x4,
		QuintupleStrike:     x5,
		GiantLuck:           giant,
		ShinyOverclockIndex: shiny,
	}
}

type MiningModifiers struct {
	GiantLuckAchievementModifier float64
	GiantLuckRuneModifier        float64
	MineSpeed                    float64
	FirstStrike                  float64
	StrikeUpgrades               strikeUpgrades
	StrikeOdds                   strikeOdds
	GiantLuckLevel               int
	GiantOdds                    float64
}

func NewMiningModifiers(giantLuckAchievementModifier, giantLuckRuneModifier, mineSpeed, firstStrike, giantOdds float64, strikeUpgrades strikeUpgrades, giantLuckLevel int, strikeOdds strikeOdds) MiningModifiers {
	strikeOdds[DoubleStrike] = strikeOdds[DoubleStrike] / 100
	strikeOdds[TripleStrike] = strikeOdds[TripleStrike] / 100
	strikeOdds[QuadrupleStrike] = strikeOdds[QuadrupleStrike] / 100
	strikeOdds[QuintupleStrike] = strikeOdds[QuintupleStrike] / 100

	return MiningModifiers{
		GiantLuckAchievementModifier: giantLuckAchievementModifier,
		GiantLuckRuneModifier:        giantLuckRuneModifier,
		MineSpeed:                    mineSpeed,
		FirstStrike:                  firstStrike,
		StrikeUpgrades:               strikeUpgrades,
		GiantLuckLevel:               giantLuckLevel,
		GiantOdds:                    giantOdds / 100,
		StrikeOdds:                   strikeOdds,
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
