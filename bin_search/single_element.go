package bin_search

func singleNonDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2

	for left <= right {
		mid = (left + right) / 2

		if (right-left+1)/2%2 == 1 {
			if mid == left || nums[mid] != nums[mid-1] {
				if mid == right || nums[mid] != nums[mid+1] {
					return nums[mid]
				}
				right = mid - 1
				continue
			} else {
				left = mid + 1
				continue
			}
		} else {
			if mid == left || nums[mid] != nums[mid-1] {
				if mid == right || nums[mid] != nums[mid+1] {
					return nums[mid]
				}
				left = mid + 2
				continue
			} else {
				right = mid - 2
				continue
			}
		}
	}

	return nums[mid]
}
