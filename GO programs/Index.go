package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "Hello, Go World!"
	fmt.Println(strings.Index(sampleString, "Go")) // Output: 7
}