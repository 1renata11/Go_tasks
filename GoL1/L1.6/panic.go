package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Горутина завершилась через panic и recover:", r)
		}
	}()
	panic("Stopped")
}

func main() {
	wg := &sync.WaitGroup{}
	go worker(wg)
	wg.Add(1)
	wg.Wait()
}
