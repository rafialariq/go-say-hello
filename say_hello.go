package go_say_hello

import "fmt"

func SayHello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}