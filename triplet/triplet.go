package triplet

func FindTriplet(arr []int, target int) (bool, []int) {
	n := len(arr)
	if n < 3 {
		return false, nil
	}

	pairSet := make(map[int]bool)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum := arr[i] + arr[j]
			complement := target - sum
			if pairSet[complement] {
				return true, []int{arr[i], arr[j], complement}
			}
		}

		pairSet[arr[i]] = true
	}

	return false, nil

}
