package main

import "fmt"

type person struct {
	name string
	sex  string
	age  int
}

func main() {
	// alex := person{"alex", "M", 59}
	alex := person{
		name: "alex",
		sex:  "M",
		age:  59,
	}
	fmt.Println(alex.age)
}
