package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for _, num := range nums {
		fmt.Println(num)
	}

	myMap := map[string]string{"name": "test", "age": "123", "name2": "test2"}

	for k, v := range myMap {
		fmt.Println(k, v)
	}

}
