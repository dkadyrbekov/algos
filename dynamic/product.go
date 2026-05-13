package dynamic

import "math"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}

	if len(nums) == 1 {
		return nums[0]
	}

	curMax := nums[0]
	//curMin := nums[0]
	maxP := curMax
	NegativeC := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			maxP = maxInt(maxP, 0)
			maxP = maxInt(curMax, maxP)
			NegativeC = 0

			if i < len(nums)-1 {
				curMax = nums[i+1]
				//curMin = nums[i+1]
				i++
			}
			continue
		}

		if nums[i] < 0 {
			NegativeC++
		}

		//case nums[i] > 0:
		//	if curMax < curMax*nums[i] {
		//		curMax = curMax * nums[i]
		//	}
		//	if curMin > curMin*nums[i] {
		//		curMin = curMin * nums[i]
		//	}
		//case nums[i] < 0:
		//	NegativeC++
		//	if NegativeC%2 == 0 {
		//		if curMin > curMax*nums[i] {
		//			curMin = curMax * nums[i]
		//		}
		//	} else {
		//		if curMax < curMax*nums[i] {
		//			curMax = curMax * nums[i]
		//		}
		//		if curMin > curMin*nums[i] {
		//			curMin = curMin * nums[i]
		//		}
		//	}
		//}
	}

	return maxInt(curMax, maxP)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

//import "math"
//
//func maxProduct(nums []int) int {
//	if len(nums) == 0 {
//		return math.MinInt32
//	}
//
//	if len(nums) == 1 {
//		return nums[0]
//	}
//
//	zeroesI := make([]int, 0)
//
//	for i, num := range nums {
//		if num == 0 {
//			zeroesI = append(zeroesI, i)
//		}
//	}
//
//	if len(zeroesI) == 0 {
//		maxP := math.MinInt32
//
//		tmp := maxPWithoutZeroes(nums)
//		if maxP < tmp {
//			maxP = tmp
//		}
//
//		return maxP
//	}
//
//	maxP := 0
//
//	for _, i := range zeroesI {
//		tmp := maxPWithoutZeroes(nums[i+1:])
//		if maxP < tmp {
//			maxP = tmp
//		}
//	}
//
//	return maxP
//}
//
//func maxPWithoutZeroes(nums []int) int {
//	if len(nums) == 0 {
//		return math.MinInt32
//	}
//
//	if len(nums) == 1 {
//		return nums[0]
//	}
//
//	maxP := math.MinInt32
//
//	firstNegI := -1
//	lastNegI := -1
//
//	if nums[0] < 0 {
//		firstNegI = 0
//		lastNegI = 0
//	}
//
//	product := nums[0]
//	for i := 1; i < len(nums); i++ {
//		product *= nums[i]
//
//		if product < 0 {
//			lastNegI = i
//
//			if firstNegI == -1 {
//				firstNegI = i
//			}
//		}
//	}
//
//	if product > 0 {
//		return product
//	}
//
//	if firstNegI != -1 && lastNegI != -1 {
//		tmp := maxProductSubArr(nums, firstNegI)
//		if maxP < tmp {
//			maxP = tmp
//		}
//
//		tmp = maxProductSubArr(nums, lastNegI)
//		if maxP < tmp {
//			maxP = tmp
//		}
//	}
//
//	return maxP
//}
//
//func maxProductSubArr(nums []int, i int) int {
//	maxP := math.MinInt32
//
//	if i != len(nums)-1 {
//		tmp := maxPWithoutZeroes(nums[i+1:])
//		if maxP < tmp {
//			maxP = tmp
//		}
//	}
//
//	tmp := maxPWithoutZeroes(nums[:i])
//	if maxP < tmp {
//		maxP = tmp
//	}
//
//	return maxP
//}
