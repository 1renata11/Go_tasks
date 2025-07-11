package main

import "fmt"

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = 0
	return slice[:len(slice)-1]
}

func L123() {
	slice := []int{1, 2, 3}
	fmt.Print(remove(slice, 1))
}
