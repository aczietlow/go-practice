package main

import "fmt"

func main() {
	fmt.Println(bSort([]int{2, 4, 1, 5, 9, 3, 7, 12, 21, 33, 6}))
}

func bSort(list []int) []int {
	for start := 0; start < len(list)-1; start++ {
		index := start
		for end := len(list) - 1; end > index; end-- {
			if list[end] < list[end-1] {
				list[end], list[end-1] = list[end-1], list[end]
			}
		}
	}
	return list
}
