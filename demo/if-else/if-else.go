package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 7, 7, 9

	if quiz1 > quiz2 {
		fmt.Println("Quiz1 scored higher than Quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz2 scored higher than Quiz1")
	} else {
		fmt.Println("Quiz1 scored the same as Quiz2")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("The scores were good")
	} else {
		fmt.Println("The instructor did a bad job!")
	}
}
