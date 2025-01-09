package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "Hello, Go World!"
	fmt.Println(strings.Replace(sampleString, "Go", "Golang", -1)) // Output: "Hello, Golang World!"
}