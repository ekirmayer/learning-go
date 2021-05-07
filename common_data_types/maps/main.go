package main

import "fmt"

func main() {

	// maps defines key-value list.
	// it can be adgjust as needed

	birthdays := map[string]string{
		"you": "01/01/2011",
		"him": "01/01/2012",
		"she": "01/01/2010",
	}
	fmt.Println(birthdays)

	// reading a value
	fmt.Println(birthdays["him"])

	// removing an item
	delete(birthdays, "him")
	fmt.Println(birthdays)

	// adding items
	birthdays["dog"] = "03/07/2019"
	fmt.Println(birthdays)
}
