package main

import "fmt"

func main() {
	// Explicit type
	var name string = "Arijit"
	var age int = 25

	// Type inference
	city := "Kolkata"
	score := 95.5

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Score:", score)
}
