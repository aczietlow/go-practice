package main

import "fmt"

func main() {
	list := []int{1, 3, 5, 2, 12, 8}
	fmt.Println(sum(list))
}

func sum(list []int) int {
	total := 0
	if len(list) == 1 {
		return list[0]
	} else {
		total += list[0] + sum(list[1:])
	}
	return total
}
