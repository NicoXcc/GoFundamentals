package main

import (
	"fmt"
	"reflect"
)

// by default gets initialized with empty string
// default gets initialized with 0.0
var (
	name, course string
	module       float64
)

var (
	_name, _course, _module = "Vikash", "Go fundamentals", 3.2
)

func main() {
	fmt.Println("Name is ", name, "and is if type ", reflect.TypeOf(name))
	fmt.Println("Module is ", module, " and of type ", reflect.TypeOf(module))
	fmt.Println("Name is ", _name, "and is if type ", reflect.TypeOf(_name))
	fmt.Println("Module is ", _module, " and of type ", reflect.TypeOf(_module))
}
