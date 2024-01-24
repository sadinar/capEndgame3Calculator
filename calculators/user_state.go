package calculators

const StoneOverclockIndex = 1337

type OverclockConfig map[int]bool

func NewOverclockConfig(stones, x2, x3, x4, x5, giant bool) OverclockConfig {
	return OverclockConfig{
		StoneOverclockIndex: stones,
		DoubleStrike:        x2,
		TripleStrike:        x3,
		QuadrupleStrike:     x4,
		QuintupleStrike:     x5,
		GiantLuck:           giant,
	}
}

type UserModifiers struct {
	GiantLuckAchievementModifier float64
	GiantLuckRuneModifier        float64
	MineSpeed                    float64
	FirstStrike                  float64
	StrikeUpgrades               strikeUpgrades
	GiantLuckLevel               int
}

func NewUserModifiers(giantLuckAchievementModifier, giantLuckRuneModifier, mineSpeed, firstStrike float64, strikeUpgrades strikeUpgrades, giantLuckLevel int) UserModifiers {
	return UserModifiers{
		GiantLuckAchievementModifier: giantLuckAchievementModifier,
		GiantLuckRuneModifier:        giantLuckRuneModifier,
		MineSpeed:                    mineSpeed,
		FirstStrike:                  firstStrike,
		StrikeUpgrades:               strikeUpgrades,
		GiantLuckLevel:               giantLuckLevel,
	}
}
