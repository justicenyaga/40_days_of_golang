package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

func mine(lhs, rhs int) {
	fmt.Println("Nothing to say.. just passing by", lhs, rhs)
}

func compute(lhs, rhs int, op func(lhs, rhs int) int) int {
	fmt.Printf("Running a computation with %v & %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {
	fmt.Println("2 + 2 =", compute(2, 2, add))

	fmt.Println("10 - 2 =", compute(10, 2, func(lhs, rhs int) int { return lhs - rhs }))

	mult := func(lhs, rhs int) int {
		fmt.Printf("multiplying %v * %v = ", lhs, rhs)
		return lhs * rhs
	}
	fmt.Println(compute(3, 3, mult))
}
