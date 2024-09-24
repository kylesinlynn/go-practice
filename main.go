package main

import "fmt"

func main() {
	// If-Else Statement
	age := 18

	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior.")
	}

	fmt.Println() // For better readability in output

	// Switch Statement
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	}

	fmt.Println() // For better readability in output

	// For Loop
	// Basic for loop
	fmt.Println("Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	fmt.Println() // For better readability in output

	// For loop with a condition
	j := 0
	fmt.Println("For loop with a condition:")
	for j < 3 {
		fmt.Println("Iteration:", j)
		j++
	}

	// Infinite loop (be careful!)
	// Uncomment the next lines to see it in action
	// fmt.Println("This will run forever!")
	// for {
	//     fmt.Println("This will run forever!")
	// }

}
