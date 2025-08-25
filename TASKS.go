package main

import (
	"fmt"
)

func main() {
	var num1, num2 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	product := num1 * num2
	sum := num1 + num2

	fmt.Printf("Product: %d\n", product)
	fmt.Printf("Sum: %d\n", sum)
}
