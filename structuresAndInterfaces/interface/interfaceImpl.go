package _interface

import "fmt"

// Calculator Any type that wants to satisfy this interface
// must implement the Add() and Subtract() methods,
// which take two integers as arguments and return
// an integer result.
type Calculator interface {
	Add(a, b int) int
	Subtract(a, b int) int
}

type SimpleCalculator struct{}

func (c SimpleCalculator) Add(a, b int) int {
	return a + b
}

func (c SimpleCalculator) Subtract(a, b int) int {
	return a - b
}

func PerformCalculation(x, y int, c Calculator) {
	fmt.Println("Added function: ", c.Add(x, y))
	fmt.Println("Subtract function: ", c.Subtract(x, y))

}
