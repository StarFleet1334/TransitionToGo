package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func negate(x *int) {
	neg := -(*x)
	*x = neg
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
