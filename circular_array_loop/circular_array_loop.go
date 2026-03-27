package circular_array_loop

func circularArrayLoop(nums []int) bool {

	// Replace this placeholder return statement with your code
	for i, _ := range nums {
		var cycleIsPositive bool

		switch {
		case nums[i] > 0:
			cycleIsPositive = true
		case nums[i] < 0:
			cycleIsPositive = false
		case nums[i] == 0:
			panic("invalid 0 value in input slice")
		}

		visitedIndexes := make(map[int]bool)
		currentIndex := i

		for {
			if cycleIsPositive && nums[currentIndex] < 0 {
				break
			}

			if !cycleIsPositive && nums[currentIndex] > 0 {
				break
			}

			if visitedIndexes[currentIndex] {
				return true
			}
			visitedIndexes[currentIndex] = true

			nextIndex := (len(nums) + (currentIndex+nums[currentIndex])%len(nums)) % len(nums)

			if currentIndex == nextIndex {
				break
			}

			currentIndex = nextIndex
		}

	}

	return false
}
