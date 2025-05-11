package main

import "fmt"

func changeNumber(a *int) {
	*a = 10
}

func main() {
	num := 1
	changeNumber(&num)
	fmt.Println(num) // Output: 10
}
