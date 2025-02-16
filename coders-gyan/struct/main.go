package main

import (
	"fmt"
	"time"
)

type Customer struct {
	name  string
	phone string
}

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	Customer
}

func (o *Order) changeStatus(status string) {
	o.status = status
}

func newOrder(id string, amount float32, status string) *Order {
	return &Order{
		id:       id,
		amount:   amount,
		status:   status,
		Customer: Customer{"ifj", "484"},
	}

}

func main() {
	// order := Order{
	// 	id:     "123",
	// 	amount: 1.0,
	// 	status: "success",
	// }

	order := newOrder("1", 10.5, "initial")

	order.createdAt = time.Now()

	order.changeStatus("hahha")

	fmt.Println(order)
}
