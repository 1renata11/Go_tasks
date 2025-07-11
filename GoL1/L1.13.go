package main

import (
	"fmt"
)

func L113() {
	a := 5
	b := 3
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
