package main

import (
	"fmt"
	"reflect"
)

func main() {
	// declaring variables inside func does not need var keyword
	// and : precedes
	// every = sign
	a := 10.0
	b := 5

	fmt.Println("\n A is of type ", reflect.TypeOf(a), " and B is of type ", reflect.TypeOf(b))

  // go does not allow operations on mismatched types
	// c := a + b

	// fmt.Println("Sum is ", c)

  d := int(a) + b
  fmt.Println("Sum is ", d)
}
