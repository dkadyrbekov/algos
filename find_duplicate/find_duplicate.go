package find_duplicate

func findDuplicate(nums []int) int {
	n := len(nums)

	switch n {
	case 0, 1:
		return -1
	case 2:
		return nums[0]
	}

	duplicateSum := 0
	maxNum := 0
	for _, num := range nums {
		duplicateSum += num
		if num > maxNum {
			maxNum = num
		}
	}
	uniqueSum := (maxNum + 1) * maxNum / 2

	duplicateCount := maxNum - len(nums)

	return (duplicateSum - uniqueSum) / duplicateCount
}
