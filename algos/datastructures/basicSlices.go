package main

import "fmt"

func main() {
	// Create a slice of integers containing numbers 1 to 5. Add the number 6 to it and print the result.
	addNumberToSlice(6)

	// Create a slice of strings with the names of three cities. Change the second city and print the modified slice.
	changeCity("toledo")

	// Create a slice of integers [10, 20, 30]. Copy it into another slice and modify the copied slice. Print both slices to observe changes.
	copyChangesSlice()
}

func copyChangesSlice() {
	s1 := []int{10, 20, 30}
	s2 := s1
	// var s3 []int
	s3 := make([]int, len(s2))
	copy(s3, s1)

	s1[0] = 33
	fmt.Printf("S1:%v\nS2:%v\nS3:%v", s1, s2, s3)
}

func addNumberToSlice(num int) {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", numbers)
	numbers = append(numbers, num)
	fmt.Printf("%v\n", numbers)
}

func changeCity(city string) {
	cities := []string{"summerville", "Augusta", "Rome"}
	fmt.Printf("%v\n", cities)
	cities[1] = city
	fmt.Printf("%v\n", cities)
}
