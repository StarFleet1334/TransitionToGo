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

func (c city) getNameOfCity() string {
	return c.name
}

func (c city) changCityName() {
	c.name = "Stockholm"
}

func (c *city) changedCityNameCorrect() {
	c.name = "Stockholm"
}

// This would give an error, because method has same name as
// structs field. That is because city has both field and
// method name same
//func (c city) tempC() float64{
//
//}

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

	fmt.Printf("Name of the city: %s\n", omitted.getNameOfCity())

	// Here we try to change name
	all.changCityName()
	// But we see it did not change, cause here we do not use
	// pointer
	fmt.Printf("All fields are presented: %v\n", all)

	// Let's call correct one here
	all.changedCityNameCorrect()
	fmt.Printf("All fields are presented: %v\n", all)

}
