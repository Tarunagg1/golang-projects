package main

import "fmt"

func main() {
	// map
	m := make(map[string]string)

	m2 := map[string]int{"price": 10, "quantity": 5}
	fmt.Println(m2)

	m["name"] = "Tarun"
	m["age"] = "23"

	fmt.Println(m["name"])

	

}
