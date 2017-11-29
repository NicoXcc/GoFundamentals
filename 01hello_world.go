package main

import (
	"fmt"
	"runtime"
)

// executable program under package main
func main() {
	// idiomatic way of calling a package function
	fmt.Println("Hello from ", runtime.GOOS)
}
