package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	Name      string
	Age       int
	Abilities []string
}

type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"What it does,omitempty"`
}

func main() {
	//birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"`
	//// Here we validate
	//if !json.Valid([]byte(birdJson)) {
	//	fmt.Println("Invalid JSON string:", birdJson)
	//	return
	//}
	// Marshalling is transforming structured data into JSON String
	// UnMarshalling it parsing raw json data into []byte variables

	//pigeon := &Bird{
	//	Species:     "Pigeon",
	//	Description: "likes to eat seed",
	//}
	//
	//data, _ := json.Marshal(pigeon)
	//fmt.Println(string(data))

	// Ignoring Empty Fields

	pigeon := &Bird{
		Species: "Pigeon",
	}

	data, _ := json.Marshal(pigeon)
	fmt.Println(string(data))

	// IF we print like this it will print
	// "what it does": "". say we want to ignore that field
	// in that case in structured data we should provide property,
	// omitempty
	// after that it will only print
	// "birdType": "Pigeon"
}
