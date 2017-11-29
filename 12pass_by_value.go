package main

import (
	"fmt"

)

func main() {
  course:= "Go fundamentals "

	fmt.Println("\n Course is ", course)
  changeCourse(course)
}

func changeCourse (course string) string {
  course = "a New " + course
   // just the equal operator as we are not declaring variable here
  fmt.Println("\n changed course is ", course)
  return course
}
