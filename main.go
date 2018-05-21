package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func main() {
	fmt.Printf("Hello world!\n")
	fmt.Printf(variables())

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func variables() string {
	var text string = "I'm the string"
	var number int = 14
	// strconv.Itoa -> integer to string
	var numberToString = strconv.Itoa(number)

	return ("This is a string: "+ text + "\n"+"This is an integer: " + numberToString + "\n")
}
