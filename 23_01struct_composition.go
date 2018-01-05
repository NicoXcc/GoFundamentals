package main

import "fmt"

// simplified composing structs
type User struct {
	Id             int
	Name, Location string
}

type Player struct {
	User
	GameId int
}

func main() {
	// initializing using dot notation
	p := Player{}
	p.Id = 42
	p.Name = "Matt"
	p.Location = "LA"
	p.GameId = 90404
	fmt.Printf("%+v", p)

	// other option is to use struct literal
	// When using a struct literal with an implicit composition,
	// we can’t just pass the composed fields.
	// We instead need to pass the types composing the struct.
	p1 := Player{
		User{Id: 42, Name: "Mattinson", Location: "LA"},
		90404,
	}
	fmt.Printf("%+v", p1)
}
