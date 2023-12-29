# Giant Calculator
Calculates the odds of hatching a giant pet and suggests an upgrade path for giant hunting in Collect All Pets.

## Usage
The calculator's constructor takes the player's current state and returns an object which can print the order of all
upgrades, suggest the next upgrade, show the chance per strike of getting a giant, and show the full probability table
for a given time period.

The constructor takes the following arguments:
- overclock configuration
- giant achievement modifier, exactly as shown in game
- rune giant luck modifier, exactly as shown in game
- mining speed, exactly as shown in game
- list of current strike levels
- giant luck level, in game giant luck divided by 0.1. For example, `6.7 % 0.1 = 67`

###### Overclock configuration
The overclock configuration constructor takes the following arguments:
- x2 overclock enabled, true for enabled and false for disabled
- x3 overclock enabled, true for enabled and false for disabled
- x4 overclock enabled, true for enabled and false for disabled
- x5 overclock enabled, true for enabled and false for disabled
- giant luck overclock enabled, true for enabled and false for disabled

###### Strike level list
Map of strike number to upgrade level and expects indexes 2, 3, 4, and 5 to be set. Index two is the x2 strike upgrade
level. Index three is the x3 strike upgrade level. And so on. To determine your strike level, divide the percentage 
shown in game by 0.25. For example, 18.25% would be `18.25 % 0.25 = 73` for a strike level of 73.

``` go
giantCalc := calculators.NewGiantCalculator(
    calculators.NewOverclockConfig(false, false, true, true, true),
    1.07,
    1.188,
    0.64,
    map[int]int{
        2: 73,
        3: 73,
        4: 73,
        5: 73,
    },
    67,
)`
```

###### Suggest next upgrade
`GetNextUpgrade` takes no arguments and will examine all current upgrades and suggest the next upgrade that will
maximize increased giant chance for the minimum number of mythic stones.

###### Giant chance per strike
To calculate the chance of getting a giant in a single strike, call `CalculateChancePerSTrike`. This method takes the
first strike chance as its only argument. First strike chance is a factional number between 0 and 1. For example, if
the first strike chance shown in game is 85.7%, the value passed in would be 0.857.

###### Show probability list
`PrintProbabilityDistribution` takes two arguments, the duration over which to calculate probabilities and first strike
chance. `PrintProbabilityDistribution` will print a table of probabilities which attempts to condense the data into
fewer data points. It will determine a reasonable upper bound to stop reporting probabilities as well as combining
the first group of low probabilities into a single line item. The output includes the median and a breakdown of the 
midpoint information. This method assumes the user will stay on the same egg for the entire duration. Duration is 
expected to be a value greater than 0 which can reduce to one or more seconds. First strike chance is a factional 
number between 0 and 1. For example, if the first strike chance shown in game is 85.7%, the value passed in would be 
0.857.