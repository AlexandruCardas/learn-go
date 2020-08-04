package main

import (
	"fmt"
	"strings"
)

// User is a user type
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

// A group of users
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

// func describeGroup
// This user group has 19 users. The newest user is Joe Smith. Accepting new users-true.
func describeGroup(g Group) string {
	if len(g.users) > 2 {
		g.spaceAvailable = false
	}
	desc := fmt.Sprintf("This user group has %d. The newsest user is %s %s. Accepting new users: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

func changeName(n *string) {
	*n = strings.ToUpper(*n)
}

type Coordinates struct {
	X, Y float64
}

type UserSimple struct {
	user  string
	email string
}

var c = Coordinates{X: 10, Y: 20}
var us = UserSimple{"alex", "email"}

func updateEmail(u *UserSimple) {
	u.email = "some new email"
	fmt.Println("I updated the email", u.email)
}

//func main() {
//	name := "Elvis"
//	changeName(&name)
//	fmt.Println(name)
//	u := User{ID: 1, FirstName: "Alex", LastName: "Cardas", Email: "calex@gmail.com"}
//
//	u2 := User{ID: 2, FirstName: "Alexandru", LastName: "Cardas", Email: "dlada@~gmail.com"}
//	u3 := User{ID: 3, FirstName: "Alexandru", LastName: "Cardas", Email: "dlada@~gmail.com"}
//	g := Group{
//		role:           "admin",
//		users:          []User{u, u2, u3},
//		newestUser:     u2,
//		spaceAvailable: true,
//	}
//
//	fmt.Println(describeUser(u))
//	fmt.Println(describeGroup(g))
//
//	pointerMemoryAddress := &c
//
//	updateEmail(&us)
//
//	pointerMemoryAddress.X = 200
//	fmt.Println(pointerMemoryAddress)
//}
