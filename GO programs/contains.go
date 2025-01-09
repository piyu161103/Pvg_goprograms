package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "Hello, Go World!"
	fmt.Println(strings.Contains(sampleString, "Go")) // Output: true
}