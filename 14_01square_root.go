package main

import "fmt"

// 1 Start with an arbitrary positive start value x (the closer to the
//    root, the better).
// 2 Initialize y = 1.
// 3. Do following until desired approximation is achieved.
//   a) Get the next approximation for root using average of x and y
//   b) Set y = n/x

func Sqrt(num float64) float64 {
	// putting x as n
	y := 1.0
	root := num
	appx := 0.0001
	for root-y > appx {
		root = (root + y) / 2.0
		y = num / root
	}
	return root
}

func main() {
	fmt.Println(Sqrt(30.0))
}
