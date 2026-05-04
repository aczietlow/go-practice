package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("test")
	string1 := "Hello world, hello universe."
	fmt.Println(countWords(string1))
}

func countWords(s string) int {
	words := strings.Split(s, " ")
	return len(words)
}
