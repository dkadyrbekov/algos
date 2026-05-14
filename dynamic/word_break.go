package dynamic

var suffixMap map[string]bool

func wordBreak(s string, wordDict []string) bool {
	suffixMap = make(map[string]bool)

	return wordBreak2("", s, wordDict)
}

func wordBreak2(suffix, s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	suffixMap[suffix] = true

	for _, word := range wordDict {
		currSuffix := suffix + word
		if suffixMap[currSuffix] {
			continue
		}

		if !matches(word, s) {
			continue
		}

		if wordBreak2(currSuffix, s[len(word):], wordDict) {
			return true
		}
	}

	return false
}

func matches(word, s string) bool {
	if len(word) > len(s) {
		return false
	}

	for i := range word {
		if word[i] != s[i] {
			return false
		}
	}

	return true
}
