package main

import "fmt"

func main() {
	// Create a slice of integers containing numbers 1 to 5. Add the number 6 to it and print the result.
	addNumberToSlice(6)

	// Create a slice of strings with the names of three cities. Change the second city and print the modified slice.
	changeCity("toledo")
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
