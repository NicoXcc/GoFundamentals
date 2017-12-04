package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrency is not parallelism. Concurrency is about dealing with lots of things at once,
// while parallelism is about doing lots of things at once.
// Go does not use threads for its concurrency.
// Actually go has a light weight constructs (go routine) on top of thread
// which achieves concurrency.
// GoRoutine is not scheduled by OS and rather its scheduled by the GoRuntime.

// Go Routine Vs Os threads
// 1. Lighter weight
// 2. Less switching (so while a goRoutine running on a thread is waiting for network I/o to complete ,
//   Go replaces that goRoutine with another routine running while running on the same thread).
//   Switching the routines have expenses but its way less than switching the entire thread.
// 3. Through Channels go routines can easily communicate with each other.
// Go concurrency model is actor model OR Communicating sequential processes.

func main() {
	// anonymous function
	// just ading the go turns a function into go routine

	// here we also need to think of main as a go routine because as soon as
	// main exits the go routines also exits.

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	// what happens: the code spawns a single thread where the first
	// go routine goes into execution.
	go func() {
		// tells the wait group to wait for execution of this function
		defer waitGrp.Done()

		// after executing ths line Go removes this process from the thread and adds
		// 2nd go routine onto the thread
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		// tells the wait group to wait for execution of this function
		defer waitGrp.Done()

		fmt.Println("Pluralsight")
	}()

	// after waiting for two wait groups the program exits
	waitGrp.Wait()
}
