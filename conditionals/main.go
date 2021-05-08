package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["john"] = 18
	ages["phill"] = 18

	// If/else if/else
	if ages["john"] < 18 {
		fmt.Println("This is a tiny kid")
	} else if ages["john"] == 18 {
		fmt.Println("He is ready to vote")
	} else {
		fmt.Println("Too old now")
	}

	// Switch with condition on each line
	switch {
	case ages["john"] < 18:
		fmt.Println("This is a tiny kid")
	case ages["phill"] == 18:
		fmt.Println("phill is ready to vote")
	default:
		fmt.Println("Too old now")
	}

	switch ages["john"] {
	case 17:
		fmt.Println("This is a tiny kid")
	case 18, 19, 20:
		fmt.Println("He is ready to vote")
	default:
		fmt.Println("Too old now")
	}
}
