package main

import "fmt"

func main() {
	fmt.Println(sum([]int{1, 3, 5, 2, 12, 8}))
	fmt.Printf("factorial of %d is %d", 5, fact(5))
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

func fact(x int) int {
	if x == 1 {
		return x
	} else {
		return x * fact(x-1)
	}
}
