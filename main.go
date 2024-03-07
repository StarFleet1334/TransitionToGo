package main

import "fmt"

var greeting = "Hello world!"

func main() {
	fmt.Println("Hello World!")

	// declare pointer
	var ptr *string

	// initialize a greeting
	greeting := "Hello World!"

	// assign
	ptr = &greeting

	fmt.Println("Greeting:", greeting)
	fmt.Println("Address of greeting:", ptr)
	fmt.Println("Value stored in ptr:", *ptr)

}
