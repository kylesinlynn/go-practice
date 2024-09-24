package main

import "fmt"

func main() {
	// Basic types
	var a int = 10
	var b float64 = 20.5
	var c bool = true
	var d string = "Hello World!"

	// Composite types
	arr := [3]int{1, 2, 3}
	slice := []string{"apple", "banana", "cherry"}
	m := map[string]int{"one": 1, "two": 2}

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Kyel", Age: 23}

	// Pointer
	var ptr *int = &a

	fmt.Println("Integer:", a)
	fmt.Println("Float:", b)
	fmt.Println("Boolean:", c)
	fmt.Println("String:", d)
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
	fmt.Println("Map:", m)
	fmt.Println("Struct:", p)
	fmt.Println("Pointer:", ptr)

}
