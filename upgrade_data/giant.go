package upgrade_data

const PerStepGiantLuckImprovement = 0.001

func GetGiantLuckPrices() map[int]int {
	return map[int]int{
		1:  100000,
		2:  1000000,
		3:  10000000,
		4:  100000000,
		5:  1000000000,
		6:  10000000000,
		7:  100000000000,
		8:  1000000000000,
		9:  10000000000000,
		10: 100000000000000,
	}
}
