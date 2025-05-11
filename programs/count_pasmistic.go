package main

import (
	"fmt"
	"sync"
)

var count int = 0

var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	count++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("hello world", count)
}
