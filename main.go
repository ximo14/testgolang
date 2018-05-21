package main

import "fmt"
import "strconv"

func main() {
	fmt.Printf("Hello world!\n")
	fmt.Printf(variables())
}

func variables() string {
	var text string = "I'm the string"
	var number int = 14
	// strconv.Itoa -> integer to string
	var numberToString = strconv.Itoa(number)

	return ("This is a string: "+ text + "\n"+"This is an integer: " + numberToString + "\n")
}

