package main

import (
	"fmt"
)

// In Go conditionals are only on Boolean true or false, and it is not evaluated
// based on integer or string hoping Go will do the type conversion.
//
// if <Boolean expression> {
//
// If <Boolean expression>
// { // results in compile time error
//
// If   <Boolean expression> {
//     // code block
// } else if <Boolean expression> {
//        // code block
// } else {
// }
//
// If also supports optional initialization statements which are scoped in the if block only.
//
// If <simple statements> ; <Boolean expression> {
// }

func main() {

	if firstRank, secondRank := 39, 65; firstRank < secondRank {
		fmt.Println("first course is doing better than the second course")
	} else if firstRank > secondRank {
		fmt.Println("second course is doing better")
	} else {
		fmt.Println("dono same hai")
	}
}
