package main

import (
	"fmt"

	"base/greetings"
)

// https://go.dev/doc/tutorial/greetings-multiple-people
// https://go.dev/doc/tutorial/compile-install
func main() {

	msg, err := greetings.Hello("James")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)

}
