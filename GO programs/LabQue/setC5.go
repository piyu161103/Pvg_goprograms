package main

import (
	"fmt"
	"strings"
)

func main() {
	
	var str1, str2 string
	fmt.Print("Enter the first string: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter the second string: ")
	fmt.Scanln(&str2)

	if strings.Contains(str2, str1) {
		fmt.Printf("'%s' is a substring of '%s'.\n", str1, str2)
	} else {
		fmt.Printf("'%s' is NOT a substring of '%s'.\n", str1, str2)
	}
}
