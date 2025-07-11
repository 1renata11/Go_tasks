package main

import (
	"fmt"
	"sync"
	"time"
)

var stopped bool

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if stopped {
			fmt.Println("Stopped")
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func stop() {
	wg := &sync.WaitGroup{}
	go worker(wg)
	wg.Add(1)
	time.Sleep(time.Millisecond)
	stopped = true
	wg.Wait()
}
