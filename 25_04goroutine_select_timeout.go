package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator: function that returns a channel

// https://talks.golang.org/2012/concurrency.slide#27

//We can instead use a fan-in function to let whosoever is ready talk.

func boring(msg string) <-chan string{ // receive only channel for strings
  c:= make(chan string)
  go func(){
    for i := 0; ; i++ {
      c <- fmt.Sprintf("%s %d", msg, i)
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
    }()
 return c
}

// In the previous program, we had a timeout for each message
// func main() {
//     c := boring("Joe")
//     for {
//         select {
//         case s := <-c:
//             fmt.Println(s)
//         case <-time.After(1 * time.Millisecond):
//             fmt.Println("You're too slow.")
//             return
//         }
//     }
// }

// Create the timer once, outside the loop, to time out the entire conversation
func main() {
    c := boring("Joe")
    timeout := time.After(5 * time.Second)
    for {
        select {
        case s := <-c:
            fmt.Println(s)
        case <-timeout:
            fmt.Println("You talk too much.")
            return
        }
    }
}
