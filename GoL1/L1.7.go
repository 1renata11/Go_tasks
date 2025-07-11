package main

import (
	"log"
	"math/rand"
	"sync"
)

var mp map[int]int
var mutex2 sync.Mutex

func Write(key int, value int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex2.Lock()
	mp[key] = value
	mutex2.Unlock()
}

func L17() {
	var wg sync.WaitGroup
	mp = make(map[int]int)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Write(rand.Int(), rand.Int(), &wg)
	}
	wg.Wait()
	for key, value := range mp {
		log.Printf("m[%v] = %v ", key, value)
	}
}
