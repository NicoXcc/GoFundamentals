package main

import (
	"fmt"
)

func main() {
	// to define a struct
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	// instantiating with 0 values
	// Go has a habit of initializing with zero values

	var course courseMeta

	fmt.Println(course)

	// instantiating with new keyword
	// gives a pointer and output will be &{  0}
	newCourse := new(courseMeta)
	fmt.Println(newCourse)

	// instantiating the literal way
	anotherCourse := courseMeta{
		Author: "vikash",
		Level:  "beginner",
		Rating: 0.0,
	}

	fmt.Println(anotherCourse)

}
