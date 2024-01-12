//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var color = "blue"
	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	birthYear, age := 2002, 21
	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial = "J"
		lastInitial  = "N"
	)
	//* Declare (but don't assign!) a variable for your age (in days),
	var ageInDays int
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	ageInDays = 365 * age

	fmt.Println("My favorite color is:", color)
	fmt.Println("I was born in", birthYear, "and I'm ", age, "years old.")
	fmt.Println("My first initial is", firstInitial, "and last initial is", lastInitial)
	fmt.Println("I am", ageInDays, "days old")
}
