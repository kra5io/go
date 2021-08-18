package main

import "fmt"

func foo(x ...int) int {
	sum := 0
	for _, i := range x {
		sum += i
	}
	return sum
}

func bar(y []int) int {
	sum := 0
	for _, i := range y {
		sum += i
	}
	return sum
}

func main() {
	defer fmt.Println(foo([]int{5, 4, 3, 2, 1}...))
	fmt.Println(bar([]int{5, 4, 3, 2}))
}
