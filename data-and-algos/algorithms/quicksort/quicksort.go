package main

import (
	"fmt"
)

func main() {
	fmt.Println(quicksort([]int{5, 1, 2, 18, 21, 4, 7, 7, 7, 5, 8, 3}))
}

func quicksort(list []int) []int {
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
		return append(append(quicksort(left), middle...), quicksort(right)...)
	}
}

func partition() {

}
