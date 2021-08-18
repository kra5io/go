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
	fmt.Println(foo([]int{5, 4, 3, 2}...))
	fmt.Println(bar([]int{5, 4, 3, 2}))
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	n := foo(ii...)
// 	fmt.Println(n)

// 	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	n2 := bar(ii2)
// 	fmt.Println(n2)
// }

// func foo(xi ...int) int {
// 	total := 0
// 	for _, v := range xi {
// 		total += v
// 	}
// 	return total
// }

// func bar(xi []int) int {
// 	total := 0
// 	for _, v := range xi {
// 		total += v
// 	}
// 	return total
// }
