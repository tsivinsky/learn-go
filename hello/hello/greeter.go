package hello

type Greeter interface {
	Greet()
}

func Greet(user Greeter) {
	user.Greet()
}
