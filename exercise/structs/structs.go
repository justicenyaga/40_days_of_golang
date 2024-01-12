//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate // Top left
	b Coordinate // Bottom right
}

func width(rect Rectangle) int {
	return rect.b.x - rect.a.x
}

func height(rect Rectangle) int {
	return rect.a.y - rect.b.y
}

func area(rect Rectangle) int {
	return width(rect) * height(rect)
}

func perimeter(rect Rectangle) int {
	return (width(rect) + height(rect)) * 2
}

func main() {
	rect := Rectangle{a: Coordinate{0, 5}, b: Coordinate{8, 0}}
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimeter is", perimeter(rect))

	rect.a.y *= 2
	rect.b.x *= 2
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimeter is", perimeter(rect))
}
