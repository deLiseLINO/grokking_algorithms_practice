package main

import "fmt"

// Recursive function
// takes slice
// returns amount its elements
func count(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + count(arr[1:])
}

func main() {
	arr := []int{1, 2, 1, 2, 4, 6, 7} // 7
	fmt.Println(count(arr))
}
