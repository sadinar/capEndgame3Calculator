package calculators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAscensionModifiers(t *testing.T) {
	am := NewAscensionModifiers(0, 0)
	assert.Equal(t, 1.0, am.giantLuckMultiplier)
}
