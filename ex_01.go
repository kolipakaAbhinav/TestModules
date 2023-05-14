package main

import (
	"fmt"
	dog "testmymodules/go/dog"
)

type canine struct {
	name string
	age  int
}

func main() {

	fido := canine{
		name: "test",
		age:  dog.Years(7),
	}
	fmt.Println(fido)
}
