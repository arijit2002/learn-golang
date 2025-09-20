package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	fmt.Println("3 + 4 =", add(3, 4))

	first, second := swap("Go", "Lang")
	fmt.Println(first, second) // Lang Go
}
