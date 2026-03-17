package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}

	bigARR := arr[0:len(arr)]
	fmt.Println(bigARR)
}
