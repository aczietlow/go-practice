package main

import (
	"fmt"
)

func main() {
	list := []int{5, 1, 2, 18, 21, 4, 7, 7, 7, 8, 3, 5}
	list2 := list
	quickSort(list, 0, len(list)-1)
	fmt.Println(list)

	fmt.Println(quickSortFunctional(list2))
}

func quickSort(list []int, low, high int) []int {
	if low < high {
		var pivot int
		list, pivot := partition(list, low, high)
		list = quickSort(list, low, pivot-1)
		list = quickSort(list, pivot+1, high)
	}
	return list
}

func partition(list []int, low, high int) ([]int, int) {
	pivot := list[high]
	i := low
	for j := low; j < high; j++ {
		if list[j] < pivot {
			list[i], list[j] = list[j], list[i]
			i++
		}
	}
	list[i], list[high] = list[high], list[i]
	return list, i
}

// The functional approach uses sub-arrays and as a result uses more memory
func quickSortFunctional(list []int) []int {
	if len(list) < 2 {
		return list
	} else {
		pivot := list[len(list)/2]
		left, right, middle := []int{}, []int{}, []int{}
		for _, value := range list {
			if value < pivot {
				left = append(left, value)
			} else if value > pivot {
				right = append(right, value)
			} else {
				middle = append(middle, value)
			}
		}
		return append(append(quickSortFunctional(left), middle...), quickSortFunctional(right)...)
	}
}
