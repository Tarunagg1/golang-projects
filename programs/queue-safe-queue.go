package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Queue struct {
	queue []int32
	mu    sync.Mutex // guards
}

func (q *Queue) Enqueue(item int32) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, item)
}

func (q *Queue) Dqueue() int32 {
	item := q.queue[0]
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = q.queue[1:]
	return item
}

func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.queue)
}

var wg sync.WaitGroup

func main() {
	queue1 := Queue{
		queue: make([]int32, 0),
	}

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			queue1.Enqueue(rand.Int31())
			wg.Done()
		}()
	}

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			queue1.Dqueue()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(queue1.Size())
}
