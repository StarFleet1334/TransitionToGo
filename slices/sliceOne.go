package slices

import "fmt"

func Test() {
	cities := make([]string, 1, 1)
	copy := cities
	cities[0] = "London"
	cities = append(cities, "New York")
	fmt.Printf("[%p]cities= %v, len = %v, cap = %v\n", &cities, cities, len(cities), cap(cities))
	fmt.Printf("[%p]cities= %v, len = %v, cap = %v\n", &copy, copy, len(copy), cap(copy))

}
