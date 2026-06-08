package cyclic_sort

func sortArrayByParityII(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if i%2 == nums[i]%2 {
			continue
		}

		for j := i + 1; j < len(nums); j += 2 {
			if j%2 != nums[j]%2 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}

	}

	return nums
}
