package main

import "fmt"

func main() {
	// creating map

	m := make(map[string]string)

	m["name"] = "test"
	m["lastname"] = "AGGATWAL"

	fmt.Println(m["ojo"])

	fmt.Println(len(m))
}
