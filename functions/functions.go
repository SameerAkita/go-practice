package main

import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, string) {
// 	return "golang", "javascript", "c"
// }

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return a
	}
}

func main() {

	// fn := func(a int) int {
	// 	return 2
	// }

	fn := processIt()
	fmt.Println(fn(6))

	// result := add(3, 5)
	// fmt.Println(result)

	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3)
}