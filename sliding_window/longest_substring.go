package sliding_window

func findLongestSubstring(str string) int {
	if len(str) == 0 {
		return 0
	}

	left, right := 0, 0
	maxLen := 1
	lettersMap := map[byte]int{
		str[0]: 1,
	}

	for right < len(str)-1 {
		right++
		lettersMap[str[right]]++

		for lettersMap[str[right]] > 1 {
			lettersMap[str[left]]--
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
