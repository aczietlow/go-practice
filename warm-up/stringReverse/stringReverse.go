package main

import "fmt"

func main() {
	// In go stings are collection of runes in the form of a slice.
	// Therefore, it should be able to reverse it

	input := "Hello World!"

	fmt.Print(reverseString(input))
}

func reverseString(inputString string) string {
	// create a slice of the same length
	stringSlice := make([]rune, len(inputString))

	for key, value := range inputString {
		index := len(stringSlice) - (key + 1)
		stringSlice[index] = value
	}

	return string(stringSlice)
}
