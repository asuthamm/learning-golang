package main

import "fmt"

const rate float64 = 0.453592

type person struct {
	name   string
	sex    string
	age    uint16
	weight float64
}

// lb->kilo method with value receiver
// func (guy person) lb2kilo() float64 {
// 	// const rate float64 = 0.453592
// 	return float64(guy.weight * rate)
// }

// lb->kilo method with pointer receiver
func (guy *person) lb2kilo() float64 {
	// const rate float64 = 0.453592
	guy.weight = guy.weight * rate
	return float64(guy.weight)
}

func main() {
	// alex := person{"alex", "M", 59}
	alex := person{
		name:   "alex",
		sex:    "M",
		age:    59,
		weight: 160,
	}
	fmt.Println(alex.weight)
	fmt.Println(alex.lb2kilo())
	fmt.Println(alex.weight)
}
