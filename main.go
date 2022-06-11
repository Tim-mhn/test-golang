package main

import (
	"fmt"
	"net/http"

	"github.com/tim-mhn/module-1/controllers"
)

const (
	first = iota
	two   = iota
)

func startServer() {
	fmt.Println("Starting ...")
}

func log(i int) {
	fmt.Println(i)
}

func sum(i, j int) {
	fmt.Println(i + j)
}

func multiply(i, k int) int {
	return i * k
}

func sumAll(numbers ...int) int {
	var sum = 0
	for _, v := range numbers {
		sum += v
	}

	return sum
}

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
	// const i int = 3
	// fmt.Println(i)
	// const j = float32(i) + 2.21312
	// fmt.Println((j))

	// fmt.Println(first, two)

	// Arrays
	// var arr [3]int = [3]int{1, 2, 3}
	// arr := [3]int{1, 2, 3}
	// slice := arr[:]

	// arr[1] = 10
	// // arr = [1, 2 ,1]
	// slice[0] = 26 // all the elements have the same ref as those in arr

	// slice2 := []int{1, 2, 10}
	// slice3 := append(slice2, 4) // copies all elements, not same ref anymore

	// s4 := slice[1:]
	// s5 := slice3[2:3]

	// slice2[0] = 5
	// fmt.Println(arr)
	// fmt.Println(slice)
	// fmt.Println(slice2)
	// fmt.Println(slice3)
	// fmt.Println(s4)
	// fmt.Println(s5)

	// Maps
	// m := map[string]int{"foo": 4}

	// fmt.Println(m)

	// m["foo"] = 10

	// fmt.Println(m)

	// m["bar"] = 5

	// fmt.Println(m)

	// delete(m, "foo")

	// fmt.Println(m)

	// Structs

	// type user struct {
	// 	name    string
	// 	surname string
	// 	age     int
	// }

	// var u user
	// u.name = "tim"
	// u.age = 24
	// u.surname = "mhn"

	// u := user{age: 24, name: "tim", surname: "mhn"}

	// var u2 user = user{age: 25, name: "tim2", surname: "mhn"}

	// fmt.Println(u)
	// fmt.Println(u2)

	// u := models.User{
	// 	Id:        1,
	// 	FirstName: "tim",
	// 	LastName:  "mhn",
	// }

	// fmt.Println(u)

	// startServer()
	// log(1)
	// sum(2, 3)

	// fmt.Println(multiply(4, 4))

	// numbers := []int{1, 2, 4}
	// fmt.Println(sumAll(numbers...))
	// fmt.Println(sumAll(1, 15))

	// models.AddUser(u)
	// var users = models.GetUsers()
	// fmt.Println(*users[0])

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
