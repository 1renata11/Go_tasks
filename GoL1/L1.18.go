package main

import (
	"fmt"
	"sync"
)

var mutex3 sync.Mutex

type Counter struct {
	counter int
}

func (c *Counter) Add(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex3.Lock()
	c.counter++
	mutex3.Unlock()
}

func L118() {
	wg := &sync.WaitGroup{}
	c := Counter{0}
	for i := 0; i < 10; i++ {
		go c.Add(wg)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Print(c.counter)

}
