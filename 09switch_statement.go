package main

import (
	"fmt"
)

// Default statement can go anywhere and not necessarily in the end.

// No need of break statement as it does not have default fall-through mech.
// Each Case statement in Go has gotten implicit break.
// If we still want fall through behavior the we need to write “fallthrough”
// as last line of case.

func main() {
	switch "docker" {
	case "docker":
		{
			fmt.Println("first print")
      fallthrough
		}
	case "docker1":
		{
			fmt.Println("second print")
		}
	default:
		fmt.Println("default print")
	}
}
