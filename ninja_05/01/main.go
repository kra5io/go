package main

import "fmt"

type person struct {
	first_name                  string
	last_name                   string
	favourite_ice_cream_flavors []string
}

func main() {
	p1 := person{
		first_name: "pepi",
		last_name:  "popa",
		favourite_ice_cream_flavors: []string{
			"vanilla",
			"cherry",
			"chocolate",
		},
	}
	p2 := person{
		first_name: "mil4o",
		last_name:  "balak",
		favourite_ice_cream_flavors: []string{
			"vanilla",
			"cherry",
			"chocolate"},
	}
	fmt.Println(p1.first_name)
	fmt.Println(p1.last_name)
	for k, v := range p1.favourite_ice_cream_flavors {
		fmt.Println(k, v)
	}
	fmt.Println(p2.first_name)
	fmt.Println(p2.last_name)
	for k, v := range p2.favourite_ice_cream_flavors {
		fmt.Println(k, v)
	}
}
