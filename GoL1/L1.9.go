package main

import (
	"fmt"
)

func generate(nums []int, out chan<- int) {
	for _, x := range nums {
		out <- x
	}
	close(out)
}

func multiply(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * 2
	}
	close(out)
}

func L19() {
	numbers := []int{1, 2, 3, 4, 5}
	in := make(chan int)
	out := make(chan int)
	go generate(numbers, in)
	go multiply(in, out)
	for result := range out {
		fmt.Println(result)
	}
}
