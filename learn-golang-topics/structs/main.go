package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	staus     string
	createdAt time.Time
}

func main() {
	order1 := order{
		id:        "1",
		amount:    100.50,
		staus:     "pending",
		createdAt: time.Now(),
	}

	fmt.Println("order", order1)
}
