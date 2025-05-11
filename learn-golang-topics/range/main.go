package main

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}

	for i, num := range nums {
		println(num, i)
	}

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for k, v := range m {
		println(k, v)
	}
}
