package main

import (
	"fmt"
  "os"
)

func main() {
	fmt.Println("\n Environment variables ", os.Environ())

  for _, env:= range os.Environ() {
    fmt.Println(env)
  }

  name:= os.Getenv("USER")
  fmt.Println("\n User Name ", name)
}
