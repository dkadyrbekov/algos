package circular_array_loop

func circularArrayLoop(nums []int) bool {

	for i := range nums {
		isForward := nums[i] > 0

		fast, slow := i, i

		for {
			slow = nextIndex(slow, nums[slow], len(nums))

			fast = nextIndex(fast, nums[fast], len(nums))
			if !validStep(nums[fast], isForward, len(nums)) {
				break
			}

			fast = nextIndex(fast, nums[fast], len(nums))
			if !validStep(nums[fast], isForward, len(nums)) {
				break
			}

			if fast == slow {
				return true
			}
		}

	}

	return false
}

func validStep(step int, isForward bool, lenArr int) bool {
	if (step > 0) != isForward {
		return false
	}
	if step%lenArr == 0 {
		return false
	}

	return true
}

func nextIndex(currI, nextI, lenArr int) int {
	return (lenArr + (currI+nextI)%lenArr) % lenArr
}
