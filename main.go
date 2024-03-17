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
	Description string `json:"What it does"`
}

var bird Bird
var result map[string]any

func main() {

	// JSON Struct Tags - Custom Field Names
	birdJson := `{
				"birdType": "pigeon",
				"what it does":"likes to perch on rocks"
				}`
	err := json.Unmarshal([]byte(birdJson), &bird)
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}
	// So using Custom Field Names
	// We tell our code to which JSON property should our
	// attribute to be mapped to.
	fmt.Println(bird)

	// Decoding JSON to Maps - Unstructured Data.
	birdJson1 := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	err1 := json.Unmarshal([]byte(birdJson1), &result)
	if err1 != nil {
		return
	}
	birds := result["birds"].(map[string]any)
	for key, value := range birds {
		// Each value is an `any` type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}

}
