package main

import "fmt"

func main() {

	// simple for loop
	for i := 0; i < 5; i++ {
		fmt.Println("The number is:", i)
	}

	ages := map[string]int{}
	ages["john"] = 11
	ages["phill"] = 28
	ages["tom"] = 67
	ages["lura"] = 18
	ages["joe"] = 16

	// Iterate over each item (This is like for each loop)
	for name, age := range ages {
		switch age {
		case 1, 2, 3, 5, 7, 11, 13, 17, 19:
			fmt.Println(name, "This is a tiny kid")
		case 16:
			fmt.Println(name, "He is ready to vote")
		case 18:
			fmt.Println(name, "He is ready to vote")
		case 67:
			fmt.Println(name, "He is ready to vote")
		default:
			fmt.Println(fmt.Sprintf("Nothing to tell about %s", name))
		}
	}

	// Conditional looping (This is like a while loop)
	a := 0
	for a < 4 {
		fmt.Println("We are runnig a while loop. Loop no:", a+1)
		a++
	}

	// Controling loop stopping
	b := 0
	for b < 10 {
		if b%2 == 0 {
			b++
			continue
		} else if b == 5 {
			break
		}
		fmt.Println("Continue counting:", b)
		b++
	}
}
