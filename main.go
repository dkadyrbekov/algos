package main

import "fmt"

func printAny[T any](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func SumIntsOrFloats[K comparable, V ~int | ~float32](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	printAny([]int{1, 2, 3, 4, 5})
	printAny([]string{"123", "123", "123"})

	m := make(map[int]float32)
	SumIntsOrFloats(m)
}
