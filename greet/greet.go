package main

import "fmt"

type Greeter interface {
	Greet()
}

func greet(user any) {
	if user, ok := user.(Greeter); ok {
		user.Greet()
	} else {
		fmt.Println("Sorry, not allowed to talk to strangers.")
	}
}
