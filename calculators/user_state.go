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
	GiantLuckLevel               int
}

func NewMiningModifiers(giantLuckAchievementModifier, giantLuckRuneModifier, mineSpeed, firstStrike float64, strikeUpgrades strikeUpgrades, giantLuckLevel int) MiningModifiers {
	return MiningModifiers{
		GiantLuckAchievementModifier: giantLuckAchievementModifier,
		GiantLuckRuneModifier:        giantLuckRuneModifier,
		MineSpeed:                    mineSpeed,
		FirstStrike:                  firstStrike,
		StrikeUpgrades:               strikeUpgrades,
		GiantLuckLevel:               giantLuckLevel,
	}
}

type ShinyModifiers struct {
	shinyAchievement      float64
	shinyGiantAchievement float64
	rune                  float64
	wall                  float64
	petScoreModifier      float64
	prismaticShine        float64
}

func NewShinyModifiers(shinyAchievement, shinyGiantAchievement, rune, wall, prismaticShine float64, petScore int) ShinyModifiers {
	petScoreModifier := float64(petScore)/float64(10000000) + 1
	return ShinyModifiers{
		shinyAchievement:      shinyAchievement,
		shinyGiantAchievement: shinyGiantAchievement,
		rune:                  rune,
		wall:                  wall,
		petScoreModifier:      petScoreModifier,
		prismaticShine:        prismaticShine,
	}
}

func (sm ShinyModifiers) CalculateShinyOdds() float64 {
	return sm.shinyGiantAchievement * sm.shinyGiantAchievement * sm.rune * sm.wall * sm.petScoreModifier * sm.prismaticShine / 1000
}
