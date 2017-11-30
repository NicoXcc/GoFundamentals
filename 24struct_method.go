package main

import (
	"fmt"
)

const pi float64 = 3.14159265359

// Go automatically handles conversion between values and pointers
// for method calls. You may want to use a pointer receiver type
// to avoid copying on method calls or to allow the
//  method to mutate the receiving struct.

type circle struct {
	radius int
}

// Methods can be defined for either pointer or value receiver types.
func (ref *circle) area() float64 {
	// with ref we can change the value of struct
	radius := ref.radius
	ref.radius = 2
	return pi * float64(radius) * float64(radius)
}

func (val circle) perimeter() float64 {
	radius := val.radius
	val.radius = 10
	return 2 * pi * float64(radius)
}

func main() {
	firstCircle := circle{radius: 5}

	fmt.Println("Radius is :", firstCircle.radius)

	// area of the Circle
	fmt.Println("Area is :", firstCircle.area())

	// discrepancy in code as the radius is changed to 2
	fmt.Println("Radius erroneously changed to :", firstCircle.radius)

	// perimeter with erroneously changed Radius

	fmt.Println("Perimeter :", firstCircle.perimeter())

	// the perimeter function is pass by value hence
	// it cannot chage the value of parameter defined for struct

	fmt.Println("Radius changed :", firstCircle.radius)

}
