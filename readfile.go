package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Reads all of the lines from the shia labeouf text file into memory
func readFileIntoMem() []string {
	var filename = "shia-labeouf.txt"

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error Or Something")
		//Do something
	}
	lines := strings.Split(string(content), "   ")

	return lines
}

func readToken() string {
	var keyfile = "key.txt"

	content, err := ioutil.ReadFile(keyfile)
	if err != nil {
		fmt.Println("Error")
	}
	return string(content[:])
}
