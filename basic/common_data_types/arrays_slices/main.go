package main

import "fmt"

func main() {
	//  Arrays are pre defined length
	// Define an array and set the values
	names := [4]string{"you", "him"}
	fmt.Println(names)

	names[2] = "this"
	names[3] = "there"
	fmt.Println(names)

	// define the array and set values later
	var numbers [2]int
	numbers[0] = 1
	numbers[1] = 2
	fmt.Println(numbers)

	// Slice
	sl := []string{}
	sl = append(sl, "ken")
	sl = append(sl, "them", "bla")
	fmt.Println(sl)

	// make
	mk := make([]string, 4)
	mk[0] = "one"
	mk[1] = "two"
	mk[2] = "three"
	mk[3] = "four"
	fmt.Println(mk)

	mk = append(mk, "five")
	fmt.Println(mk)

	// This will not work as numbers is an array
	// numbers = append(numbers, 4)
}
