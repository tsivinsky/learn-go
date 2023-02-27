package main

import (
	"fmt"
	"log"
	"os"
)

type Admin struct {
	Name string
}

func (a *Admin) Greet() {
	f, err := os.OpenFile("greeting.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "Hello, %s!\n", a.Name)
}
