package calculators

import (
	"math/big"
)

func Factorial(factor *big.Int) *big.Int {
	if factor.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}
	oneLess := big.NewInt(0)
	oneLess = oneLess.Sub(factor, big.NewInt(1))

	return factor.Mul(factor, Factorial(oneLess))
}

func BinomialProbability(trials uint64, successes uint64, pSuccess float64) float64 {
	combinations := totalCombinations(trials, successes)
	fCombinations := big.NewFloat(0).SetInt(combinations)

	pFail := float64(1) - pSuccess
	s := bigFloatPow(pSuccess, successes)
	f := bigFloatPow(pFail, trials-successes)
	t := fCombinations.Mul(fCombinations, s)
	t = t.Mul(t, f)
	probability, _ := t.Float64()

	return probability
}

func FindReasonableSuccessCeiling(trials uint64, pSuccess float64) (uint64, float64) {
	totalProbability := float64(0)
	successCount := uint64(0)
	for {
		if totalProbability >= 0.95 {
			break
		}
		p := BinomialProbability(trials, successCount, pSuccess)
		totalProbability += p
		successCount++
	}

	return successCount - 1, totalProbability
}

func totalCombinations(trials uint64, successes uint64) *big.Int {
	if successes == 0 {
		return big.NewInt(1)
	}

	numerator := big.NewInt(int64(trials))
	for i := trials - 1; i > trials-successes; i-- {
		numerator.Mul(numerator, big.NewInt(int64(i)))
	}

	return numerator.Div(numerator, Factorial(big.NewInt(int64(successes))))
}

func bigFloatPow(base float64, exponent uint64) *big.Float {
	if exponent == 0 {
		return big.NewFloat(1)
	}

	answer := big.NewFloat(base)
	for {
		if exponent == 1 {
			return answer
		}
		answer = answer.Mul(answer, big.NewFloat(base))
		exponent--
	}
}
