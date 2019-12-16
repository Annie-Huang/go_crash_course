package main

import "fmt"

func main()  {
	// The reason you will use pointer is that e.g. if your data stored in the memory is huge
	// It will be better for performance to just pass the memory address in rather than data value

	a := 5
	b := &a	// b is the pointer memory address

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)	// int
	fmt.Printf("%T\n", b)	// *int

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
