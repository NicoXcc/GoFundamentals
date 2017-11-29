package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}
	anotherSlice := []int{4, 5, 6}
	slice = append(slice, anotherSlice...)
	fmt.Println("appended slice :", slice)
}
