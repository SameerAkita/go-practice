package main

import "fmt"

func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// type stack[T any] struct {
// 	elements []T
// }
 
func main() {
	// nums := []int{1, 2, 3}
	// names := []string{"golang", "typescript"}
	vals := []bool{false, true, true}


	printSlice(vals, "john")
	// myStack := stack[string]{
	// 	elements: []string{"golang", "typescript"},
	// }

	// fmt.Println(myStack)
}