package main

import "fmt"

// Simple loop function
// takes slice and returns
// summ of elements in it
func sum(arr []int) int {
	total := 0
	for _, x := range arr {
		total += x
	}
	return total
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6} // 21
	fmt.Println(sum(arr))
}
