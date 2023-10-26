package gotest

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome to Golang!!!", name)
	return message
}

