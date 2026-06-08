package cyclic_sort

func findCorruptPair(nums []int) []int {

	if len(nums) < 2 {
		return []int{}
	}

	var duplicate, missing int

	for i := 0; i < len(nums); {
		if nums[i]-1 == i || nums[i] == -1 {
			i++
			continue
		}

		swI, swJ := nums[i], nums[nums[i]-1]

		nums[i], nums[nums[i]-1] = swJ, swI

		if swI == swJ {
			duplicate = nums[i]
			nums[i] = -1
		}
	}

	for i, v := range nums {
		if v == -1 {
			missing = i + 1
		}
	}

	return []int{missing, duplicate}
}
