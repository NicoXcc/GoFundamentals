package main

import (
	"fmt"
)

// Maps are unordered list. We cannot expect maps to return the values in the same sequence as we have added it.
// 1. Maps <key>:<value> pair
// 2. Maps are dynamically resizable
// 3. Maps are references  (pass by reference)
// Since Go is strongly typed, we need to declare type for both key and value
// Map[keyType]valueType
//
// The “keyType” and “valueType” are both comparable types.
// A type that can be comapared with the “==“ or “!="

// SPECIFY SIZE FOR LARGE MAPS, IT CAN IMPROVE YOUR CODES PERFORMANCE

// make(map[<key type>]<value type>, size) // specify size if your map can get very big
// at the back of the hood, adding new element to maps does the same thing as it happens
// for arrays

func main() {
	// declaring a map
	courseMarks := make(map[string]int)
	courseMarks["maths"] = 90
	courseMarks["hindi"] = 70
	courseMarks["english"] = 85

	fmt.Println(courseMarks)

	// declaring a map literal form
	previousCourseMarks := map[string]int{
		"maths":   99,
		"hindi":   80,
		"english": 75}

	fmt.Println(previousCourseMarks)
}
