package main

import (
	"fmt"
	"time"
)

func processNum(numchan chan int) {
	for num := range numchan {
		fmt.Println("Processing number...", num)
		time.Sleep(time.Second * 1)
	}
}

func sum(result chan int, num1 int, num2 int) {
	result <- num1 + num2
}

func emailSender(emailChan chan string, done chan bool) {
	defer func() {
		done <- true // Signal that the email sender is ready
	}()

	for email := range emailChan {
		fmt.Println("Sending email to:", email)
		time.Sleep(2 * time.Second) // Simulate time taken to send an email
	}
}

func main() {
	// messageChan := make(chan string)

	// messageChan <- "Hello, World!"

	// fmt.Println(<-messageChan)

	// EXM 2 sender
	// numchan := make(chan int)

	// go processNum(numchan)

	// for {
	// 	numchan <- rand.Intn(100)
	// }

	// time.Sleep(2 * time.Second)

	// reciving
	// result := make(chan int)
	// go sum(result, 4, 5)

	// res := <-result

	// fmt.Println("Result:", res)

	// buffer channel

	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%dgmail.com", i)
	}

	fmt.Println("Email sending done...")

	close(emailChan) // Close the channel to signal completion

	<-done

}
