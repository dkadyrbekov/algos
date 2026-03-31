package sliding_window

//func longestRepeatingCharacterReplacement(s string, k int) int {
//	runeStr := []rune(s)
//	if len(runeStr) == 0 {
//		return 0
//	}
//
//	maxLength := 1
//	runesMap := map[rune]int{
//		runeStr[0]: 1,
//	}
//	maxRuneCount := 1
//
//	left := 0
//	right := 0
//
//	for right < len(runeStr)-1 {
//		right++
//		runesMap[runeStr[right]]++
//		for _, v := range runesMap {
//			if v > maxRuneCount {
//				maxRuneCount = v
//			}
//		}
//
//		if len(runeStr[left:right+1])-maxRuneCount > k {
//			runesMap[runeStr[left]]--
//
//			left++
//		}
//
//		if right-left+1 > maxLength {
//			maxLength = right - left + 1
//		}
//	}
//
//	return maxLength
//}

func longestRepeatingCharacterReplacement(s string, k int) int {
	stringLength := len(s)
	lengthOfMaxSubstring := 0
	start := 0
	charFreq := make(map[byte]int)
	mostFreqChar := 0

	for end := 0; end < stringLength; end++ {
		if _, ok := charFreq[s[end]]; !ok {
			charFreq[s[end]] = 1
		} else {
			charFreq[s[end]]++
		}

		mostFreqChar = max(mostFreqChar, charFreq[s[end]])

		if end-start+1-mostFreqChar > k {
			charFreq[s[start]]--
			start++
		}

		lengthOfMaxSubstring = max(end-start+1, lengthOfMaxSubstring)
	}

	return lengthOfMaxSubstring
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
