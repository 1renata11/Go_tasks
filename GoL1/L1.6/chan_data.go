package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for val := range ch {
		fmt.Println("Получено:", val)
	}
	fmt.Println("Closed")
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	go worker(wg, ch)
	wg.Add(1)
	ch <- 1
	close(ch)
	wg.Wait()
}
