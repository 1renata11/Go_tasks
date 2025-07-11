package main

import (
	"fmt"
	"sync"
	"time"
)

func sleep(duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	for {
		if time.Since(start) >= duration {
			break
		}
	}
}

func L125() {
	var wg sync.WaitGroup
	fmt.Println("Start:", time.Now())
	wg.Add(1)
	go sleep(time.Second, &wg)
	fmt.Println("End:", time.Now())
	wg.Wait()
}
