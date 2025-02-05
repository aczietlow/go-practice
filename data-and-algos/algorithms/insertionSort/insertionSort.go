package main

import "fmt"

func main() {
	fmt.Println(insertionSort([]int{36, 24, 10, 6, 12}))
}

func insertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0 && list[j-1] > list[j]; j-- {
			list[j], list[j-1] = list[j-1], list[j]
		}
	}

	return list
}
