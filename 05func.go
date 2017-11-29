package main

import (
	"fmt"
	"strings"
)

// main function doesnot take any arguments
// it does not return any value
func main() {
	module := "functions basics"
	author := "vikash singh"

	fmt.Println(converter(module, author))
}

// in Go function can return multiple values
func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}
