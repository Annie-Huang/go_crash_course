package main

import "fmt"

//var name = "Brad"


func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr		<< unsign int, meaning only positive number, no negative.
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	//var name string = "Brad"	// Don't need to explicit the type as it's already inferred
	//var name = "Brad"
	//var age int = 37	// Same as above
	var age int32 = 37
	const isCool = true
	var size float32 = 2.3

	// Shorthand
	//name := "Brad"
	//size := 1.3	// default type is float64
	//email := "brad@gmail.com"

	// Assigned multiple values in one line
	name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}


















