package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator: function that returns a channel

// https://talks.golang.org/2012/concurrency.slide#27

func boring(msg string) chan string{
  c:= make(chan string)
  go func(){
    for i := 0; ; i++ {
      c <- fmt.Sprintf("%s %d", msg, i)
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
    }()
 return c
}


func main(){
  joe := boring("Joe")
  ann := boring("Ann")
  for i := 0; i < 5; i++ {
      fmt.Println(<-joe)
      fmt.Println(<-ann)
  }
  fmt.Println("You're both boring; I'm leaving.")
}
