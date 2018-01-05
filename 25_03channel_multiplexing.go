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

// func fanIn(input1, input2 <-chan string) chan string{
//   c:= make(chan string) // channel of channels
//   go func ()  {
//     for {
//       c <- <-input1
//     }
//   }()
//   go func ()  {
//     for {
//       c <- <-input2
//     }
//   }()
//
//   return c
// }


// The select statement provides another way to handle multiple channels.
// It's like a switch, but each case is a communication:
// - All channels are evaluated.
// - Selection blocks until one communication can proceed, which then does.
// - If multiple can proceed, select chooses pseudo-randomly.
// - A default clause, if present, executes immediately if no channel is ready.
//     select {
//     case v1 := <-c1:
//         fmt.Printf("received %v from c1\n", v1)
//     case v2 := <-c2:
//         fmt.Printf("received %v from c2\n", v1)
//     case c3 <- 23:
//         fmt.Printf("sent %v to c3\n", 23)
//     default:
//         fmt.Printf("no one was ready to communicate\n")
//     }



func fanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() {
        for {
            select {
            case s := <-input1:  c <- s
            case s := <-input2:  c <- s
            }
        }
    }()
    return c
}


func main(){
  fifs_channel := fanIn(boring("Joe"),boring("Ann"))

  for i := 0; i < 10; i++ {
      fmt.Println(<-fifs_channel)
  }
  fmt.Println("You're both boring; I'm leaving.")
}
