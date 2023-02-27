package hello

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func SayHello(name string) {
	fmt.Printf("%s\n", Hello(name))
}
