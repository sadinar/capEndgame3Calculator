package character_config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

func TestGetEggIndex(t *testing.T) {
	c := characterConfig{}
	c.GenerationMods.MinedEgg = "common"
	assert.Equal(t, 1, c.getEggIndex())
	c.GenerationMods.MinedEgg = "COMMON"
	assert.Equal(t, 1, c.getEggIndex())

	c.GenerationMods.MinedEgg = "uncommon"
	assert.Equal(t, 2, c.getEggIndex())

	c.GenerationMods.MinedEgg = "rare"
	assert.Equal(t, 3, c.getEggIndex())
	c.GenerationMods.MinedEgg = "Rare"
	assert.Equal(t, 3, c.getEggIndex())

	c.GenerationMods.MinedEgg = "epic"
	assert.Equal(t, 4, c.getEggIndex())
	c.GenerationMods.MinedEgg = "EpIc"
	assert.Equal(t, 4, c.getEggIndex())

	c.GenerationMods.MinedEgg = "legendary"
	assert.Equal(t, 5, c.getEggIndex())

	c.GenerationMods.MinedEgg = "prodigious"
	assert.Equal(t, 6, c.getEggIndex())

	c.GenerationMods.MinedEgg = "ascended"
	assert.Equal(t, 7, c.getEggIndex())

	c.GenerationMods.MinedEgg = "mythic"
	assert.Equal(t, 8, c.getEggIndex())
}

func TestConfigureCalculators(t *testing.T) {
	shinyMods, giantCalc, stoneCalc, speedCost, cloneCost := ConfigureCalculators("./pre_ascend_sadinar.json")

	assert.Equal(t, "speed", giantCalc.GetNextUpgrade(speedCost))
	assert.Equal(t, "speed", stoneCalc.FindNextUpgrade(speedCost, cloneCost))

	median, shinies, medianOdds := giantCalc.PrintProbabilityMedian(time.Hour*24, shinyMods)
	assert.Equal(t, 1327, median)
	assert.Equal(t, 1327, shinies)
	assert.Greater(t, 0.5015952014327, medianOdds)
	assert.Less(t, 0.5015952014326, medianOdds)

	gennedStones, minedStones := stoneCalc.CalculateStonesProduced(time.Hour * 24)
	assert.Equal(t, 365836, gennedStones)
	assert.Equal(t, 4153276, minedStones)

	expectedDmgOutput := "ascended generated: 467,025 (14,010,750 shiny score): ascended dmg multiplier " +
		"gained: x77.83750 (+151,783 dmg)\nmythic generated: 155,675 (6,227,000 shiny score): mythic dmg " +
		"multiplier gained: x25.94583 (+51,891 dmg)"
	assert.Equal(t, expectedDmgOutput, stoneCalc.PrintDamageChange(time.Hour*24, shinyMods))
}

func TestAscensionPets(t *testing.T) {
	character := parseCharacterFile("./AllAscensionPets.json")
	assert.Equal(t, 2, character.AscensionMods.WingboltLevel)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}
