package main

import (
	"fmt"
)

// append operation on slice doubles the size of the underlying array
// if the length exceeds capacity.
// Doubling the size of the array is just creating a new array with double the Size
// and adding all the elements from the previous array to the new array and referencing
// it through slice
func main() {
	autosize := []int{1, 2, 3}
	fmt.Println(autosize)
	fmt.Println("Slice ", autosize, " capacity is : ", cap(autosize))
	autosize = append(autosize, 5, 6, 7)
	fmt.Println("Append [5,6,7] the slice becomes ", autosize, " and capacity gets doubled : ", cap(autosize))
	autosize = append(autosize, 8)
	fmt.Println("Append 8, the slice becomes ", autosize, " but capacity gets doubled again : ", cap(autosize))
}
