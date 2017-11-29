package main

import (
	"fmt"
)

// Slices have variable length.
// Caveat: Slices are built on top of an array.
// Slices are references to contiguous positions in the array(pass by reference),
// there are no direct way to save data in
// Size of the slice cannot be bigger than the array backing it
//
// myCourse := make([]string, 5, 10) // make(<type>, <length>, <capacity>)
// []string —> signifies a slice
// Len(myCourse) —> gives length
// Cap(myCourse) —> gives capacity
//
// SliceOfSlice := myCourse[0:5]  // from position 0 to position 4
// [0:] // 0 to last position
// [:5] // 0 to position 4
//
// Slices are by default pass by reference, so need of defining pointers

func main() {
	myCourse := make([]string, 5, 10)
	fmt.Printf("Length id %d\nCapacity is %d \n", len(myCourse), cap(myCourse))

	// initializes with 0 OR empty string
	fmt.Println("1st element is : ", myCourse[0])

	anotherSlice := make([]string, 5) // by default capacity is 5

	yetAnotherSlice := []string{"maths", "science"} // length 2 capacity 2

	sliceOfSlice := yetAnotherSlice[0:1]

	fmt.Println(anotherSlice)
	fmt.Println(yetAnotherSlice)
	fmt.Println(sliceOfSlice)
}
