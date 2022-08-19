package greet

import "fmt"

func Greet(name string) string {
	message := fmt.Sprintf("Hello %s", name)
	return message
}
