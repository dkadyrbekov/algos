package main

import "fmt"

func f(s []int) {
	s = append(s, 4)
	s[0] = 100
}

func main() {
	a := make([]int, 3, 10)
	f(a)
	fmt.Println(a) // [100 2 3]
}
