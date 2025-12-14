package main

import (
	"fmt"
	"maps"
)

func main() {
	// m := make(map[string]string)
	// m["name"] = "golang"
	// m["area"] = "backend"

	// fmt.Println(m["name"], m["area"])

	// m := make(map[string]int)
	// m["age"] = 20
	// m["price"] = 50
	// // fmt.Println(m["age"])
	// // fmt.Println(len(m))
	// clear(m)
	// fmt.Println(m)

	// v, ok := m["price"]
	// fmt.Println(v)
	// fmt.Println(ok)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }
	m1 := map[string]int{"price": 40, "phone": 3}
	m2 := map[string]int{"price": 40, "phone": 3}

	fmt.Println(maps.Equal(m1, m2))
}