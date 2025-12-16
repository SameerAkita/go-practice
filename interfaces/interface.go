package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
}

func (p payment) refundPayment(amount float32) {
	p.gateway.refund(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

func (r razorpay) refund(amount float32) {
	fmt.Println("refund payment from razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

type paypal struct {}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32) {
	fmt.Println("refund payment from paypal", amount)
}

func main() {
	razorpayPaymentGw := razorpay{}
	newPayment := payment{gateway: razorpayPaymentGw}
	newPayment.makePayment(100)
	newPayment.refundPayment(100)

	paypalPaymentGw := paypal{}
	newPayment2 := payment{gateway: paypalPaymentGw}
	newPayment2.makePayment(75)
	newPayment2.refundPayment(75)
}
