package main

import (
	"fmt"
	"os"
)

// Its idiomatic to return an error as the last return from functions and methods.
// nil is used to indicate success
func testConn(target string) (respTime float64, err error)
func main() {
	_, err := os.Open("/Users/visingh/go/src/pluralsight")
	if err != nil {
		fmt.Println("Error returned was ", err)
	}
}
