package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("other")
	// }

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its a weekend")
	default:
		fmt.Println("its a weekday")
	}
}