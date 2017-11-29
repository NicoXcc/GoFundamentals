package main

import (
	"fmt"
)

func main() {
	array := []int{5, 4, 6}
	fmt.Println(len(array)) // prints 5 because all other empty spaces are filled with 0
	fmt.Println(cap(array)) // prints 5

	passByRef(array)
}

func passByRef(ptr []int) {
	fmt.Println(ptr)
}
