package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// What I found funny is the fact that only area() seens somehow indicate the linkage between a struct and a interface.
// Does that means you need to choose your a interface method very carefully? That it cannot have the same name as the other interface.
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}


func main()  {
	circle := Circle{x:0, y:0, radius:5}

	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}
