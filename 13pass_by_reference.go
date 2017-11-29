package main

import (
	"fmt"

)

func main() {
  course:= "Go fundamentals "

	fmt.Println("\n Course is ", course)

  // pass address of variable
  changeCourse(&course)

  fmt.Println("\n Course is ", course)
}

// takes a ponter reference to the address passed
func changeCourse (course *string) string {
  *course = "a New Go Fundamentals"
   // just the equal operator as we are not declaring variable here
  fmt.Println("\n changed course is ", *course)
  return *course
}
