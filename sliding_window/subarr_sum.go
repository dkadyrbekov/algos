package sliding_window

func minSubArrayLen(target int, nums []int) int {
	minLen := len(nums) + 1

	if len(nums) == 0 {
		return 0
	}
	if target <= 0 {
		return 0
	}

	left, right := 0, 1
	windowSum := nums[left]

	if windowSum >= target {
		return 1
	}

	for right < len(nums) {
		windowSum += nums[right]

		for windowSum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}

			if minLen == 1 {
				return 1
			}

			windowSum -= nums[left]
			left++
		}

		right++
	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}
