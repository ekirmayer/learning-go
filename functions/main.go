package main

import (
	"fmt"
	"learning-go/functions/message" // This is relative to the defined GOPATH
)

func main() {
	m := message.Message("You", "Hello")
	fmt.Println(m)

	r := message.New_Message("New", "whatsup")
	fmt.Println(r)

	vr, i := message.Very_New_Message("This", "Test")
	fmt.Println(i, vr)
}
