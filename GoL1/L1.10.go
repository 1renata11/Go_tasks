package main

import "fmt"

func L110() {
	var numbers = []float32{25.1, -0.6, 1.5}
	mp := make(map[int][]float32)
	for _, x := range numbers {
		mp[int(x/10)*10] = append(mp[int(x/10)*10], x)
	}
	fmt.Println(mp)
}
