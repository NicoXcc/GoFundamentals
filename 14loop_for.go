package main

import _ "net/http/pprof"

import (
	"fmt"
	"time"
	"net/http"
	"log"
)

// For <expression>
//
// <expression>
// 1. Can be blank (infinite loop)
// 2. Can be a proper boolean expression
// 3. Can be a range

// For i < 10 { // is a boolean expression. Is a while loop in disguise
// }
//
// It can have simple pre and post statements
//
// For i:= 0; i<10; i++ {
//     <code>
// }
//
// It can be a range, course list can be a slice Or it can be a map
// for i:= range courseList {
// }

// break command breaks out of the loop
// break <label> breaks the loop and takes us to the label

func main() {
	// being idiomatic and initialized timer inside for loop
	for timer := 3; timer > 0; timer-- {
		fmt.Println("time left :", timer)
		time.Sleep(1 * time.Second)
	}

	// declaring a slices

	course := []string{"math", "english", "hindi"}

	// range returns both index and value
	// if we donot want the index then we ignore it with blank '_' identifier
	for _, i := range course {
		fmt.Println(i)
	}

	for true {
		fmt.Println("conditional looping")
		break
	}

	for {
		fmt.Println("loop forever")
		break
	}
}
