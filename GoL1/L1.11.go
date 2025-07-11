package main

import (
	"fmt"
	"sort"
)

func L111() {
	var numbers1 = []int{1, 2, 3}
	var numbers2 = []int{2, 3, 4}
	var result []int
	sort.Ints(numbers1)
	sort.Ints(numbers2)
	var i = 0
	var j = 0
	for i < len(numbers1) && j < len(numbers2) {
		if numbers1[i] == numbers2[j] {
			result = append(result, numbers1[i])
			i++
			j++
			continue
		}
		if numbers1[i] > numbers2[j] {
			j++
			continue
		}
		if numbers1[i] < numbers2[j] {
			i++
			continue
		}
	}
	fmt.Println(result)
}
