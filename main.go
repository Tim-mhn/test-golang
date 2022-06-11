package main

import "fmt"

const (
	first = iota
	two   = iota
)

func main() {

	// Variable Declarations
	// var i = 2
	// var f float64 = 3
	// firstName := "tim"
	// firstName = "timmm"
	// b := true

	// name, surname := "tim", "mnh"
	// fmt.Println(name, surname)
	// fmt.Println("hello Go from a module !", i, f, firstName)

	// var firstName *string = new(string)
	// *firstName = "Arthur"
	// fmt.Println(*firstName)

	// Pointers
	// firstName := "Arthur"
	// fmt.Println(firstName)

	// ptr := &firstName
	// fmt.Println(ptr, *ptr)

	// firstName = "Tim"
	// fmt.Println(firstName)
	// fmt.Println(*ptr)

	// Constant declarations
	const i int = 3
	fmt.Println(i)
	const j = float32(i) + 2.21312
	fmt.Println((j))

	fmt.Println(first, two)
}
