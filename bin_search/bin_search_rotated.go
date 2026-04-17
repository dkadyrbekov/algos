package bin_search

func binarySearchRotated(nums []int, target int) int {
	n := len(nums)

	left := 0
	right := n - 1

	for {
		if left > right {
			return -1
		}

		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < nums[right] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
}
