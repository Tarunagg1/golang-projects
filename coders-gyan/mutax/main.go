package main

import (
	"fmt"
	"sync"
)

type Post struct {
	views int
	mu    sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	defer p.mu.Unlock()

	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := Post{
		views: 0,
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println(myPost.views)
}
