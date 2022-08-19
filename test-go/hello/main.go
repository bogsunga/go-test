package main

import (
	"fmt"

	"github.com/bogsunga/go-test/blob/master/test-go/greet"
)

func main() {
	message := greet.Greet("David")
	fmt.Println(message)
}
