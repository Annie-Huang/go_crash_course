package main

import "fmt"

// A closure is the combination of a function bundled together (enclosed) with references to its surrounding state (the lexical environment).
// In other words, a closure gives you access to an outer function’s scope from an inner function.
// In JavaScript, closures are created every time a function is created, at function creation time.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main()  {
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
