package main

import "fmt"

type Client struct {
	Name string
}

func (c *Client) Greet() {
	fmt.Printf(`{"greeting": "Hello, %s!"}`, c.Name)
	fmt.Printf("\n")
}
