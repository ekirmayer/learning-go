package main

import "fmt"

func main() {
	// define vars
	// var <name> <type> [= <value>]
	var myInt int = 19

	// Define multiple vars at the same time
	// The value will be set automatically
	var val, ok = "yes", true

	// Define multiple vars at the same time when type is known
	var one, two, three int = 1, 2, 3

	fmt.Println("myInt is:", myInt)
	fmt.Println("myInt times two:", myInt*2)
	fmt.Println("val is:", val)
	fmt.Println("ok is:", ok)
	fmt.Println(one, two, three)

	// Define var and assign later
	var name string
	name = "MyName"

	fmt.Println("My name is:", name)

	// A short version to define and assign
	number := 256
	str := "a new string"
	fmt.Println("This is the new number", number)
	fmt.Println("This is the new string", str)
}
