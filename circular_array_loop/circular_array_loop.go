package circular_array_loop

func circularArrayLoop(nums []int) bool {

	for i := range nums {
		isForward := nums[i] > 0

		fast, slow := i, i

		for {
			if nums[fast]%len(nums) == 0 {
				break
			}

			slow = (len(nums) + (slow+nums[slow])%len(nums)) % len(nums)

			fast = (len(nums) + (fast+nums[fast])%len(nums)) % len(nums)
			if (nums[fast] > 0) != isForward {
				break
			}
			if nums[fast]%len(nums) == 0 {
				break
			}

			fast = (len(nums) + (fast+nums[fast])%len(nums)) % len(nums)
			if (nums[fast] > 0) != isForward {
				break
			}
			if nums[fast]%len(nums) == 0 {
				break
			}

			if fast == slow {
				return true
			}
		}

	}

	return false
}
