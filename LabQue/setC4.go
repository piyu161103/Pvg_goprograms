package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	var input string
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}
	if num >= -9 && num <= 9 {
		fmt.Println("The number is a single digit.")
	} else {
		fmt.Println("The number is not a single digit.")
	}
}
