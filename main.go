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

var player Player

func main() {
	// Two Types of Data, when working with JSON:
	// 1. Structured Data.
	// 2. UnStructured Data.

	// Structured data
	rawJson := `{"name": "Jax","age": 2400,"abilities": ["Raw-Strength","Fly","Invincibility"]}`
	err := json.Unmarshal([]byte(rawJson), &player)
	if err != nil {
		return
	}
	fmt.Printf("Player name: %s\nPlayer age: %d\nPlayer abilities: %q\n", player.Name, player.Age, player.Abilities)

}
