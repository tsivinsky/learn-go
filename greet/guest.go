package main

import "fmt"

type Guest struct {
	Name string
}

func (g *Guest) Greet() {
	fmt.Printf("Hello, %s!\n", g.Name)
}
