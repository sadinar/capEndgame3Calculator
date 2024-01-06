package calculators

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestBigFloatPow(t *testing.T) {
	assert.Equal(t, big.NewFloat(1), bigFloatPow(7, 0))
	assert.Equal(t, big.NewFloat(6), bigFloatPow(6, 1))
	assert.Equal(t, big.NewFloat(25), bigFloatPow(5, 2))
	assert.Equal(t, big.NewFloat(27), bigFloatPow(3, 3))
}

func TestBinomialProbability(t *testing.T) {
	assert.Equal(t, 0.3125, BinomialProbability(5, 2, 0.5))
	assert.Equal(t, 0.15625, BinomialProbability(5, 4, 0.5))
	assert.Equal(t, 0.03125, BinomialProbability(5, 5, 0.5))

	rounded := fmt.Sprintf("%.4f", BinomialProbability(10, 7, 0.8))
	assert.Equal(t, "0.2013", rounded)

	rounded = fmt.Sprintf("%.4f", BinomialProbability(5, 3, .75))
	assert.Equal(t, "0.2637", rounded)
}

func TestFactorial(t *testing.T) {
	assert.Equal(t, int64(3*2), Factorial(big.NewInt(3)).Int64())
	assert.Equal(t, int64(9*8*7*6*5*4*3*2), Factorial(big.NewInt(9)).Int64())
}
