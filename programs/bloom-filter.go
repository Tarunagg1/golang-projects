package main

import "fmt"

type BloomFilter struct {
	filter []bool
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		make([]bool, size),
	}
}

func (f *BloomFilter) Add(key string) {
	// num := fn(key) % size
}

func (f *BloomFilter) Exists(key string) bool {
	return true
}

func main() {
	bloom := NewBloomFilter(16)

	keys := []string{"a", "b", "c", "d", "e", "f"}

	for _, key := range keys {
		bloom.Add(key)
	}

	for _, key := range keys {
		fmt.Println(key, bloom.Exists(key))
	}
}
