package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped")
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func contextt() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	go worker(wg, ctx)
	wg.Add(1)
	time.Sleep(time.Millisecond)
	cancel()
	wg.Wait()
}
