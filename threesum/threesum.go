package threesum

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	slices.Sort(nums)

	var lastNotNegativeIndex int

	for i, val := range nums {
		if val >= 0 {
			lastNotNegativeIndex = i
			break
		}
	}

	if lastNotNegativeIndex == 0 {
		return result
	}

	result = append(result, search(nums[:lastNotNegativeIndex], nums)...)
	result = append(result, search(nums[lastNotNegativeIndex:], nums)...)

	return result
}

func search(croppedArr, originalArr []int) [][]int {
	var result [][]int

	for i := 0; i < len(croppedArr); i++ {
		if i > 0 && croppedArr[i] == croppedArr[i-1] {
			continue
		}

		for j := i + 1; j < len(croppedArr); j++ {
			if j > i+1 && croppedArr[j] == croppedArr[j-1] {
				continue
			}

			searchNumber := croppedArr[i] + croppedArr[j]

			k, found := slices.BinarySearch(originalArr, -searchNumber)

			if found {
				result = append(result, []int{croppedArr[i], croppedArr[j], originalArr[k]})
			}
		}
	}

	return result
}
