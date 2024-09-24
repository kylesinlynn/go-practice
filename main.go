package main

import "fmt"

// Define the Rectangle struct
type Rectangle struct {
	width, height float64
}

// Method to calculate area of the rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Method to calculate perimeter of the rectangle
func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Method with a pointer receiver to modify the rectangle
func (r *Rectangle) resize(newWidth, newHeight float64) {
	r.width = newWidth
	r.height = newHeight
}

func main() {
	// Create a new rectangle
	rect := Rectangle{width: 10, height: 5}

	// Calculate and display the area and perimeter
	fmt.Println("Area:", rect.area())
	fmt.Println("Perimeter:", rect.perimeter())

	// Modify the rectangle dimensions
	rect.resize(20, 10)
	fmt.Println("Resized Rectangle - Area:", rect.area())
	fmt.Println("Resized Rectangle - Perimeter:", rect.perimeter())
}
