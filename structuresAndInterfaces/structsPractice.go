package main

import "fmt"

type city struct {
	name  string
	tempC float64
}

func createCity(name string) city {
	return city{
		name: name,
	}
}

func main() {

	all := city{
		name:  "Tbilisi",
		tempC: 4,
	}
	omitted := city{
		name: "Paris",
	}
	fmt.Printf("All fields are presented: %v\n", all)
	fmt.Printf("Some fields are omitted: %v\n", omitted)
	fmt.Printf("Using function: %v\n", createCity("Los Angeles"))
}
