package main

import (
	"fmt"
)

func main() {
	testMap := map[string]int{"A": 10, "B": 20, "C": 30, "D": 40}
	// add a new values
	testMap["E"] = 50

	fmt.Println(testMap)
	// delete a key
	delete(testMap, "A")

	fmt.Println(testMap)
}
