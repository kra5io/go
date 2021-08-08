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
		first_name: "bobi",
		last_name:  "turboto",
		favourite_ice_cream_flavors: []string{
			"vanilla",
			"cherry",
			"chocolate"},
	}

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.first_name)
		fmt.Println(v.last_name)
		for i, val := range v.favourite_ice_cream_flavors {
			fmt.Println("\t", i, val)
		}
		fmt.Println("---------")
	}
}
