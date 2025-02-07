package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{1, 3, 7, 11, 12, 2, 4, 6}))
}

func mergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	mid := len(list) / 2
	left := mergeSort(list[:mid])
	right := mergeSort(list[mid:])
	return merge(left, right)

}

func merge(leftArray []int, rightArray []int) []int {
	tempArray := []int{}
	i, j := 0, 0
	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] < rightArray[j] {
			tempArray = append(tempArray, leftArray[i])
			i++
		} else {
			tempArray = append(tempArray, rightArray[j])
			j++
		}
	}

	// if one array was longer than the other
	for ; i < len(leftArray); i++ {
		tempArray = append(tempArray, leftArray[i])
	}
	for ; j < len(rightArray); j++ {
		tempArray = append(tempArray, rightArray[j])
	}
	return tempArray
}
