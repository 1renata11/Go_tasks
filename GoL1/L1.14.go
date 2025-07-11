package main

import (
	"fmt"
)

func detectType(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println("Тип: int", val)
	case string:
		fmt.Println("Тип: string", val)
	case bool:
		fmt.Println("Тип: bool", val)
	case chan int:
		fmt.Println("Тип: chan int")
	case chan string:
		fmt.Println("Тип: chan string")
	case chan bool:
		fmt.Println("Тип: chan bool")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func L114() {
	detectType(42)
	detectType("hello")
	detectType(true)
	c1 := make(chan int)
	c2 := make(chan string)
	c3 := make(chan bool)
	detectType(c1)
	detectType(c2)
	detectType(c3)
	detectType(3.14)
}
