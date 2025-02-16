package main

import "fmt"

func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	// var items = []int{1, 2, 3, 4, 5}
	names := []string{"tarun", "aggarwal"}

	printSlice(names)
}
