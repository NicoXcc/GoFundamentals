package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Goroutines
// What is a goroutine? It's an independently executing function, launched by a go statement.
// It has its own call stack, which grows and shrinks as required.
// It's very cheap. It's practical to have thousands, even hundreds of thousands of goroutines.
// It's not a thread.
// There might be only one thread in a program with thousands of goroutines.
// Instead, goroutines are multiplexed dynamically onto threads as needed to keep all the goroutines running.
// But if you think of it as a very cheap thread, you won't be far off.

// A channel in Go provides a connection between two goroutines, allowing them to communicate.
//     Declaring and initializing.
//     var c chan int
//     c = make(chan int)
//     or
//     c := make(chan int)
//     Sending on a channel.
//     c <- 1
//     Receiving from a channel.
//     The "arrow" indicates the direction of data flow.
//     value = <-c

// HOW CHANNELS communicate

// When the main function executes <–c, it will wait for a value to be sent.
// Similarly, when the boring function executes c <– value, it waits for a receiver to be ready.
// A sender and receiver must both be ready to play their part in the communication.
// Otherwise we wait until they are.
// Thus channels both communicate and synchronize.

// Don't communicate by sharing memory, share memory by communicating.

func boring(msg string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Exiting!")
}
