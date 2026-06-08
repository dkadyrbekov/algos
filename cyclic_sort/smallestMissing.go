package cyclic_sort

func smallestMissingPositiveInteger(nums []int) int {

	n := len(nums)
	j := 0
	for i := 0; i < n-j; {
		if nums[i] > 0 {
			i++
			continue
		}

		nums[i], nums[n-j-1] = nums[n-j-1], nums[i]
		j++
	}
	nums = nums[:n-j]
	n = len(nums)

	if len(nums) == 0 {
		return 1
	}

	for i := 0; i < n; {
		if nums[i]-1 == i || nums[i] > n || nums[nums[i]-1] == nums[i] {
			i++
			continue
		}

		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}
