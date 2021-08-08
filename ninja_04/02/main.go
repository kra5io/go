package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 32, 15, 12, 15, 55}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
