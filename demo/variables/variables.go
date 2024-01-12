package main

import "fmt"

func main() {
	var myName = "Justice Nyaga"
	fmt.Println("My name is", myName)

	var name string = "Amos"
	fmt.Println("name =", name)

	username := "admin"
	fmt.Println("username =", username)

	var sum int
	fmt.Println("The sum is =", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "and other is", other)

	part2, other := 2, 0
	fmt.Println("part 2 is", part2, "and other is", other)

	sum = part1 + part2
	fmt.Println("Sum =", sum)

	var (
		lessonName = "Variable"
		lessonType = "Demo"
	)
	fmt.Println("lessonName =", lessonName)
	fmt.Println("lessonType =", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}
