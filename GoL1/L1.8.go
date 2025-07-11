package main

import "fmt"

func SetBit(val int64, b int64, i int) int64 {
	a := int64(1 << (i - 1))
	if b == 1 {
		return val | a
	} else {
		return val &^ a
	}
}

func L18() {
	a := int64(0b1101)
	fmt.Println(SetBit(a, 0, 3))
}
