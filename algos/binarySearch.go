package main

import "fmt"

func main() {
	fmt.Printf("%v\n", binarySearch([]int{2, 4, 5, 7, 19, 21}, 7))
}

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list)
	for low <= high {
		// go will round here, bc mid is a type int
		mid := (low + high) / 2
		if item == list[mid] {
			return mid
		} else if item < list[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
