package main

import "fmt"

func BinarySort(mas []int, element int) int {
	left := -1
	right := len(mas)
	for left < right-1 {
		middle := (left + right) / 2
		if mas[middle] < element {
			left = middle
		} else {
			right = middle
		}
	}
	if mas[right] == element {
		return right
	}
	return -1
}

func L117() {
	mas := []int{1, 2, 3, 4, 6}
	element := 5
	fmt.Print(BinarySort(mas, element))
}
