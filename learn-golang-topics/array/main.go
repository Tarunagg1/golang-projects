package main

import "fmt"

func main() {
	var num [4]int

	fmt.Println(len(num))
	fmt.Println(num)

	num[0] = 1
	num[1] = 2
	num[2] = 3
	num[3] = 4

	fmt.Println(num)

	

}
