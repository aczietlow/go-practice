package main

import "fmt"

func main() {
	testArray := [4]int{2, 4, 6, 8}

	for i, value := range testArray {
		fmt.Printf("testArray[%d] memory address: %v\n", i, &value)
	}
}
