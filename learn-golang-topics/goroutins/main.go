package main

import (
	"sync"
)

func tasks(id int, w *sync.WaitGroup) {
	defer w.Done() // Decrement the WaitGroup counter when the goroutine completes
	println("Task", id, "is running")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		go tasks(i, &wg)
		wg.Add(1)
	}
	
	wg.Wait()
}
