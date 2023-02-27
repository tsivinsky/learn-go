package main

import (
	"fmt"
	"hello"
)

type Guest struct {
	Name string
}

func (g *Guest) Greet() {
	hello.SayHello(g.Name)
}

type Admin struct {
	Name string
}

func (a *Admin) Greet() {
	hello.SayHello(fmt.Sprintf("%s [admin]", a.Name))
}

func main() {
	guest1 := &Guest{Name: "Kate"}
	guest2 := &Guest{Name: "Jack"}

	admin := &Admin{Name: "Dan"}

	hello.Greet(guest1)
	hello.Greet(guest2)
	hello.Greet(admin)
}
