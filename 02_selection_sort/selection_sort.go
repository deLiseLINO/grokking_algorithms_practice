package main

import "fmt"

// Takes unsorted slice, finds and returns
// index of the smallest value in it
func find_smallest(arr []int) int {
	// Stores the index of the smallest value
	smallest_index := 0

	// Stores the smallest value
	smallest := arr[smallest_index]
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

// Takes unsorted slice, returns sorted slice
func selection_sort(arr []int) []int {
	size := len(arr)

	// Make new slice with
	// capacity of the old one
	new_arr := make([]int, 0, size)
	for i := 0; i < size; i++ {
		smallest := find_smallest(arr)
		fmt.Println(i)
		// Append the smalles value to the new slice
		new_arr = append(new_arr, arr[smallest])

		// Pop it from the old slice
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return new_arr
}

func main() {
	arr := []int{9, 8, 125, 1, 5, 7, 3} // [1 3 5 7 8 9 125]
	fmt.Println(selection_sort(arr))
}
