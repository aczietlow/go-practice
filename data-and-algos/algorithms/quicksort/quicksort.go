package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(quicksort([]int{5, 1, 2, 18, 21, 4}))
}

func quicksort(list []int) []int {
	if len(list) < 2 {
		return list
	} else {
		// Select a random Pivot to get better average time complexity
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomIndex := r.Intn(len(list))
		pivot := list[randomIndex]
		left, right := []int{}, []int{}
		for _, value := range list {
			if value < pivot {
				left = append(left, value)
			} else if value > pivot {
				right = append(right, value)
			}
		}
		return append(append(quicksort(left), pivot), quicksort(right)...)
	}
}
