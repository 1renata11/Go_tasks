package main

import "fmt"

func QuickSort(mas []int) []int {
	if len(mas) < 2 {
		return mas
	}
	pivot := len(mas) / 2
	var left []int
	var right []int
	var middle []int
	for _, x := range mas {
		if x < mas[pivot] {
			left = append(left, x)
		} else if mas[pivot] == x {
			middle = append(middle, x)
		} else {
			right = append(right, x)
		}
	}
	return append(append(QuickSort(left), middle...), QuickSort(right)...)

}

func L116() {
	mas := []int{2, 3, 9, 1, 5, 2, 6, 3}
	fmt.Print(QuickSort(mas))
}
