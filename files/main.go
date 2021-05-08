package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	//  writing to file with ioutil
	var ioutil_text string = "This is written to file using the ioutil package"
	_ = ioutil.WriteFile("output/ioutil.txt", []byte(ioutil_text), 0664)

	// Working with files by openning them to read/write
	f, err := os.OpenFile("./output/openfile.txt", os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error: unable to open the file for writing.", err)
		os.Exit(1)
	}
	defer f.Close()
	var os_messsage string = "This file is using os.OpenFile"
	_, err = f.Write([]byte(os_messsage))
	if err != nil {
		fmt.Println("Error: unable to write to the file")
		os.Exit(1)
	}
}
