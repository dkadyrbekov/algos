package dynamic

import (
	"fmt"
	"strconv"
)

func numOfDecodings(decodeStr string) int {
	if len(decodeStr) == 0 {
		return 0
	}

	counter := 1
	prev := ""
	prevC := 0
	for i := 0; i < len(decodeStr); i++ {
		valid, err := validNum(prev + string(decodeStr[i]))
		if err != nil {
			return 0
		}
		if valid {
			tmp := counter
			counter += prevC
			prevC = tmp
		}
		prev = string(decodeStr[i])
	}

	return counter
}

func validNum(numStr string) (bool, error) {
	if len(numStr) == 0 {
		return false, fmt.Errorf("empty str input")
	}

	if numStr[0] == '0' || numStr == "00" {
		return false, fmt.Errorf("leading zeros")
	}

	num, err := strconv.Atoi(numStr)

	return err == nil && num >= 1 && num <= 26, nil
}
