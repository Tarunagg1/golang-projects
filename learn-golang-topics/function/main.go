package main

func add(a int, b int) (error, int) {
	return nil, a + b
}

// func parseInt(fn func(int) int) {

// }




func main() {
	_, ans := add(1, 2)
	println(ans)
}
