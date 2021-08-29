package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go kur()
	bar()

	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func foo() {
	for i := -9; i < 1; i++ {
		fmt.Println("foo:", i)
	}
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Done()
}

func kur() {
	for i := 1; i < 11; i++ {
		fmt.Println("kur:", i)
	}
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Done()
}

func bar() {
	for i := 1; i < 11; i++ {
		fmt.Println("bar:", i)
	}

}
