package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "Hello, Go World! Go Rocks!"
	fmt.Println(strings.LastIndex(sampleString, "Go")) // Output: 19
}