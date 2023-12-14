package upgrade_data

const PerStepStrikeImprovement = 0.01

func GetSingleStrikePrices() map[int]int {
	return map[int]int{
		1:  10000,
		2:  20000,
		3:  30000,
		4:  40000,
		5:  50000,
		6:  60000,
		7:  70000,
		8:  80000,
		9:  90000,
		10: 100000,
	}
}

func GetDoubleStrikePrices() map[int]int {
	return map[int]int{
		1:  20000,
		2:  40000,
		3:  60000,
		4:  80000,
		5:  100000,
		6:  120000,
		7:  140000,
		8:  160000,
		9:  180000,
		10: 200000,
	}
}

func GetTripleStrikePrices() map[int]int {
	return map[int]int{
		1:  30000,
		2:  60000,
		3:  90000,
		4:  120000,
		5:  150000,
		6:  180000,
		7:  210000,
		8:  240000,
		9:  270000,
		10: 300000,
	}
}

func GetQuadrupleStrikePrices() map[int]int {
	return map[int]int{
		1:  40000,
		2:  80000,
		3:  120000,
		4:  160000,
		5:  200000,
		6:  240000,
		7:  280000,
		8:  320000,
		9:  360000,
		10: 400000,
	}
}

func GetQuintupleStrikePrices() map[int]int {
	return map[int]int{
		1:  50000,
		2:  100000,
		3:  150000,
		4:  200000,
		5:  250000,
		6:  300000,
		7:  350000,
		8:  400000,
		9:  450000,
		10: 500000,
	}
}
