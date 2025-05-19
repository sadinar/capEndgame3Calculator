package calculators

type PetScore struct {
	bonusPetScoreOdds float64
}

func NewBonusPetScoreCalculator(bonusPetScoreOdds float64) PetScore {
	return PetScore{bonusPetScoreOdds: bonusPetScoreOdds}
}

func (p PetScore) BonusPetScore(generatedMythics float64) int {
	return int(generatedMythics * p.bonusPetScoreOdds)
}
