package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p *payment) makePayment(amount float32) {

	p.gateway.pay(amount)

	// stripePayment := stripe{}
	// stripePayment.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Razorpay payment of amount:", amount)
}

type stripe struct{}

func (r stripe) pay(amount float32) {
	fmt.Println("Stripe payment of amount:", amount)
}

func main() {
	stripePaymentGW := stripe{}

	newPayment := payment{
		gateway: stripePaymentGW,
	}

	newPayment.makePayment(1000)
}
