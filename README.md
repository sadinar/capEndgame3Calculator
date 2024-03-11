# CAP Endgame 3 Calculator
Calculates the odds of hatching a giant pet, estimate stones gained, and suggests an upgrade path for giant hunting in Collect All Pets.

## Usage
The stones and giants calculators rely on information about the player's current upgrades which are set through
the `MiningModifiers` and `EggGenerationModifiers` configuration objects.

### Configuration Objects
Both the `MiningModifiers` and `EggGenerationModifiers` configuration objects have constructors which take values directly
from the newly added stats panel in game. All upgrades are no longer required to be listed out since the stats panel
now takes care of all of that work.

In order to analyze the costs of potential upgrades, however, players' current strike and giant levels still need to be
provided. 

###### Strike level list
Map of strike number to upgrade level and expects indexes 2, 3, 4, and 5 to be set. Index two is the x2 strike upgrade 
level. Index three is the x3 strike upgrade level. And so on. To determine your strike level, divide the percentage 
shown in game by 0.25. For example, 18.25% would be 18.25 % 0.25 = 73 for a strike level of 73.

###### Giant luck level
The giant luck percentage shown in game above the giant luck cart divided by 0.1. For example, `6.7 % 0.1 = 67`

``` go
miningMods := calculators.NewMiningModifiers(
		.75+.5, // exactly as on stats screen
		100,    // exactly as shown on the wooden board behind egg
		.149,   // exactly as on stats screen
		408.8,  // stats screen
		map[int]int{
			2: 74,
			3: 74,
			4: 74,
			5: 74,
		},
		70,
		map[int]float64{
			2: 29.4,  // exactly as on stats screen
			3: 8.702, // exactly as on stats screen
			4: 2.898, // exactly as on stats screen
			5: 1.076, // exactly as on stats screen
		},
	)
```


``` go
generationMods := calculators.NewEggGenerationModifiers(
		51,    // as shown on stats screen
		6.1,   // as shown on stats screen
		127.5, // as shown in stats pane
		calculators.MythicEgg,
		true,
	)
```

``` go
shinyMods := calculators.NewShinyModifiers(75.12) // exactly as seen on stats screen
```

### Giant Calculator

###### Suggest next giant upgrade
Given the cost of the next mining speed upgrade, `GetNextUpgrade` on the `Giant` calculator will examine all current 
upgrades and suggest the next upgrade that will maximize increased giant chance for the minimum number of mythic stones.

###### PrintProbabilityMedian
Given a duration and a shiny luck config, `PrintProbabilityMedian` on the `Giant` calculator prints the median number 
of regular and shiny giants for the specified time period.

###### Show probability list
`PrintProbabilityDistribution` on the `Giant` calculator takes two arguments, the duration over which to calculate probabilities and first strike
chance. `PrintProbabilityDistribution` will print a table of probabilities which attempts to condense the data into
fewer data points. It will determine a reasonable upper bound to stop reporting probabilities as well as combining
the first group of low probabilities into a single line item. The output includes the median and a breakdown of the 
midpoint information. This method assumes the user will stay on the same egg for the entire duration. Duration is 
expected to be a value greater than 0 which can reduce to one or more seconds. First strike chance is a factional 
number between 0 and 1. For example, if the first strike chance shown in game is 85.7%, the value passed in would be 
0.857.

### Stones Calculator

###### Suggest next stones upgrade
Given the costs of the next mining speed and clone luck upgrades, `FindNextUpgrade` on the `Stones` calculator will 
examine all upgrades and pick the next upgrade which maximizes additional stones gained for the amount of stones spent.

###### Estimate stones gained
Given a duration, `CalculateCombinedStones` will estimate the number of stones gained during that time period. Returns
two values, one for stones gained through generation and another for stones gained through mining.

###### Estimate pet damage gain
Given a duration, PrintDamageChange will estimate how much each pet class' damage will increase as a result of egg
generation for that duration. This allows a comparison between different egg luck values alongside their stone effect.

```go
giantCalc := calculators.NewGiantCalculator(miningMods, true)
stoneCalc := calculators.NewStonesCalculator(miningMods, generationMods)

giantCalc.GetNextUpgrade(2100000)
giantCalc.PrintProbabilityMedian(duration, shinyMods)

stoneCalc.FindNextUpgrade(2100000, 700000)
stoneCalc.CalculateCombinedStones(duration)
stoneCalc.PrintDamageChange(duration, shinyMods)
```