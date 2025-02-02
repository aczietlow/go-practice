package main

import "fmt"

// Sort the list by find a given element in a list, copying it to a new list, AND removing it from the unsorted list
func main() {
	list := []int{5, 3, 2, 1, 21, 13, 8}
	fmt.Println(selectionSort(list))
}

func selectionSort(list []int) []int {
	newList := []int{}
	for range list {
		smallest, smallestIndex := findSmallest(list)
		list = append(list[:smallestIndex], list[smallestIndex+1:]...)
		newList = append(newList, smallest)
	}
	return newList
}

func findSmallest(list []int) (int, int) {
	var smallest, smallestIndex int = list[0], 0
	for i, number := range list {
		if number < smallest {
			smallest = number
			smallestIndex = i
		}
	}
	return smallest, smallestIndex
}
