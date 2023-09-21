package triplet

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("Finding tripple in an array",
	func(source []int, target int, expectedBool bool, tripletArray []int) {
		ok, value := FindTriplet(source, target)
		Ω(ok).Should(Equal(expectedBool))
		Ω(value).Should(ContainElements(tripletArray))
	},

	Entry("The triple target is availble in the source array - set 1",
		[]int{1, 15, 4, 8, 5, 3, 10}, 8, true, []int{3, 4, 1}),

	Entry("The triple target is not availble in the source array",
		[]int{1, 15, 4, 8, 5, 3, 10}, 100, false, nil),

	Entry("The source is less length 3",
		[]int{1, 15}, 1, false, nil),
)
