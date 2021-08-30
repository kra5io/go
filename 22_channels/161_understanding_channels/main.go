package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	c := make(chan int, 2) # buffer
// 	c <- 42
// 	fmt.Println(<-c)

// 	c <- 43
// 	c <- 44
// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// }
