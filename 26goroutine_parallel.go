package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// If we want the go routine to run in parallel
// then we need to import the runtime package
// and use inbuilt GOMAXPROCS to up the number of virtual
// processes to two.

func main() {
	runtime.GOMAXPROCS(2)
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {

		defer waitGrp.Done()

		fmt.Println("Pluralsight")
	}()

	waitGrp.Wait()
}
