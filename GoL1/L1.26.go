package main

import (
	"fmt"
)

func check(str string) bool {
	mp := make(map[rune]bool)
	for _, x := range str {
		if x >= 97 {
			if mp[x-32] {
				return true
			}
			mp[x-32] = true
		} else {
			if mp[x] {
				return true
			}
			mp[x] = true
		}

	}
	return false
}

func main() {
	str := "Aabc"
	fmt.Print(check(str))
}
