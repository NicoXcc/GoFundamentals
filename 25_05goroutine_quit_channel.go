package main

import (
	"fmt"
	"math/rand"
)

// Generator: function that returns a channel

// https://talks.golang.org/2012/concurrency.slide#27

//We can instead use a fan-in function to let whosoever is ready talk.

func boring(msg string, quit chan bool) <-chan string{ // receive only channel for strings
  c:= make(chan string)
  go func(){
    for i := 0; ; i++ {
      select {
      case c <- fmt.Sprintf("%s: %d", msg, i):
            // do nothing
      case <-quit:
        return
      }
    }
  }()
 return c
}


// We can turn this around and tell Joe to stop when we're tired of listening to him.
func main() {
  quit := make(chan bool)
  c := boring("Joe", quit)

    for i := rand.Intn(10); i >= 0; i-- { fmt.Println(<-c) }
    quit <- true
}
