package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("First method (Arithmetics):")
	a = 4234
	b = 1252
	fmt.Printf("Before swapping a: %d, b: %d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("After swapping a: %d, b: %d\n", a, b)

	fmt.Println("Second method (Binary operations):")
	a = 4234
	b = 1252
	fmt.Printf("Before swapping a: %d, b: %d\n", a, b)
	// Побитовое XOR
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("After swapping a: %d, b: %d\n", a, b)
}
