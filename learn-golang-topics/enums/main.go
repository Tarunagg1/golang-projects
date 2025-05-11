package main

import "fmt"


type orderStatus int


const (
	Revived orderStatus = 10
	Confirmed = 11
	Prepared
	Shipped
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("Order status changed to: ", status)
}

func main() {
	changeOrderStatus(Revived)
	changeOrderStatus(Confirmed)
}
