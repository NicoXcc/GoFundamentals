package main

import (
	"log"
	"os"
)

type Job struct {
	Command string
	Logger  *log.Logger
}

// implementing a Job struct that can also behave like a Logger

func main() {
	job := &Job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
	// same as
	// job := &Job{Command: "demo",
	//            Logger: log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Logger.Print("test")
}

// Our Job struct has a field called Logger which is a pointer to another type (log.Logger)

// When we initialize our value, we set the logger so we can then call its
// Print function by chaining the calls: job.Logger.Print()

// But Go lets you go even further and use implicit composition.

// We can skip defining the field for our logger and now all the methods
// available on a pointer to log.Logger are available from our struct

// package main
//
// import (
// 	"log"
// 	"os"
// )
//
// type Job struct {
// 	Command string
// 	*log.Logger
// }
//
// func main() {
// 	job := &Job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
// 	job.Print("starting now...")
// }
