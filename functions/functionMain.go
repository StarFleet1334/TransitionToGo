package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func negate(x *int) {
	neg := -(*x)
	*x = neg
}

// Here we use named return values and because of that
// we just use = operator instead of :=
func rectangle(x int, y int) (area int, circumf int) {
	area = x * y
	circumf = 2 * (x + y)
	return
}

// Here we do not use named return values, so we do not do naked return
// we use := operator and return results as tuple
func rectangleDefault(x int, y int) (int, int) {
	area := x * y
	circumf := 2 * (x + y)
	return area, circumf
}

// Here we use shorthand variable in if statement
func printParity(x int) {
	if r := x % 2; r == 0 {
		fmt.Printf("%v is even.\n", x)
		return
	}
	// r is out of scope
	fmt.Printf("%v is odd.\n", x)
}

func main() {

	// This is referred to anonymous function
	sub := func(x int, y int) int {
		return x - y
	}
	fmt.Printf("Sub: %v\n", sub(10, 3))
	fmt.Printf("Add: %v\n", add(10, 3))

	x := 3
	fmt.Printf("Before negation: %v\n", x)
	negate(&x)
	fmt.Printf("After negation: %v\n", x)

}
