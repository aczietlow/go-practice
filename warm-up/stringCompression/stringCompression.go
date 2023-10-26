package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "aabcccccaaa"

	var newString string
	count := 1

	for key, value := range str1 {
		if key > 0 && str1[key-1] == str1[key] {
			count++
		} else {
			count = 1
		}
		if key+1 < len(str1) && str1[key] == str1[key+1] {
			newString += string(value) + strconv.FormatInt(int64(count), 10)
		}
	}

	if len(str1) <= len(newString) {
		fmt.Println(newString)
	} else {
		fmt.Println(str1)
	}
}
