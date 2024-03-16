package main

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
	Species     string
	Description string
	Dimensions  Dimensions
}

// var player Player
// var birds []Bird
//var birds Bird
//var n int
//
//var pi float64
//var str string

func main() {
	// Two Types of Data, when working with JSON:
	// 1. Structured Data.
	// 2. UnStructured Data.

	// Structured data
	//rawJson := `{"name": "Jax","age": 2400,"abilities": ["Raw-Strength","Fly","Invincibility"]}`
	//err := json.Unmarshal([]byte(rawJson), &player)
	//if err != nil {
	//	return
	//}
	//fmt.Printf("Player name: %s\nPlayer age: %d\nPlayer abilities: %q\n", player.Name, player.Age, player.Abilities)
	//birdJson := `[{"species":"pigeon","description":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	//err := json.Unmarshal([]byte(birdJson), &birds)
	//if err != nil {
	//	return
	//}
	//fmt.Printf("Birds : %+v\n", birds)

	// Nested Objects

	//birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	//err := json.Unmarshal([]byte(birdJson), &birds)
	//if err != nil {
	//	return
	//}
	//fmt.Println(birds)
	//
	//// Primitive Data Types
	//
	//numberJson := "3"
	//floatJson := "3.1414"
	//stringJson := `"bird"`
	//
	//json.Unmarshal([]byte(numberJson), &n)
	//json.Unmarshal([]byte(floatJson), &pi)
	//json.Unmarshal([]byte(stringJson), &str)
	//
	//fmt.Println(n)
	//fmt.Println(pi)
	//fmt.Println(str)

}
