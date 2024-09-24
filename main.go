package main

import "fmt"

// Function that take two integers and returns their sum
func add(a int, b int) int {
	return a + b
}

// Function that returns the sum and difference of two integers
func sumAndDifference(a int, b int) (int, int) {
	return a + b, a - b
}

// Function with named return values
func namedReturns(a int, b int) (sum int, difference int) {
	sum = a + b
	difference = a - b
	return // Returns named return values
}

func main() {
	// Calling the add function
	result := add(5, 7)
	fmt.Println("The sum is:", result)

	// Calling sumAndDifference function
	sum, diff := sumAndDifference(3, 5)
	fmt.Println("Sum:", sum, ", Difference:", diff)

	// Calling namedReturns function
	sumNamed, diffNamed := namedReturns(4, 2)
	fmt.Println("Sum:", sumNamed, ", Difference:", diffNamed)
}
