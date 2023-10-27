package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "aabcccccaaa"

	var (
		newString string
		prevChar  int32
	)

	count := 1
	for key, value := range str1 {
		if key > 0 && prevChar == value {
			count++
		} else {
			count = 1
		}
		if key+1 == len(str1) || byte(value) != str1[key+1] {
			newString += string(value) + strconv.FormatInt(int64(count), 10)
		}
		prevChar = value
	}

	if len(newString) < len(str1) {
		fmt.Println(newString)
	} else {
		fmt.Println(str1)
	}
}
