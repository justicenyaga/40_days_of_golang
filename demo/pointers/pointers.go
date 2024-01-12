package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}

	hello, world := "Hello", "World!"
	fmt.Println(hello, world)

	increment(&counter)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrases := []string{hello, world}
	replace(&phrases[1], "Go!", &counter)
	fmt.Println(phrases)
}
