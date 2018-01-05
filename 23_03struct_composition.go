package main

import "fmt"

// Looking at the User / Player example, you might
// have noticed that we composed Player using User but it might be better
// to compose it with a pointer to a User struct. The reason why a pointer
// might be better is because in Go, arguments are passed by value and not reference.
// If you have a small struct that is inexpensive to copy, that is fine,
// but more than likely, in real life, our User struct will be bigger and should not be copied.
// Instead we would want to pass by reference (using a pointer).

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

type Player struct {
	*User
	GameId int
}

func NewPlayer(id int, name, location string, gameId int) *Player {
	return &Player{
		User:   &User{id, name, location},
		GameId: gameId,
	}
}

func main() {
	p := NewPlayer(1, "vikash", "blr", 1)

	fmt.Printf("%+v", p)

}
