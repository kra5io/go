// write a sum function that works with all types

package main

import (
	"fmt"
)

func main() {
	a1 := []int{5, 5, 5}
	a2 := []float64{5.0, 5.0, 5.0}

	sum(a1)
	sum(a2)
}

type kur interface {
}

func sum(s interface{}) {
	switch s.(type) {
	case []int:
		sum := 0
		for _, i := range s {
			sum += i
		}
		fmt.Println(sum)
	case []float64:
		sum := 0.0
		for _, i := range s {
			sum += i
		}
		fmt.Println(sum)
	}
}
