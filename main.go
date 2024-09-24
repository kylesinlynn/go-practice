package main

import "fmt"

// Define the Shape interface
type Shape interface {
	area() float64
	perimeter() float64
}

// Define the Rectangle struct
type Rectangle struct {
	width, height float64
}

// Implement the Shape interface for Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Define the Circle struct
type Circle struct {
	radius float64
}

// Implement the Shape interface for Circle
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// Function to print area and perimeter of a shape
func printShapeDetails(s Shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

func main() {
	// Create a rectangle and circle
	rect := Rectangle{width: 10, height: 5}
	circle := Circle{radius: 7}

	// Print details for the rectangle
	fmt.Println("Rectangle details:")
	printShapeDetails(rect)

	// Print details for the circle
	fmt.Println("Circle details:")
	printShapeDetails(circle)

}
