package arrays

import "fmt"

func addCity(city string, cities *[2]string) {
	cities[1] = city
}

func Test() {
	cities := [2]string{"London"}
	copy := cities

	cities[1] = "New York"
	addCity("Paris", &copy)

	fmt.Println("Cities: ", cities)
	fmt.Println("Copy: ", copy)

}
