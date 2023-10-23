package main

import "fmt"

func main() {
	// In go stings are collection of runes in the form of a slice.
	// Therefore, it should be able to reverse it

	input := "Hello World!"

	fmt.Print(reverseString(input))
}

func reverseString(inputString string) string {
	// create a slice of the same length by casting the string as a slice of runes.
	stringSlice := []rune(inputString)

	i, j := 0, len(inputString)-1

	for i < j {
		stringSlice[i], stringSlice[j] = stringSlice[j], stringSlice[i]
		i++
		j--
	}

	return string(stringSlice)
}
