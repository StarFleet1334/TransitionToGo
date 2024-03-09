package _interface

import "fmt"

// This function takes an empty interface as its argument.
// This means it can accept any value, regardless of its type.
func PrintAny(value interface{}) {
	fmt.Println(value)
}
