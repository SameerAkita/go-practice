package main

import (
	"fmt"
	// "time"
)

// // order struct
// type order struct {
// 	id string
// 	amount float32
// 	status string
// 	createdAt time.Time
// }

// // receiver type
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o order) getAmount() float32 {
// 	return o.amount
// }

// func newOrder(id string, amount float32, status string) *order {
// 	myOrder := order{
// 		id: id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrder
// }

func main() {

	language := struct {
		name string
		isGood bool
	} {"golang", true}
	fmt.Println(language)

	// myOrder := newOrder("1", 30.50, "received")

	// fmt.Println(myOrder)
	// fmt.Println(myOrder.getAmount())
	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder.getAmount())


	// myOrder.createdAt = time.Now()

	// fmt.Println(myOrder.status)

	// myOrder2 := order{
	// 	id: "2",
	// 	amount: 100,
	// 	status: "delivered",
	// 	createdAt: time.Now(),
	// }

	// fmt.Println("Order struct", myOrder)
	// fmt.Println("Order2", myOrder2)
}