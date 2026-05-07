package main

import "fmt"

func main() {
	// list := []int{1, 4, 17, 28, 3, 5, 5, 5, 1, 7, 12, 23, 28, 38, 47, 12}
	list := []int{5, 4, 3, 2, 1}

	sortedList := quicksort(list)
	fmt.Printf("%v", sortedList)

}

func quicksort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	lower, higher := partition(list, len(list)-1)
	lower = quicksort(lower)
	higher = quicksort(higher)
	fmt.Printf("lower: %v\n", lower)
	fmt.Printf("higher: %v\n", higher)
	var newList []int
	newList = append(newList, lower...)
	newList = append(newList, list[len(list)-1])
	newList = append(newList, higher...)

	return newList
}

func partition(list []int, pivotIndex int) ([]int, []int) {
	pivot := list[pivotIndex]
	lower := 0
	for i := range list {
		if list[i] < pivot {
			list[lower], list[i] = list[i], list[lower]
			lower++
		}
	}
	// return lower and higher partitioned subarray.
	return list[0:lower], list[lower:pivotIndex]
}
