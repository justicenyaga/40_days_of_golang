package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from the greet function")
}

func main() {
	greet()

	dozen := double(6)
	fmt.Println("A dozen has", dozen, "items")

	backersDozen := add(dozen, 1)
	fmt.Println("A backersDozen has", backersDozen, "items")
}
