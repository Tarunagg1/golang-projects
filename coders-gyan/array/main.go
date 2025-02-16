package main

import "fmt"

func main() {
	// var nums [4]int

	// fmt.Println(len(nums))

	// slices
	// var ages []int
	// fmt.Println(nums == nil)

	// var ages = make([]int, 2)

	var nums = []int{1, 2}
	nums = append(nums, nums...)
	nums = append(nums, 6)

	

	fmt.Println(nums)

}
