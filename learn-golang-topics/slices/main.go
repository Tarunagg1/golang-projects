package main

import "fmt"

func main() {
	var nums []int

	var nums1 = make([]int, 2)

	nums1 = append(nums1, 1)
	nums1 = append(nums1, 2)

	fmt.Println("nums:", nums) // nil slice
	fmt.Println(len(nums))
	fmt.Println(cap(nums1))
	fmt.Println(nums1)

	

}
