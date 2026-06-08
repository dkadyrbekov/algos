package cyclic_sort

func findMissingNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	nums = append(nums, len(nums))

	for i := 0; i < len(nums); {
		if nums[nums[i]] == nums[i] {
			i++
			continue
		}

		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}

	for i, v := range nums {
		if i != v {
			return i
		}
	}

	return len(nums) - 1
}
