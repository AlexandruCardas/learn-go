package main

import (
	"fmt"
)

// Instead of passing in the struct, you are calling the method with an instance of that struct
//func (u *User) describe() string {
//	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
//	return desc
//}

func main() {
	user := User{ID: 1, FirstName: "Alex", LastName: "Cardas", Email: "calexc@gmail.com"}

	// desc := describeUser(user) // function
	desc := user.describe() // method

	fmt.Println(desc)
}
