package main

import (
	"fmt"
	"time"
)

func appendLen(numbers []*int) {
	size := len(numbers)
	numbers = append(numbers, &size)
}

func main() {
	numbers := make([]*int, 0, 5)
	var number int
	for range 3 {
		number++
		numbers = append(numbers, &number)
	}
	appendLen(numbers)

	for _, number := range numbers {
		fmt.Printf("%d ", *number)
	}

	time.Time{}
}
