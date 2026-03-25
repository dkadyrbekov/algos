package reverse_words

func reverseWords(sentence string) string {
	var reversedSentence []byte

	sentenceByteArr := append([]byte{' '}, []byte(sentence)...)

	left := len(sentenceByteArr)
	right := len(sentenceByteArr)

	for left > 0 {
		left--

		if sentenceByteArr[left] == ' ' {
			if right-left > 1 {
				word := sentenceByteArr[left+1 : right]

				delim := []byte{' '}
				if len(reversedSentence) == 0 {
					delim = []byte{}
				}

				reversedSentence = append(reversedSentence, append(delim, word...)...)
			}
			right = left
		}
	}

	return string(reversedSentence)
}
