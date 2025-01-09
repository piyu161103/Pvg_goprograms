package main

import (
	"fmt"
)

func main() {

	var num1, num2 float64
	var choice int

	fmt.Println("Select an operation:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")

	fmt.Print("Enter your choice (1-4): ")
	fmt.Scan(&choice)

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	switch choice {
	case 1:
		fmt.Printf("Result: %.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case 2:
		fmt.Printf("Result: %.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case 3:
		fmt.Printf("Result: %.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case 4:
		if num2 != 0 {
			fmt.Printf("Result: %.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	default:
		fmt.Println("Invalid choice. Please select a valid operation (1-4).")
	}
}
