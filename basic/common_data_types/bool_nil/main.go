package main

import "fmt"

func main() {
	fmt.Println("Greater than: ", 4 > 3)
	fmt.Println("Less than: ", 2 < 3)
	fmt.Println("Greater than or equal: ", 2 >= 2)
	fmt.Print("Less than or equal", 4 <= 4)
	fmt.Println("Equivalent:", 4.0 == 4)
	fmt.Println("Not equivalent:", 4.1 != 4.1)

	/*
		Testing for Null == nil
		if err != nil {
			//do something
		}
	*/
}
