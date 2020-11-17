package main

import (
	"fmt"
)

func main() {
	numGenerator := closureFunc()
	// fmt.Print(numGenerator, "\n")
	// fmt.Print("start", "\n")
	for j := 0; j < 5; j++ {
		fmt.Print("=", numGenerator(), "\n")
	}
}

func closureFunc() func() int {
	var i = 0
	fmt.Print(i, "-", "\n")
	return func() int {
		fmt.Print(i, "\t")
		i += 1
		return i
	}
}
