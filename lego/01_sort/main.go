package main

import "fmt"

func main() {
	input := []int{1, 7, 4, 5, 4, 9, 1, 3, 3, 2, 4, 8, 3, 8, 1} //output = [1, 1, 1, 7, 4, 4, 4, 5, 9, 3, 3, 3, 2, 8, 8]
	var output []int

	for _, n := range input {
		if notin(output, n) {
			for _, m := range input {
				if n == m {
					output = append(output, m)
				}
			}
		}
	}
	fmt.Println(output)
}

func notin(output []int, n int) bool {
	for _, a := range output {
		if a == n {
			return false
		}
	}
	return true
}
