package main

import "fmt"

// sending

// func processNum(numchan chan int) {
// 	for num := range numchan {
// 		fmt.Println("process num: ", num)
// 		time.Sleep(time.Second * 1)
// 	}
// }

func sum(result chan int, num1 int, num2 int) {
	numresult := num1 + num2
	result <- numresult
}

func main() {
	// recive
	// result := make(chan int)

	// go sum(result, 5, 40)

	// res := <-result

	// fmt.Println(res)

	// sending

	// numchan := make(chan int)
	// go processNum(numchan)

	// for {
	// 	numchan <- rand.Intn(100)
	// }

}
