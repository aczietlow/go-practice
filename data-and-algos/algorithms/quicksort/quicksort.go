package main

import (
	"fmt"
)

func main() {
	fmt.Println(quickSortStart([]int{5, 1, 2, 18, 21, 4, 7, 7, 7, 5, 8, 3}))
}

func quickSortStart(list []int) []int {
	return quickSort(list, 0, len(list)-1)
}

func quickSort(list []int, low, high int) []int {
	for low < high {
		pivot := 0
		list, pivot := partition(list, low, high)
		fmt.Println(pivot)
		fmt.Println(list)
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
