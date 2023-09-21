package powercost

import (
	"math"
)

func CalculateElectricityCost(solarPanelHouses []int, connections [][]int) []int {
	numHouses := int(math.Max(float64(len(connections)+1), float64(len(solarPanelHouses))))
	electricityCosts := make([]int, numHouses)

	for i := range electricityCosts {
		electricityCosts[i] = math.MaxInt
	}

	for _, solarPanelHouse := range solarPanelHouses {
		electricityCosts[solarPanelHouse-1] = 0
	}

	for _, connection := range connections {
		house1, house2 := connection[0], connection[1]
		if electricityCosts[house2-1] > electricityCosts[house1-1]+1 {
			electricityCosts[house2-1] = electricityCosts[house1-1] + 1
		}
	}

	return electricityCosts
}
