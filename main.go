package main

import (
	"fmt"
	subtract "golang/assignment"
)

var name string = "golang practice"

func main() {
	a, b := 50, 20
	fmt.Println(name)
	if a > b {
		subtract.Sub(a, b)
	} else {
		subtract.Sub(b, a)
	}
	var i = 0
	for i < 5 {
		fmt.Println("i value is", i)
		i++
	}

}
