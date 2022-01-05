package main

import "fmt"

// Quck sort function, returns sorted slice
func quick_sort(arr []int) []int {

	// Base case, arrays with 1 or 0 length are already sorted
	if len(arr) < 2 {
		return arr
	} else {
		// Recursive case
		pivot := arr[0]

		// 2 slices that store values less
		// or greater than the pivot
		less := []int{}
		greater := []int{}
		for _, num := range arr[1:] {
			if num < pivot {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}
		less = append(quick_sort(less), pivot)
		greater = quick_sort(greater)
		return (append(less, greater...))
	}
}

func main() {
	arr := []int{20, 3, 1}
	fmt.Println(quick_sort(arr))
}
