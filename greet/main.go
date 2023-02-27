package main

func main() {
	guest1 := &Guest{Name: "Kate"}
	guest2 := &Guest{Name: "Jack"}

	admin := &Admin{Name: "Daniil"}

	client1 := &Client{Name: "frontend"}
	client2 := &Client{Name: "service 2"}

	greet(guest1)
	greet(admin)
	greet(guest2)
	greet(client1)
	greet(client2)
	greet("kek")
}
