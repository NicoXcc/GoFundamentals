package main

import (
	"fmt"
)

func main() {
	testMap := map[string]int{"A": 10, "B": 20, "C": 30, "D": 40}

  for key, value := range testMap {
    fmt.Printf("\n key is : %v and value is : %v", key, value)
  }

  fmt.Println("\n",len(testMap))
}
