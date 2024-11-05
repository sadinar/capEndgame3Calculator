package character_config

import (
	"capEndgame3Calculator/calculators"
	"encoding/json"
	"os"
	"strings"
)

type characterConfig struct {
	MiningMods struct {
		Speed             float64 `json:"speed"`
		FirstStrike       float64 `json:"firstStrike"`
		GiantOdds         float64 `json:"giantOdds"`
		StonesFromMining  float64 `json:"stonesFromMining"`
		CartUpgradeLevels struct {
			X2Strike  int `json:"x2Strike"`
			X3Strike  int `json:"x3Strike"`
			X4Strike  int `json:"x4Strike"`
			X5Strike  int `json:"x5Strike"`
			GiantLuck int `json:"giantLuck"`
		} `json:"cartUpgradeLevels"`
		GiantLuckLevel int `json:"giantLuckLevel"`
		StrikeOdds     struct {
			X2Strike float64 `json:"x2Strike"`
			X3Strike float64 `json:"x3Strike"`
			X4Strike float64 `json:"x4Strike"`
			X5Strike float64 `json:"x5Strike"`
		} `json:"strikeOdds"`
		Overclocks struct {
			X2 bool `json:"x2"`
			X3 bool `json:"x3"`
			X4 bool `json:"x4"`
			X5 bool `json:"x5"`
		} `json:"overclocks"`
	} `json:"miningMods"`
	GenerationMods struct {
		EggLuck           float64 `json:"eggLuck"`
		CloneLuck         float64 `json:"cloneLuck"`
		CalcifyChance     float64 `json:"calcifyChance"`
		MinedEgg          string  `json:"minedEgg"`
		HasRecursiveClone bool    `json:"hasRecursiveClone"`
	} `json:"generationMods"`
	ShinyLuck     float64 `json:"shinyLuck"`
	GiantLuckMods struct {
		LabTier7         float64 `json:"labTier7"`
		LabTier8         float64 `json:"labTier8"`
		Achievement      float64 `json:"achievement"`
		Runes            float64 `json:"runes"`
		Overclocked      bool    `json:"overclocked"`
		ShinyOverclocked bool    `json:"shinyOverclocked"`
	} `json:"giantLuckMods"`
	NextMineSpeedCost int `json:"nextMineSpeedCost"`
	NextCloneLuckCost int `json:"nextCloneLuckCost"`
	AscensionMods     struct {
		TrunkyLevel   int `json:"trunkyLevel"`
		HoppityLevel  int `json:"hoppityLevel"`
		GrimLevel     int `json:"grimLevel"`
		WingboltLevel int `json:"wingboltLevel"`
		NovaLevel     int `json:"novaLevel"`
		RadiLevel     int `json:"radiLevel"`
		BattackLevel  int `json:"battackLevel"`
		FlutterLevel  int `json:"flutterLevel"`
	} `json:"ascensionMods"`
}

func (c characterConfig) getEggIndex() int {
	switch strings.ToLower(c.GenerationMods.MinedEgg) {
	case "common":
		return calculators.CommonEgg
	case "uncommon":
		return calculators.UncommonEgg
	case "rare":
		return calculators.RareEgg
	case "epic":
		return calculators.EpicEgg
	case "legendary":
		return calculators.LegendaryEgg
	case "prodigious":
		return calculators.ProdigiousEgg
	case "ascended":
		return calculators.AscendedEgg
	case "mythic":
		return calculators.MythicEgg
	default:
		panic("unknown egg type: " + c.GenerationMods.MinedEgg)
	}
}

func parseCharacterFile(fileLoc string) characterConfig {
	fileData, err := os.ReadFile(fileLoc)
	if err != nil {
		panic(err)
	}

	character := characterConfig{}
	err = json.Unmarshal(fileData, &character)
	if err != nil {
		panic(err)
	}

	return character
}

func ConfigureCalculators(fileLoc string) (calculators.ShinyModifiers, calculators.Giant, calculators.Stones, int, int) {
	character := parseCharacterFile(fileLoc)

	miningMods := calculators.NewMiningModifiers(
		character.MiningMods.Speed,
		character.MiningMods.FirstStrike,
		character.MiningMods.GiantOdds,
		character.MiningMods.StonesFromMining,
		map[int]int{
			2: character.MiningMods.CartUpgradeLevels.X2Strike,
			3: character.MiningMods.CartUpgradeLevels.X3Strike,
			4: character.MiningMods.CartUpgradeLevels.X4Strike,
			5: character.MiningMods.CartUpgradeLevels.X5Strike,
		},
		character.MiningMods.CartUpgradeLevels.GiantLuck,
		map[int]float64{
			2: character.MiningMods.StrikeOdds.X2Strike,
			3: character.MiningMods.StrikeOdds.X3Strike,
			4: character.MiningMods.StrikeOdds.X4Strike,
			5: character.MiningMods.StrikeOdds.X5Strike,
		},
		character.MiningMods.Overclocks.X2,
		character.MiningMods.Overclocks.X3,
		character.MiningMods.Overclocks.X4,
		character.MiningMods.Overclocks.X5,
	)
	generationMods := calculators.NewEggGenerationModifiers(
		character.GenerationMods.EggLuck,
		character.GenerationMods.CloneLuck,
		character.GenerationMods.CalcifyChance,
		character.getEggIndex(),
		character.GenerationMods.HasRecursiveClone,
	)
	shinyMods := calculators.NewShinyModifiers(character.ShinyLuck)
	giantLuckMods := calculators.NewGiantModifiers(1, 1, 1.1, 1.2, true, false)
	ascensionMods := calculators.NewAscensionModifiers(character.AscensionMods.WingboltLevel)

	giantCalc := calculators.NewGiantCalculator(miningMods, giantLuckMods)
	stoneCalc := calculators.NewStonesCalculator(miningMods, generationMods, ascensionMods)

	return shinyMods, giantCalc, stoneCalc, character.NextMineSpeedCost, character.NextCloneLuckCost
}
