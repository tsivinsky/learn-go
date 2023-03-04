package main

import "fmt"

type User struct {
	Name string
}

func modifyUserWithPointer(u *User) {
	u.Name = "Jack"
}

func modifyUser(u User) {
	u.Name = "Cisco"
}

func main() {
	u := &User{Name: "Dan"}

	fmt.Printf("u.Name: %v\n", u.Name)

	modifyUser(*u)

	fmt.Printf("u.Name: %v\n", u.Name)

	modifyUserWithPointer(u)

	fmt.Printf("u.Name: %v\n", u.Name)
}
