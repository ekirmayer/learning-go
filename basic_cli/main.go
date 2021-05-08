package main

import (
	"flag"	"os"
)


func main() {

	// define the values to get from the cli arguments
	var name string
	var greeting string
	var prompt bool
	var preview bool

	// Define flags
	flag.StringVar(&name, "name", "", "name to use inthe message")
	flag.StringVar(&greeting, "greeting", "", "phrase to use in the greeting message")
	flag.BoolVar(&prompt, "prompt", false, "use the prompt for input")
	flag.BoolVar(&preview, "preview", false, "Show preview of the message without writing to file")

	// parse the flags
	flag.Parse()

	// Optional, print flags base on OS environment
	if os.Getenv("DEBUG") != "" {
		// Print the values here
		// ...

		os.Exit(0)
	}

	// If no arguments parsed, show usage
	if prompt == false && (name == "" || greeting == "") {
		flag.Usage()
		os.Exit(1)
	}
}
