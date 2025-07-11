package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, done <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Println("Stopped")
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func chann() {
	wg := &sync.WaitGroup{}
	done := make(chan struct{})
	go worker(wg, done)
	wg.Add(1)
	time.Sleep(time.Millisecond)
	close(done)
	wg.Wait()
}
