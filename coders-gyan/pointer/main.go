package main

import "fmt"

func changeNumber(num *int) {
	fmt.Println("In change nunber", *num)
	*num = 5
}

func main() {
	num := 1
	changeNumber(&num)

	fmt.Println("After change number", num)
}
