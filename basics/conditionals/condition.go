package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	if score := 75; score >= 50 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}
