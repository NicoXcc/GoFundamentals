package main

import (
	"fmt"
)

// main function doesnot take any arguments
// it does not return any value
func main() {
  bestFinish := bestLeagueFinishes(13,10,13,14, 16,17,19)
  fmt.Println(bestFinish)
}

// ...int signifies that 'finishes' will take slices of int
func bestLeagueFinishes(first int, finishes ...int) int {
  best := finishes[0]
  for _, i := range finishes {
    if i < best {
      best = i
    }
  }
  fmt.Println(first)
  return best
}
