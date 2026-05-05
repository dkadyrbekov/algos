package dynamic

import "math"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}

	if len(nums) == 1 {
		return nums[0]
	}

	zeroesI := make([]int, 0)

	for i, num := range nums {
		if num == 0 {
			zeroesI = append(zeroesI, i)
		}
	}

	if len(zeroesI) == 0 {
		maxP := math.MinInt32

		tmp := maxPWithoutZeroes(nums)
		if maxP < tmp {
			maxP = tmp
		}

		return maxP
	}

	maxP := 0

	for _, i := range zeroesI {
		tmp := maxPWithoutZeroes(nums[i+1:])
		if maxP < tmp {
			maxP = tmp
		}
	}

	return maxP
}

func maxPWithoutZeroes(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}

	if len(nums) == 1 {
		return nums[0]
	}

	maxP := math.MinInt32

	firstNegI := -1
	lastNegI := -1

	if nums[0] < 0 {
		firstNegI = 0
		lastNegI = 0
	}

	product := nums[0]
	for i := 1; i < len(nums); i++ {
		product *= nums[i]

		if product < 0 {
			lastNegI = i

			if firstNegI == -1 {
				firstNegI = i
			}
		}
	}

	if product > 0 {
		return product
	}

	if firstNegI != -1 && lastNegI != -1 {
		tmp := maxProductSubArr(nums, firstNegI)
		if maxP < tmp {
			maxP = tmp
		}

		tmp = maxProductSubArr(nums, lastNegI)
		if maxP < tmp {
			maxP = tmp
		}
	}

	return maxP
}

func maxProductSubArr(nums []int, i int) int {
	maxP := math.MinInt32

	if i != len(nums)-1 {
		tmp := maxPWithoutZeroes(nums[i+1:])
		if maxP < tmp {
			maxP = tmp
		}
	}

	tmp := maxPWithoutZeroes(nums[:i])
	if maxP < tmp {
		maxP = tmp
	}

	return maxP
}
