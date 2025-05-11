package main

import "fmt"

func printSclice[T any](data []T) {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	printSclice(nums)

}
