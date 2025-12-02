package main

import (
	"fmt"
	// "time"
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

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its a weekend")
	// default:
	// 	fmt.Println("its a weekday")
	// }

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case string:
			fmt.Println("its a string")
		case int: 
			fmt.Println("its an int")
		case bool:
			fmt.Println("its a bool")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI(1.0)
}