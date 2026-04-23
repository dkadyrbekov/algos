package bin_search

func findClosestElements(nums []int, k int, target int) []int {
	if len(nums) == 0 {
		return nums
	}

	closest := findClosest(nums, target)

	left := closest - 1
	right := closest + 1
	for {
		if left < 0 {
			right = minInt(k, len(nums))

			return nums[0:right]
		}

		if right > len(nums)-1 {
			left = maxInt(0, len(nums)-k)

			return nums[left:]
		}
		if right-left-1 >= k {
			return nums[left+1 : right]
		}

		if abs(target-nums[left]) <= abs(target-nums[right]) {
			left = left - 1
		} else {
			right = right + 1
		}

	}
}

func findClosest(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	closest := 0

	for left <= right {
		closest = (left + right) / 2

		if nums[closest] == target {
			break
		}

		if right-left <= 1 {
			if abs(target-nums[right]) < abs(target-nums[left]) {
				closest = right
			} else {
				closest = left
			}

			break
		}

		if nums[closest] < target {
			left = closest
		} else {
			right = closest
		}
	}

	return closest
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
