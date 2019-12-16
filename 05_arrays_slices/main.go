package main

import "fmt"

func main()  {
	// Arrays
	//var fruitArr [2]string	// 2 is the size

	// Assign values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	//// Declare and assign
	//fruitArr := [2]string{"Apple", "Orange"}
	//
	//fmt.Println(fruitArr)
	//fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}	// If you give a size of 2, it will throw array index 2 out of bounds [0:2]

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
