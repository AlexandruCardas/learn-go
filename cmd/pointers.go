package main

import (
	"fmt"
	"strings"
)

// Describer interface decsribes the struct being passed in
// Convention is to name the interface by the job the interface will do
type Describer interface {
	describe() string // Don't forget to specify whats the return type of that method
}

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

// repeatable, not concerned about modifying the state
func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

// Not necessary to use pointer here, if you modify state, use this
func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

// Can have multiple methods with the same name but different receiving parameters
func (g *Group) describe() string {
	if len(g.users) > 2 {
		g.spaceAvailable = false
	}
	desc := fmt.Sprintf("This user group has %d. The newsest user is %s %s. Accepting new users: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

// DoTheDescribing describes whatever struct is passed in
func DoTheDescribing(d Describer) string {
	return d.describe()
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

func main() {
	name := "Elvis"
	changeName(&name)
	fmt.Println(name)
	u := User{ID: 1, FirstName: "Alex", LastName: "Cardas", Email: "calex@gmail.com"}

	u2 := User{ID: 2, FirstName: "Alexandru", LastName: "Cardas", Email: "dlada@~gmail.com"}
	u3 := User{ID: 3, FirstName: "Alexandru", LastName: "Cardas", Email: "dlada@~gmail.com"}
	g := Group{
		role:           "admin",
		users:          []User{u, u2, u3},
		newestUser:     u2,
		spaceAvailable: true,
	}

	fmt.Println(describeUser(u))
	fmt.Println(describeGroup(g))

	desc := u.describe()
	fmt.Println(desc)
	fmt.Println(g.describe())

	userDescriptionWithInterface := DoTheDescribing(&u)
	groupDescriptionWithInterface := DoTheDescribing(&g)

	fmt.Println(userDescriptionWithInterface)
	fmt.Println(groupDescriptionWithInterface)
	pointerMemoryAddress := &c

	updateEmail(&us)

	pointerMemoryAddress.X = 200
	fmt.Println(pointerMemoryAddress)
}
