package powercost

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("There are 5 houses. Houses with solar panels: [1, 4]. House connections are like [[1, 2], [2, 3], [2, 4], [4,5]]", func() {
	result := CalculateElectricityCost([]int{1, 4}, [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}})

	//House 1 has a solar panel, so its electricity cost is 0.
	Expect(result[0]).To(Equal(0))
	//House 2 is one wire away from house 1, so its electricity cost is 1.
	Expect(result[1]).To(Equal(1))
	//House 3 is two wires away from the nearest house with a solar panel (house 1), so its electricity cost is 2.
	Expect(result[2]).To(Equal(2))
	//House 4 has a solar panel, so its electricity cost is 0.
	Expect(result[3]).To(Equal(0))
	//House 5 is one wire away from house 4, so its electricity cost is 1.
	Expect(result[4]).To(Equal(1))
})

var _ = Describe("There are 4 houses. Houses with solar panels: [1, 4]. House connections are like [[1, 2], [2, 3], [2, 4], [4,5]]", func() {
	result := CalculateElectricityCost([]int{1, 3}, [][]int{{1, 2}, {2, 3}, {2, 4}})

	fmt.Println(result)

	//House 1 has a solar panel, so its electricity cost is 0.
	Expect(result[0]).To(Equal(0))
	//House 2 is one wire away from house 1, so its electricity cost is 1.
	Expect(result[1]).To(Equal(1))
	//House 3 has a solar panel, so its electricity cost is 0.
	Expect(result[2]).To(Equal(0))
	//House 4 is two wires away from the nearest house with a solar panel (house 1), so its electricity cost is 2.
	Expect(result[3]).To(Equal(2))
})
