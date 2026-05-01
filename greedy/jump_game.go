package greedy

func jumpGame(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	if nums[0] >= len(nums)-1 {
		return true
	}

	maxIndex := 0
	maxVal := 0

	for i := 1; i <= nums[0]; i++ {
		if nums[i] > maxVal {
			maxIndex = i
			maxVal = nums[i]
		}
	}

	if maxIndex == 0 {
		return false
	}

	return jumpGame(nums[maxIndex:])
}
