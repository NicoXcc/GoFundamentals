package main

import "fmt"

func main() {
	// Note how we used the fmt package using
	// Printf and used the %q “verb” to print each element quoted.

	// set the array entries as you declare
	a := [2]string{"hello", "world!"}
	fmt.Printf("%s\n", a)
	// use an ellipsis to use an implicit length when you pass the values
	a1 := [...]string{"hello", "world!"}
	fmt.Printf("%q\n", a1)
}
