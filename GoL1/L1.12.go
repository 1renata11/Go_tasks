package main

import (
	"fmt"
)

func L112() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)
	for _, s := range strings {
		set[s] = true
	}
	var set_r []string
	for key, _ := range set {
		set_r = append(set_r, key)
	}

	fmt.Println(set_r)
}
