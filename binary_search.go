package main

import "fmt"

// Takes sorted list with values and a value needs to be found
// Returns index of this value in the list
func binary_search(list []int, item int) int {
	// The lowest and the highest index in the list
	low := 0
	high := len(list) - 1

	// While we still have some elements
	for low <= high {

		// Guess would be the middle of the list
		guess := (low + high) / 2

		// Returns the index of needed value
		if list[guess] == item {
			return guess
		}

		// If guess value was too high or too low
		if list[guess] > item {
			high = guess - 1
		} else {
			low = guess + 1
		}
	}

	// Unable to find element in the list, return -1
	return -1
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binary_search(list, 4))  // 3
	fmt.Println(binary_search(list, -1)) // -1
}
