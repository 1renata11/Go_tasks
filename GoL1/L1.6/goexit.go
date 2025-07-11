package main

import (
	"fmt"
	"runtime"
	"sync"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Stopped")
	runtime.Goexit()
}

func goexit() {
	wg := &sync.WaitGroup{}
	go worker(wg)
	wg.Add(1)
	wg.Wait()
}
