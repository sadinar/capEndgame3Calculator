package calculators

import (
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
