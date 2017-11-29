package main

import (
	"fmt"
)

// Go passes arguments by value and not by reference.
// Hence a primitive value passed can be used and changed
//  without changing the actual variable passed as
// argument to the function.

func main() {
	a := 10
	ptr := &a // holds the memory address of a
	// *ptr dereference the address and gives the actual value of the variable
	fmt.Println("memory address of variable a ", ptr, "while the actual value is ", *ptr)
}
