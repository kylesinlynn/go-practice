package main

import (
	"errors"
	"fmt"
)

// Function that returns an error if division by zero is attempted
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Function that returns an error if a value is out of bounds
func checkBounds(val, min, max int) error {
	if val < min || val > max {
		return fmt.Errorf("value %d is out of bounds (%d - %d)", val, min, max)
	}
	return nil
}

// Function to demonstrate defer, panic, and recover
func riskyFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Running riskyFunction")
	panic("Something went wrong!")
}

func main() {
	// Error handling for division
	result, err := divide(0, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	fmt.Println() // For better readability in output

	// Error handling for bounds checking
	err = checkBounds(15, 1, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value is within bounds.")
	}

	fmt.Println() // For better readability in output

	riskyFunction()
	fmt.Println("Program continues after recover.")
}
