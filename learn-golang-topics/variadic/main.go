package main

func sum(num ...int) int {
	total := 0
	for _, n := range num {
		total += n
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	println(result)
}
