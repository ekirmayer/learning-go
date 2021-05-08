package main

import (
	"bufio"
	"fmt"
	"learning-go/functions/message" // This is relative to the defined GOPATH
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your greeting? ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Name please? ")
	name, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	m := message.Message(name, phrase)
	fmt.Println(m)
}
