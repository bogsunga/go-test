package main

import (
	"https://github.com/bogsunga/go-test/test-go/greet"
	"fmt"
)

func main() {
	message := greet.Greet("David")
	fmt.Println(message)
}
