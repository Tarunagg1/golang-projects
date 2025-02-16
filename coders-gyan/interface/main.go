package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type Payment struct {
	gateway paymenter
}

func (p *Payment) makePayment(amount float32) {
	// razorPayModule := razorPay{}
	p.gateway.pay(amount)
}

type razorPay struct{}

func (p *razorPay) pay(amount float32) {
	fmt.Println("Making payment using razor payment", amount)
}

type Stripe struct{}

func (p *Stripe) pay(amount float32) {
	fmt.Println("Making payment using Stripe", amount)
}

func main() {
	razorPayModule := razorPay{}

	newPayment := Payment{
		gateway: &razorPayModule,
	}
	newPayment.makePayment(1000)
}
