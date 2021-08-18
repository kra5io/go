package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (a person) speak() {
	fmt.Println("hello, my name is", a.first, "and my age is", a.age)
}

func main() {
	a := person{
		first: "kur",
		last:  "cici",
		age:   32,
	}
	a.speak()
}
