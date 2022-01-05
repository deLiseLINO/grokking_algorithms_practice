package main

import "fmt"

// Recursive function
// takes slice, returns
// its highest value
func highest(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if arr[0] > highest(arr[1:]) {
		return arr[0]
	}
	return highest(arr[1:])
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 10} // 10
	fmt.Println(highest(arr))
}
