package main

import "fmt"

type OrderStatus int

const (
	Recived OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Chaning order status:", status)
}

func main() {
	changeOrderStatus(Delivered)
}
