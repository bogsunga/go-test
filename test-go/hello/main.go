package main

import (
	"github.com/bogsunga/go-test/blob/master/test-go/greet",
	"fmt"
)

func main() {
	message := greet.Greet("David")
	fmt.Println(message)
}
