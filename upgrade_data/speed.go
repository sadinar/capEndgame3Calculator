package upgrade_data

const PerStepSpeedImprovement = 0.0025

func GetSpeedPrices() map[int]int {
	return map[int]int{
		1:  60000,
		23: 1800000,
	}
}
